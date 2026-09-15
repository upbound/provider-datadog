/*
Copyright 2021 Upbound Inc.
*/

package cluster

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/cluster/api"
	"github.com/upbound/provider-datadog/config/cluster/apm"
	"github.com/upbound/provider-datadog/config/cluster/application"
	"github.com/upbound/provider-datadog/config/cluster/authn"
	"github.com/upbound/provider-datadog/config/cluster/child"
	"github.com/upbound/provider-datadog/config/cluster/cloud"
	"github.com/upbound/provider-datadog/config/cluster/dashboard"
	"github.com/upbound/provider-datadog/config/cluster/downtime"
	"github.com/upbound/provider-datadog/config/cluster/integration"
	"github.com/upbound/provider-datadog/config/cluster/ip"
	"github.com/upbound/provider-datadog/config/cluster/logs"
	"github.com/upbound/provider-datadog/config/cluster/metric"
	"github.com/upbound/provider-datadog/config/cluster/monitor"
	"github.com/upbound/provider-datadog/config/cluster/organization"
	"github.com/upbound/provider-datadog/config/cluster/powerpack"
	"github.com/upbound/provider-datadog/config/cluster/restriction"
	"github.com/upbound/provider-datadog/config/cluster/role"
	"github.com/upbound/provider-datadog/config/cluster/rum"
	"github.com/upbound/provider-datadog/config/cluster/securitymonitoring"
	"github.com/upbound/provider-datadog/config/cluster/sensitivedatascanner"
	"github.com/upbound/provider-datadog/config/cluster/service"
	"github.com/upbound/provider-datadog/config/cluster/spans"
	"github.com/upbound/provider-datadog/config/cluster/synthetics"
	"github.com/upbound/provider-datadog/config/cluster/team"
	"github.com/upbound/provider-datadog/config/cluster/user"
	"github.com/upbound/provider-datadog/config/cluster/webhook"
)

// Configurators lists the resource configurators applied to the cluster-scoped provider.
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
