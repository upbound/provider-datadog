/*
Copyright 2026 Upbound Inc.
*/

// Package common holds resource configuration helpers shared by the
// cluster-scoped and the namespaced configurators.
package common

import (
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
// tfjson conversion infers for every plugin-framework single-nested block
// with at least one required descendant. Without it, such a block is
// generated only under status.atProvider. The conversion also marks the
// block required, so pass optional for a block that upstream treats as
// optional. tfPath may be a dotted path into a nested block, e.g.
// "list_block.single_block"; only the final segment's flags are changed.
//
// Safe only for a plugin framework resource: for a plugin SDK resource,
// TerraformResource is the live schema its controller uses for real CRUD,
// and mutating it here would corrupt that client.
func MarkSingleNestedBlockConfigurable(r *config.Resource, tfPath string, optional bool) {
	segments := strings.Split(tfPath, ".")
	s := r.TerraformResource.Schema[segments[0]]
	for _, seg := range segments[1:] {
		s = s.Elem.(*schema.Resource).Schema[seg]
	}
	s.Computed = false
	if optional {
		s.Required = false
		s.Optional = true
	}
}
