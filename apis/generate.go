//go:build generate
// +build generate

/*
Copyright 2026 Upbound Inc.
*/

// NOTE: See the below link for details on what is happening here.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

// Remove existing CRDs
//go:generate rm -rf ../package/crds

// Remove generated files
//go:generate bash -c "find . -iname 'zz_*' ! -iname 'zz_generated.managed*.go' -delete"
//go:generate bash -c "find . -type d -empty -delete"
//go:generate bash -c "find ../internal/controller -iname 'zz_*' -delete"
//go:generate bash -c "find ../internal/controller -type d -empty -delete"
//go:generate rm -rf ../examples-generated

// Generate documentation from Terraform docs.
//go:generate go run github.com/crossplane/upjet/v2/cmd/scraper -n ${TERRAFORM_PROVIDER_SOURCE} -r ../.work/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_DOCS_PATH} -o ../config/provider-metadata.yaml

// Run Upjet generator
//go:generate go run ../cmd/generator/main.go ..

// Generate deepcopy methodsets and CRD manifests
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=./... crd:allowDangerousTypes=true,crdVersions=v1 output:artifacts:config=../package/crds

// Generate crossplane-runtime methodsets (resource.Claim, etc)
//go:generate go run -tags generate github.com/crossplane/crossplane-tools/cmd/angryjet generate-methodsets --header-file=../hack/boilerplate.go.txt ./...

// Transform the generated cross-resource reference resolvers to look the
// referenced kinds up through the resolver runtime scheme instead of importing
// their API packages, which would create import cycles. The API group is the
// resource short group plus the suffix, so the suffix is the bare root group.
//go:generate go run github.com/crossplane/upjet/v2/cmd/resolver -g upbound.io -a github.com/upbound/provider-datadog/internal/apis -s -p ./cluster/...
//go:generate go run github.com/crossplane/upjet/v2/cmd/resolver -g m.upbound.io -a github.com/upbound/provider-datadog/internal/apis -s -p ./namespaced/...

package apis

import (
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen" //nolint:typecheck

	_ "github.com/crossplane/crossplane-tools/cmd/angryjet" //nolint:typecheck
	_ "github.com/crossplane/upjet/v2/cmd/resolver"         //nolint:typecheck
	_ "github.com/crossplane/upjet/v2/cmd/scraper"          //nolint:typecheck
)
