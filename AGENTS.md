# provider-datadog — Agent Guide

This is a **Crossplane v2 Upjet provider** for [Datadog](https://www.datadoghq.com/). It wraps the [`DataDog/terraform-provider-datadog`](https://github.com/DataDog/terraform-provider-datadog) Go module (imported via a `replace` in `go.mod`, not shelled out to) and generates cluster-scoped and namespaced Crossplane managed resources from a declarative config, with no hand-written controllers.

> **Adding a resource**: Use the `add-upjet-resource` skill.

---

## Architecture

### In-process Terraform provider (no forked CLI)

Unlike most Upjet providers, `provider-datadog` never shells out to a Terraform binary at runtime. It imports the wrapped provider's Go packages directly and drives them in-process:

- `github.com/terraform-providers/terraform-provider-datadog/datadog` — the **plugin SDK v2** provider (`datadog.Provider()`).
- `github.com/terraform-providers/terraform-provider-datadog/datadog/fwprovider` — the **plugin framework** provider (`fwprovider.New()`).

Every resource is implemented by exactly one of the two; `config/external_name.go` records the split in `terraformPluginSDKExternalNameConfigs` / `terraformPluginFrameworkExternalNameConfigs`, and `config/provider.go` turns each into an upjet include list (`WithTerraformPluginSDKIncludeList` / `WithTerraformPluginFrameworkIncludeList`). `WithIncludeList([]string{})` is passed empty — **no resource runs through the Terraform CLI**. `cmd/provider/main.go` constructs both providers once and shares them across the cluster-scoped and namespaced controller sets via `clients.TerraformSetupBuilder`.

**Runtime schema-mutation safety — read before touching a resource's `TerraformResource.Schema`:**
- For a **plugin SDK** resource, `r.TerraformResource` at runtime *is* the live `*schema.Resource` used by that resource's controller for real CRUD. Mutating its `.Schema` in place (changing a field's type, deleting a field, flipping `.Sensitive`) corrupts the live client.
- For a **plugin framework** resource, runtime CRUD goes through a separate framework client; `r.TerraformResource.Schema` is only used to derive the CRD and is safe to mutate.
- `GetProvider`/`GetProviderNamespaced` take a `generationProvider bool`. When `true` (only `cmd/generator`), `sdkProvider` is swapped for a throwaway provider built from the JSON schema document (`codegenSDKProvider`) before any configurator runs, so schema mutations for plugin SDK resources are safe there and unsafe everywhere else. `config/codegen.go`'s `codegenConfigurators` map holds exactly the mutations that need this gate; new schema-mutating configurators for a plugin SDK resource belong there, not in a group's `config.go`.
- `config/common/single_block.go`'s `MarkSingleNestedBlockConfigurable`/`EmbedSingleNestedBlocks` mutate `TerraformResource.Schema` unconditionally — safe only for plugin framework resources (its doc comment predates the in-process switch and is stale on this point; verify SDK vs. framework per resource before reusing it on a new one).

### Dual-scope generation

Every resource is generated **twice** from the same `config/`:

| Scope | Root group | ProviderConfig kind | Example CRD group |
|---|---|---|---|
| Cluster | `upbound.io` | `ProviderConfig` | `datadog.upbound.io`, `synthetics.datadog.upbound.io` |
| Namespaced | `m.upbound.io` | `ClusterProviderConfig` (or namespaced `ProviderConfig`) | `datadog.m.upbound.io` |

CRD group formula: `<shortGroup>.<rootGroup>`, where `shortGroup` is set per-resource (see below). A namespaced resource that omits `spec.providerConfigRef` uses the `ClusterProviderConfig` named `default`.

`config/provider.go` calls `GetProvider(sdkProvider, fwProvider, generationProvider)` (cluster) and `GetProviderNamespaced(...)` (namespaced), each delegating to `newProvider`. `cmd/generator/main.go` runs `pipeline.Run(GetProvider(...), GetProviderNamespaced(...), rootDir)` with `generationProvider=true`; `cmd/provider/main.go` builds the same two with `generationProvider=false`.

### Codegen pipeline

