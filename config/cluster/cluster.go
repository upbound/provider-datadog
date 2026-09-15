/*
Copyright 2021 Upbound Inc.
*/

package cluster

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/cluster/action"
	"github.com/upbound/provider-datadog/config/cluster/agentlessscanning"
	"github.com/upbound/provider-datadog/config/cluster/api"
	"github.com/upbound/provider-datadog/config/cluster/apicatalog"
	"github.com/upbound/provider-datadog/config/cluster/apm"
	"github.com/upbound/provider-datadog/config/cluster/appbuilder"
	"github.com/upbound/provider-datadog/config/cluster/application"
	"github.com/upbound/provider-datadog/config/cluster/appsec"
	"github.com/upbound/provider-datadog/config/cluster/authn"
	"github.com/upbound/provider-datadog/config/cluster/child"
	"github.com/upbound/provider-datadog/config/cluster/cloud"
	"github.com/upbound/provider-datadog/config/cluster/cloudcost"
	"github.com/upbound/provider-datadog/config/cluster/compliance"
	"github.com/upbound/provider-datadog/config/cluster/csmthreats"
	"github.com/upbound/provider-datadog/config/cluster/dashboard"
	"github.com/upbound/provider-datadog/config/cluster/data"
	"github.com/upbound/provider-datadog/config/cluster/downtime"
	"github.com/upbound/provider-datadog/config/cluster/governance"
	"github.com/upbound/provider-datadog/config/cluster/incident"
	"github.com/upbound/provider-datadog/config/cluster/integration"
	"github.com/upbound/provider-datadog/config/cluster/ip"
	"github.com/upbound/provider-datadog/config/cluster/logs"
	"github.com/upbound/provider-datadog/config/cluster/metric"
	"github.com/upbound/provider-datadog/config/cluster/monitor"
	"github.com/upbound/provider-datadog/config/cluster/observabilitypipeline"
	"github.com/upbound/provider-datadog/config/cluster/oncall"
	"github.com/upbound/provider-datadog/config/cluster/organization"
	"github.com/upbound/provider-datadog/config/cluster/powerpack"
	"github.com/upbound/provider-datadog/config/cluster/reference"
	"github.com/upbound/provider-datadog/config/cluster/restriction"
	"github.com/upbound/provider-datadog/config/cluster/role"
	"github.com/upbound/provider-datadog/config/cluster/rum"
	"github.com/upbound/provider-datadog/config/cluster/saml"
	"github.com/upbound/provider-datadog/config/cluster/securityfindings"
	"github.com/upbound/provider-datadog/config/cluster/securitymonitoring"
	"github.com/upbound/provider-datadog/config/cluster/sensitivedatascanner"
	"github.com/upbound/provider-datadog/config/cluster/service"
	"github.com/upbound/provider-datadog/config/cluster/spans"
	"github.com/upbound/provider-datadog/config/cluster/statuspage"
	"github.com/upbound/provider-datadog/config/cluster/synthetics"
	"github.com/upbound/provider-datadog/config/cluster/tag"
	"github.com/upbound/provider-datadog/config/cluster/team"
	"github.com/upbound/provider-datadog/config/cluster/user"
	"github.com/upbound/provider-datadog/config/cluster/webhook"
	"github.com/upbound/provider-datadog/config/cluster/workflowautomation"
)

// Configurators lists the resource configurators applied to the cluster-scoped provider.
var Configurators = []func(*ujconfig.Provider){
	action.Configure,
	agentlessscanning.Configure,
	api.Configure,
	apicatalog.Configure,
	apm.Configure,
	appbuilder.Configure,
	application.Configure,
	appsec.Configure,
	authn.Configure,
	child.Configure,
	cloud.Configure,
	cloudcost.Configure,
	compliance.Configure,
	csmthreats.Configure,
	dashboard.Configure,
	data.Configure,
	downtime.Configure,
	governance.Configure,
	incident.Configure,
	integration.Configure,
	ip.Configure,
	logs.Configure,
	metric.Configure,
	monitor.Configure,
	observabilitypipeline.Configure,
	oncall.Configure,
	organization.Configure,
	powerpack.Configure,
	reference.Configure,
	restriction.Configure,
	role.Configure,
	rum.Configure,
	saml.Configure,
	securityfindings.Configure,
	securitymonitoring.Configure,
	sensitivedatascanner.Configure,
	service.Configure,
	spans.Configure,
	statuspage.Configure,
	synthetics.Configure,
	tag.Configure,
	team.Configure,
	user.Configure,
	webhook.Configure,
	workflowautomation.Configure,
}
