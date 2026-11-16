/*
Copyright 2026 Upbound Inc.
*/

package config

import (
	"context"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/go-cmp/cmp"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestNilUUIDPlaceholder(t *testing.T) {
	stateType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"id":    tftypes.String,
		"email": tftypes.String,
	}}

	type args struct {
		externalName string
	}
	type want struct {
		id    string
		empty bool
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"NoExternalName": {
			args: args{externalName: ""},
			want: want{id: "00000000-0000-0000-0000-000000000000", empty: true},
		},
		"ExternalName": {
			args: args{externalName: "573fb00e-e1be-11f1-8729-2ee1bd86e8ac"},
			want: want{id: "573fb00e-e1be-11f1-8729-2ee1bd86e8ac", empty: false},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := &config.Resource{Name: "datadog_service_account"}
			resourceConfigurator()(r)

			id, err := r.ExternalName.GetIDFn(context.Background(), tc.args.externalName, nil, nil)
			if err != nil {
				t.Fatalf("GetIDFn(...): unexpected error: %v", err)
			}
			// The state the upstream Read returns for a service account
			// that answers 404: the request state, id included.
			state := tftypes.NewValue(stateType, map[string]tftypes.Value{
				"id":    tftypes.NewValue(tftypes.String, id),
				"email": tftypes.NewValue(tftypes.String, "desired@example.com"),
			})
			empty, err := r.TerraformPluginFrameworkIsStateEmptyFn(context.Background(), state, rschema.Schema{})
			if err != nil {
				t.Fatalf("TerraformPluginFrameworkIsStateEmptyFn(...): unexpected error: %v", err)
			}

			if diff := cmp.Diff(tc.want, want{id: id, empty: empty}, cmp.AllowUnexported(want{})); diff != "" {
				t.Errorf("placeholder id and state emptiness: -want, +got:\n%s", diff)
			}
		})
	}
}
