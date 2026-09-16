/*
Copyright 2026 Upbound Inc.
*/

package clients

import (
	"maps"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// keyPair is the api_key/app_key pair most cases start from, as the JSON
// members and as the configuration they produce.
const keyPair = `"api_key": "k", "app_key": "a"`

func withKeyPair(cfg map[string]any) map[string]any {
	out := map[string]any{"api_key": "k", "app_key": "a"}
	maps.Copy(out, cfg)
	return out
}

func TestProviderConfiguration(t *testing.T) {
	type args struct {
		credentials string
	}
	type want struct {
		cfg map[string]any
		err string
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"KeyPair": {
			args: args{credentials: `{` + keyPair + `, "api_url": "https://api.datadoghq.eu"}`},
			want: want{cfg: withKeyPair(map[string]any{"api_url": "https://api.datadoghq.eu"})},
		},
		"BearerToken": {
			args: args{credentials: `{"bearer_token": "ddpat_x"}`},
			want: want{cfg: map[string]any{"bearer_token": "ddpat_x"}},
		},
		"RetryKnobsAsNumbers": {
			args: args{credentials: `{` + keyPair + `, "http_client_retry_enabled": "true", "http_client_retry_max_retries": 5, "http_client_retry_jitter": 2}`},
			want: want{cfg: withKeyPair(map[string]any{"http_client_retry_enabled": "true", "http_client_retry_max_retries": float64(5), "http_client_retry_jitter": float64(2)})},
		},
		"RetryKnobsAsStrings": {
			args: args{credentials: `{` + keyPair + `, "http_client_retry_timeout": "30", "http_client_retry_backoff_base": "3", "http_client_retry_backoff_multiplier": "4"}`},
			want: want{cfg: withKeyPair(map[string]any{"http_client_retry_timeout": float64(30), "http_client_retry_backoff_base": float64(3), "http_client_retry_backoff_multiplier": float64(4)})},
		},
		"FlagsAsBooleans": {
			args: args{credentials: `{"validate": false, "http_client_retry_enabled": true}`},
			want: want{cfg: map[string]any{"validate": "false", "http_client_retry_enabled": "true"}},
		},
		"IgnoreTagKeys": {
			args: args{credentials: `{` + keyPair + `, "ignore_tag_keys": ["team", "env"]}`},
			want: want{cfg: withKeyPair(map[string]any{"ignore_tag_keys": []any{"team", "env"}})},
		},
		"EmptyIgnoreTagKeysOmitted": {
			args: args{credentials: `{` + keyPair + `, "ignore_tag_keys": []}`},
			want: want{cfg: withKeyPair(nil)},
		},
		"UnknownKeysIgnored": {
			args: args{credentials: `{` + keyPair + `, "cloud_provider_type": "aws", "comment": 1}`},
			want: want{cfg: withKeyPair(nil)},
		},
		"NotJSON": {
			args: args{credentials: `api_key=k`},
			want: want{err: "cannot unmarshal datadog credentials as JSON: invalid character 'a' looking for beginning of value"},
		},
		"NonIntegerNumber": {
			args: args{credentials: `{"http_client_retry_jitter": 1.5}`},
			want: want{err: "cannot unmarshal datadog credentials as JSON: 1.5 is not an integer"},
		},
		"NonNumericString": {
			args: args{credentials: `{"http_client_retry_max_retries": "many"}`},
			want: want{err: `cannot unmarshal datadog credentials as JSON: "many" is not an integer`},
		},
		"FlagNotABoolean": {
			args: args{credentials: `{"validate": "maybe"}`},
			want: want{err: `cannot unmarshal datadog credentials as JSON: "maybe" is not a boolean`},
		},
		"StringKeyNotAString": {
			args: args{credentials: `{"api_key": 1}`},
			want: want{err: "cannot unmarshal datadog credentials as JSON: json: cannot unmarshal number into Go struct field credentials.api_key of type string"},
		},
		"IgnoreTagKeysNotAnArray": {
			args: args{credentials: `{"ignore_tag_keys": "team,env"}`},
			want: want{err: "cannot unmarshal datadog credentials as JSON: json: cannot unmarshal string into Go struct field credentials.ignore_tag_keys of type []string"},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := providerConfiguration([]byte(tc.args.credentials))
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}
			if diff := cmp.Diff(tc.want.err, gotErr); diff != "" {
				t.Errorf("providerConfiguration(...) error: -want, +got:\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.cfg, got); diff != "" {
				t.Errorf("providerConfiguration(...): -want, +got:\n%s", diff)
			}
		})
	}
}
