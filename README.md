# Provider Datadog

`provider-datadog` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Datadog API.

## Prerequisites

This provider interacts with a
[Datadog account](https://www.datadoghq.com/). It authenticates
to the account using a Datadog API Key, an Application key, and
a Datadog account endpoint URL.
The keys can be generated inside the account
and be stored in a Kubernetes secret on the Crossplane
management cluster. The format of the secret is as follows:
```
    {
      "api_key": "INSERT_API_KEY",
      "app_key": "INSERT_APP_KEY",
      "api_url": "https://api.datadoghq.com/"
    }
```
Note that your preferred endpoint may differ.
The Kubernertes secret can be referenced by
the ProviderConfig, so that the provider-datadog can connect
to the desired Datadog account. A ProviderConfig may look
as follows:
```
apiVersion: datadog.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: datadog-creds
      namespace: upbound-system
      key: credentials
```

To run local tests, create a datadog-secret file per above.
Then create an `UPTEST_CLOUD_CREDENTIALS` environment variable
as follows
```
export UPTEST_CLOUD_CREDENTIALS=$(cat <PATH-TO-DATADOG-SECRET-FILE> )
```
Once complete, specify the tests that you would like to run
in the `UPTEST_EXAMPLE_LIST` environment variable. An example
is as follows:
```
export UPTEST_EXAMPLE_LIST="./examples/datadog/dashboardjson.yaml"
```
Note that you may specify multiple comma separated tests.
Now run `UPTEST_EXAMPLE_LIST="./examples/datadog/v1alpha1/dashboardjson.yaml" make e2e`. This will create a local kind cluster,
install Crossplane and the provider-datadog from a local build
and run Uptests managed resources apply, update, import, delete
tests.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/upbound/provider-datadog):
```
up ctp provider install upbound/provider-datadog:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-datadog
spec:
  package: upbound/provider-datadog:v0.1.0
EOF
```

You can see the API reference [here](https://doc.crds.dev/github.com/upbound/provider-datadog).

## Developing

Run code-generation pipeline:

```console
make clean; make gen-clean; make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/upbound/provider-datadog/issues).
