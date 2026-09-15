/*
Copyright 2026 Upbound Inc.
*/

// Package roundtrip holds the API round-trip tests: every managed resource
// type must survive a serialization cycle, and every multi-version kind must
// survive spoke -> hub -> spoke and hub -> spoke -> hub conversions.
package roundtrip

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/apitesting/roundtrip"
	"github.com/terraform-providers/terraform-provider-datadog/datadog"
	"github.com/terraform-providers/terraform-provider-datadog/datadog/fwprovider"
	"k8s.io/apimachinery/pkg/runtime"

	clusterapis "github.com/upbound/provider-datadog/apis/cluster"
	namespacedapis "github.com/upbound/provider-datadog/apis/namespaced"
	"github.com/upbound/provider-datadog/config"
)

func TestRoundTrip(t *testing.T) {
	testScheme := runtime.NewScheme()
	if err := clusterapis.AddToScheme(testScheme); err != nil {
		t.Fatalf("cluster-scoped apis AddToScheme: %s", err)
	}
	if err := namespacedapis.AddToScheme(testScheme); err != nil {
		t.Fatalf("namespaced apis AddToScheme: %s", err)
	}

	sdkProvider := datadog.Provider()
	fwProvider := fwprovider.New()
	provider, err := config.GetProvider(sdkProvider, fwProvider, false)
	if err != nil {
		t.Fatalf("GetProvider: %s", err)
	}
	providerNamespaced, err := config.GetProviderNamespaced(sdkProvider, fwProvider, false)
	if err != nil {
		t.Fatalf("GetProviderNamespaced: %s", err)
	}

	rt, err := roundtrip.NewRoundTripTest(provider, providerNamespaced, testScheme,
		roundtrip.WithFuzzerConfig(
			roundtrip.FuzzerIterations(10),
			roundtrip.FuzzerNilChance(0)),
		roundtrip.WithFuzzerConfig(
			roundtrip.FuzzerIterations(30),
			roundtrip.FuzzerNilChance(0.3)),
		roundtrip.WithComparisonOptions(
			roundtrip.EquateEmptyAndSingleZeroSlice(),
			roundtrip.EquateNilAndZeroValuePtr(),
		),
	)
	if err != nil {
		t.Fatalf("NewRoundTripTest: %s", err)
	}

	t.Run("TestSerializationRoundtrip", func(t *testing.T) {
		rt.TestSerializationRoundtrip(t)
	})

	t.Run("TestConversionRoundtrip", func(t *testing.T) {
		rt.TestConversionRoundtrip(t)
	})
}
