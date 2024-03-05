/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/upbound/provider-datadog/config/api"
	"github.com/upbound/provider-datadog/config/apm"
	"github.com/upbound/provider-datadog/config/application"
	"github.com/upbound/provider-datadog/config/authn"
	"github.com/upbound/provider-datadog/config/child"
	"github.com/upbound/provider-datadog/config/cloud"
	"github.com/upbound/provider-datadog/config/dashboard"
	"github.com/upbound/provider-datadog/config/downtime"
	"github.com/upbound/provider-datadog/config/integration"
	"github.com/upbound/provider-datadog/config/ip"
	"github.com/upbound/provider-datadog/config/logs"
	"github.com/upbound/provider-datadog/config/metric"
	"github.com/upbound/provider-datadog/config/monitor"
	"github.com/upbound/provider-datadog/config/organization"
	"github.com/upbound/provider-datadog/config/powerpack"
	"github.com/upbound/provider-datadog/config/restriction"
	"github.com/upbound/provider-datadog/config/role"
	"github.com/upbound/provider-datadog/config/rum"
	"github.com/upbound/provider-datadog/config/securitymonitoring"
	"github.com/upbound/provider-datadog/config/sensitivedatascanner"
	"github.com/upbound/provider-datadog/config/service"
	"github.com/upbound/provider-datadog/config/spans"
	"github.com/upbound/provider-datadog/config/synthetics"
	"github.com/upbound/provider-datadog/config/team"
	"github.com/upbound/provider-datadog/config/user"
	"github.com/upbound/provider-datadog/config/webhook"
)

const (
	resourcePrefix = "datadog"
	modulePath     = "github.com/upbound/provider-datadog"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		api.Configure,
		apm.Configure,
		application.Configure,
		authn.Configure,
		child.Configure,
		cloud.Configure,
		dashboard.Configure,
		downtime.Configure,
		integration.Configure,
		ip.Configure,
		logs.Configure,
		metric.Configure,
		monitor.Configure,
		organization.Configure,
		powerpack.Configure,
		restriction.Configure,
		role.Configure,
		rum.Configure,
		securitymonitoring.Configure,
		sensitivedatascanner.Configure,
		service.Configure,
		spans.Configure,
		synthetics.Configure,
		team.Configure,
		user.Configure,
		webhook.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
