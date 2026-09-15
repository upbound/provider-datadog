/*
Copyright 2026 Upbound Inc.
*/

package config

import (
	"context"
	"slices"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/go-cmp/cmp"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/terraform-providers/terraform-provider-datadog/datadog"
	datadogfw "github.com/terraform-providers/terraform-provider-datadog/datadog/fwprovider"
)

func TestExternalNameConfigsMatchImplementingLibrary(t *testing.T) {
	sdkResources := map[string]bool{}
	for name := range datadog.Provider().ResourcesMap {
		sdkResources[name] = true
	}
	type args struct {
		configs     map[string]config.ExternalName
		implemented map[string]bool
	}
	cases := map[string]struct {
		reason string
		args   args
		want   []string
	}{
		"TerraformPluginSDK": {
			reason: "Every resource in the plugin SDK table must be implemented by the plugin SDK provider.",
			args: args{
				configs:     terraformPluginSDKExternalNameConfigs,
				implemented: sdkResources,
			},
		},
		"TerraformPluginFramework": {
			reason: "Every resource in the plugin framework table must be implemented by the plugin framework provider.",
			args: args{
				configs:     terraformPluginFrameworkExternalNameConfigs,
				implemented: frameworkResourceNames(datadogfw.New()),
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var got []string
			for resource := range tc.args.configs {
				if !tc.args.implemented[resource] {
					got = append(got, resource)
				}
			}
			slices.Sort(got)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("\n%s\nunimplemented resources: -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestResourceConfigurator(t *testing.T) {
	cases := map[string]struct {
		args string
		want bool
	}{
		"TerraformPluginSDKResource": {
			args: "datadog_logs_index",
			want: true,
		},
		"TerraformPluginFrameworkResource": {
			args: "datadog_team",
			want: true,
		},
		"UnconfiguredResource": {
			args: "datadog_dashboard",
			want: false,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := &config.Resource{Name: tc.args}
			resourceConfigurator()(r)
			got := r.ExternalName.GetIDFn != nil
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("resourceConfigurator() configured an external name: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestResourceList(t *testing.T) {
	cases := map[string]struct {
		args map[string]config.ExternalName
		want []string
	}{
		"NoResources": {
			args: map[string]config.ExternalName{},
			want: []string{},
		},
		"ExactMatch": {
			args: map[string]config.ExternalName{"datadog_team": config.IdentifierFromProvider},
			want: []string{"^datadog_team$"},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := resourceList(tc.args)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("resourceList(): -want, +got:\n%s", diff)
			}
		})
	}
}

func frameworkResourceNames(p fwprovider.Provider) map[string]bool {
	ctx := context.Background()
	var providerMeta fwprovider.MetadataResponse
	p.Metadata(ctx, fwprovider.MetadataRequest{}, &providerMeta)
	names := map[string]bool{}
	for _, newResource := range p.Resources(ctx) {
		var meta fwresource.MetadataResponse
		newResource().Metadata(ctx, fwresource.MetadataRequest{ProviderTypeName: providerMeta.TypeName}, &meta)
		names[meta.TypeName] = true
	}
	return names
}
