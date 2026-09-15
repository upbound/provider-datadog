/*
Copyright 2021 Upbound Inc.
*/

package namespaced

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/namespaced/api"
	"github.com/upbound/provider-datadog/config/namespaced/apm"
	"github.com/upbound/provider-datadog/config/namespaced/application"
	"github.com/upbound/provider-datadog/config/namespaced/authn"
	"github.com/upbound/provider-datadog/config/namespaced/child"
	"github.com/upbound/provider-datadog/config/namespaced/cloud"
	"github.com/upbound/provider-datadog/config/namespaced/dashboard"
	"github.com/upbound/provider-datadog/config/namespaced/downtime"
	"github.com/upbound/provider-datadog/config/namespaced/integration"
	"github.com/upbound/provider-datadog/config/namespaced/ip"
	"github.com/upbound/provider-datadog/config/namespaced/logs"
	"github.com/upbound/provider-datadog/config/namespaced/metric"
	"github.com/upbound/provider-datadog/config/namespaced/monitor"
	"github.com/upbound/provider-datadog/config/namespaced/organization"
	"github.com/upbound/provider-datadog/config/namespaced/powerpack"
	"github.com/upbound/provider-datadog/config/namespaced/restriction"
	"github.com/upbound/provider-datadog/config/namespaced/role"
	"github.com/upbound/provider-datadog/config/namespaced/rum"
	"github.com/upbound/provider-datadog/config/namespaced/securitymonitoring"
	"github.com/upbound/provider-datadog/config/namespaced/sensitivedatascanner"
	"github.com/upbound/provider-datadog/config/namespaced/service"
	"github.com/upbound/provider-datadog/config/namespaced/spans"
	"github.com/upbound/provider-datadog/config/namespaced/synthetics"
	"github.com/upbound/provider-datadog/config/namespaced/team"
	"github.com/upbound/provider-datadog/config/namespaced/user"
	"github.com/upbound/provider-datadog/config/namespaced/webhook"
)

// Configurators lists the resource configurators applied to the namespaced provider.
var Configurators = []func(*ujconfig.Provider){
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
}
