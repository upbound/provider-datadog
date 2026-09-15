/*
Copyright 2026 Upbound Inc.
*/

package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestNotFoundDiagnostic(t *testing.T) {
	isNotFound := datadogExternalNameWithInjectedUUIDAndNotFoundDiagnostics("invalid policy_id", "404 Not Found").IsNotFoundDiagnosticFn

	type args struct {
		diags []*tfprotov6.Diagnostic
	}
	type want struct {
		notFound bool
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"MatchingSummary": {
			args: args{diags: []*tfprotov6.Diagnostic{{Severity: tfprotov6.DiagnosticSeverityError, Summary: "Error retrieving TagRule: 400 Bad Request: invalid policy_id"}}},
			want: want{notFound: true},
		},
		"MatchingDetail": {
			args: args{diags: []*tfprotov6.Diagnostic{{Severity: tfprotov6.DiagnosticSeverityError, Summary: "Error reading budget", Detail: "404 Not Found: {\"errors\":[\"Not found\"]}"}}},
			want: want{notFound: true},
		},
		"OtherError": {
			args: args{diags: []*tfprotov6.Diagnostic{{Severity: tfprotov6.DiagnosticSeverityError, Summary: "Error reading budget", Detail: "403 Forbidden"}}},
			want: want{notFound: false},
		},
		"MatchingWarning": {
			args: args{diags: []*tfprotov6.Diagnostic{{Severity: tfprotov6.DiagnosticSeverityWarning, Summary: "404 Not Found"}}},
			want: want{notFound: false},
		},
		"NoDiagnostics": {
			args: args{diags: nil},
			want: want{notFound: false},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := isNotFound(tc.args.diags)
			if diff := cmp.Diff(tc.want.notFound, got); diff != "" {
				t.Errorf("IsNotFoundDiagnosticFn(...): -want, +got:\n%s", diff)
			}
		})
	}
}
