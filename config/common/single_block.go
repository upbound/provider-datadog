/*
Copyright 2026 Upbound Inc.
*/

// Package common holds resource configuration helpers shared by the
// cluster-scoped and the namespaced configurators.
package common

import "github.com/crossplane/upjet/v2/pkg/config"

// EmbedSingleNestedBlocks embeds the given plugin-framework single-nested
// blocks as objects without a singleton list conversion. Terraform exchanges
// a single-nested block as an object, while the SingletonListEmbedder treats
// every MaxItems=1 list as a singleton list to convert at runtime, which would
// break the state round trip for these blocks.
func EmbedSingleNestedBlocks(r *config.Resource, tfPaths ...string) {
	for _, p := range tfPaths {
		r.RemoveSingletonListConversion(p)
		r.SchemaElementOptions.SetEmbeddedObject(p)
	}
}

// MarkSingleNestedBlockConfigurable clears the Computed flag that upjet's
// tfjson conversion infers for every plugin-framework single-nested block.
// Without it a block with a required descendant is generated only under
// status.atProvider. The conversion also marks such a block required, so pass
// optional for a block that upstream treats as optional. The mutated schema
// only feeds code generation: the resource runs through the Terraform CLI.
func MarkSingleNestedBlockConfigurable(r *config.Resource, tfPath string, optional bool) {
	s := r.TerraformResource.Schema[tfPath]
	s.Computed = false
	if optional {
		s.Required = false
		s.Optional = true
	}
}
