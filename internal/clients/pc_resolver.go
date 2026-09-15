/*
Copyright 2026 Upbound Inc.
*/

package clients

import (
	"context"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/apis/configuration/v1alpha1"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ReconciliationPolicy returns the reconciliation policy configured on the
// ProviderConfig referenced by the given managed resource, or nil when the
// ProviderConfig sets none.
func ReconciliationPolicy(ctx context.Context, c client.Client, mg resource.Managed) (*v1alpha1.ReconciliationPolicy, error) {
	pcSpec, err := resolveProviderConfig(ctx, c, mg)
	if err != nil {
		return nil, errors.Wrap(err, errResolveProviderConfig)
	}
	return pcSpec.ReconciliationPolicy, nil
}