`make generate` runs these steps in order:
1. Downloads Terraform ≤1.5 via `build/` submodule tools (feeds codegen only — never the runtime provider, which is in-process).
2. Initialises a minimal `.work/terraform/main.tf.json` and runs `terraform providers schema -json` → **`config/schema.json`** (do not edit).
3. Clones the TF provider repo (sparse, tag pinned) and scrapes its docs → **`config/provider-metadata.yaml`** (do not edit).
4. Runs `go generate ./apis/...` which invokes:
   - `cmd/generator/main.go` → rewrites `apis/cluster/`, `apis/namespaced/`, `internal/controller/cluster/`, `internal/controller/namespaced/`, `examples-generated/`, `config/generated.lst`.
   - `controller-gen` → `package/crds/`.
   - `angryjet` → managed-resource methodsets.

---

## Repository structure

```
config/
  provider.go                  ← GetProvider/GetProviderNamespaced/newProvider, WithRootGroup     ✏️
  external_name.go              ← terraformPluginSDK/FrameworkExternalNameConfigs (gates generation)✏️
  codegen.go                    ← schema mutations that must run ONLY for generationProvider=true   ✏️
  schema.json                   ← TF provider schema (embedded)                                    🚫 generated
  provider-metadata.yaml        ← scraped TF provider docs                                          🚫 generated
  generated.lst                 ← list of generated resource names                                  🚫 generated
  cluster/cluster.go            ← Configurators []func(*ujconfig.Provider) (cluster scope)          ✏️
  namespaced/namespaced.go      ← same, namespaced scope (must match cluster)                       ✏️
  cluster/<group>/config.go     ← per-group AddResourceConfigurator (cluster scope)                 ✏️
  namespaced/<group>/config.go  ← same, namespaced scope (must match cluster)                       ✏️
  common/                       ← shared configurator helpers (single-nested-block handling)        ✏️
  test/roundtrip/               ← API round-trip (webhook conversion) tests                         ✏️

apis/
  cluster/                      ← generated API types (zz_* files); version dir varies per group     🚫 generated
  namespaced/                   ← generated API types (zz_* files)                                   🚫 generated
  generate.go                   ← //go:generate directives that drive make generate                  ✏️ rarely

internal/
  apis/scheme.go                ← resolver runtime scheme registration (distinct from apis/ above)   ✏️ rarely
  clients/                      ← TerraformSetupBuilder, credential extraction, PC resolver           ✏️
  controller/cluster/           ← generated controllers                                              🚫 generated
  controller/namespaced/        ← generated controllers                                              🚫 generated
  features/                     ← feature flags                                                      ✏️ rarely

cmd/
  generator/main.go             ← runs pipeline.Run(GetProvider(..., true), GetProviderNamespaced(..., true))  ✏️ rarely
  provider/main.go              ← provider binary entry point; builds both providers once, generationProvider=false ✏️ rarely

package/
  crossplane.yaml                ← provider name in the OCI package                                 ✏️
  crds/                          ← generated CRD manifests                                           🚫 generated

examples/
  cluster/<group>/<version>/     ← curated E2E examples (cluster scope)                              ✏️
  namespaced/<group>/<version>/  ← curated E2E examples (namespaced scope)                            ✏️
  cluster/providerconfig/        ← ProviderConfig + Secret template                                  ✏️
  namespaced/providerconfig/     ← ProviderConfig + ClusterProviderConfig + Secret template           ✏️
  install.yaml                   ← provider install manifest                                         ✏️

examples-generated/              ← raw generated examples (starting point only)                      🚫 generated

cluster/test/setup.sh            ← E2E cluster setup: creates Secret + ProviderConfig                ✏️

Makefile                         ← provider knobs at the top (see below)                             ✏️
build/                           ← crossplane/build submodule — do not edit                          🚫 submodule
```

`✏️` = hand-written (safe to edit). `🚫` = generated or submodule (never edit directly). API version directories are **not uniformly `v1beta1`** — most groups are, but `synthetics`, `sensitivedatascanner`, and `securitymonitoring` are `v1alpha1`; check `apis/<scope>/<group>/` for the actual version before writing an example.

---

## Makefile knobs

The provider-specific variables at the top of `Makefile`:

```makefile
PROJECT_NAME                    # provider-datadog
TERRAFORM_PROVIDER_SOURCE       # DataDog/datadog
TERRAFORM_PROVIDER_REPO         # https://github.com/DataDog/terraform-provider-datadog (docs scrape only)
TERRAFORM_PROVIDER_VERSION      # pinned Datadog TF provider version (feeds codegen + doc scrape only —
                                 # the runtime provider version comes from the go.mod replace, see below)
TERRAFORM_DOCS_PATH             # docs/resources
```

`PROJECT_REPO` is derived as `github.com/upbound/$(PROJECT_NAME)` — matches the Go module path `github.com/upbound/provider-datadog`. The actual in-process provider version is pinned by the `replace github.com/terraform-providers/terraform-provider-datadog => github.com/DataDog/terraform-provider-datadog <version>` line in `go.mod`; keep it in step with `TERRAFORM_PROVIDER_VERSION` when bumping.

---

## Key commands

```bash
# After any fresh clone or worktree — MUST run first
git submodule update --init --recursive

# Full codegen (schema pull + doc scrape + code generation)
make generate

# Run only the Go generator (schema + metadata already present)
go run cmd/generator/main.go "$PWD"

# Verify compilation
go build ./...

# Run provider out-of-cluster (needs a kubeconfig)
make run

# E2E test (builds provider, spins up a KinD cluster, runs uptest/chainsaw)
UPTEST_EXAMPLE_LIST="examples/cluster/<group>/<version>/<kind>.yaml" \
UPTEST_CLOUD_CREDENTIALS="$(cat path/to/creds.json)" \
UPTEST_DATASOURCE_PATH="path/to/datasource.ini" \
make e2e

# After any E2E run — ALWAYS clean up in this order
kubectl delete managed --all --all-namespaces
kind delete cluster --name local-dev
```

---

## Config key concepts

### ExternalNameConfigs gates generation, split by SDK

A resource absent from **both** `terraformPluginSDKExternalNameConfigs` and `terraformPluginFrameworkExternalNameConfigs` in `config/external_name.go` is **never generated**. Adding a resource here is always step one — and you must pick the right map: check the resource's own source file in the wrapped provider (plugin SDK `schema.Resource` vs. plugin framework `resource.Resource`), not just what the provider module as a whole supports.

### Group and Kind naming

There is no `config/groups.go` in this provider. `shortGroup`/`Kind` are set inline, per resource, inside its `AddResourceConfigurator` call (`r.ShortGroup = "..."`, `r.Kind = "..."`) in `config/{cluster,namespaced}/<group>/config.go`.

### References

Cross-resource references (`r.References["field"] = config.Reference{...}`) go in both `config/cluster/<group>/config.go` AND `config/namespaced/<group>/config.go`. These files must stay in sync.

### Sensitive fields

Any TF attribute marked `Sensitive: true` becomes `<field>SecretRef` in the CRD (references a `v1/Secret`). The plain field name is absent from `forProvider`. Check CRD `keys` when a field seems missing.

### codegen.go: schema mutations gated on generationProvider

`config/codegen.go`'s `codegenConfigurators` holds resource-schema mutations that are only safe when `generationProvider=true` — see the runtime schema-mutation safety note above. Currently covers `datadog_powerpack` (widget list collapsed to a JSON string) and `datadog_synthetics_test` (a sensitive numeric field exposed as plain). Framework-resource mutations that are safe unconditionally (e.g. dropping write-only attributes on `datadog_integration_fastly_account`, `datadog_synthetics_global_variable`) live there too, ungated, for locality with the other schema tweaks on those same resources.

---

## E2E setup

`cluster/test/setup.sh` runs during `make e2e` (passed via `--setup-script=cluster/test/setup.sh` in the Makefile's `uptest e2e` invocation). It:
1. Creates a `provider-secret` Secret in `upbound-system` from `$UPTEST_CLOUD_CREDENTIALS` (JSON with `api_key`, `app_key`, `api_url`).
2. Applies a `ProviderConfig` (cluster scope) and a `ClusterProviderConfig` (namespaced scope).

The `UPTEST_DATASOURCE_PATH` file resolves `${data.<key>}` placeholders in example YAML files.

---

## Crossplane v2 notes

- Both scopes are `SafeStart`-capable (`package/crossplane.yaml`).
- Cluster-scoped ProviderConfig kind for namespaced resources is **`ClusterProviderConfig`** (not `ProviderConfig`).
- Managed resources in the namespaced scope reference it with `providerConfigRef.kind: ClusterProviderConfig`.
- Namespace for namespaced managed resources and Secrets: **`upbound-system`** in examples (or whatever namespace the user creates them in).
