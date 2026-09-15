/*
Copyright 2021 Upbound Inc.
*/

package namespaced

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/namespaced/action"
	"github.com/upbound/provider-datadog/config/namespaced/agentlessscanning"
	"github.com/upbound/provider-datadog/config/namespaced/api"
	"github.com/upbound/provider-datadog/config/namespaced/apicatalog"
	"github.com/upbound/provider-datadog/config/namespaced/apm"
	"github.com/upbound/provider-datadog/config/namespaced/appbuilder"
	"github.com/upbound/provider-datadog/config/namespaced/application"
	"github.com/upbound/provider-datadog/config/namespaced/appsec"
	"github.com/upbound/provider-datadog/config/namespaced/authn"
	"github.com/upbound/provider-datadog/config/namespaced/child"
	"github.com/upbound/provider-datadog/config/namespaced/cloud"
	"github.com/upbound/provider-datadog/config/namespaced/cloudcost"
	"github.com/upbound/provider-datadog/config/namespaced/compliance"
	"github.com/upbound/provider-datadog/config/namespaced/csmthreats"
	"github.com/upbound/provider-datadog/config/namespaced/dashboard"
	"github.com/upbound/provider-datadog/config/namespaced/data"
	"github.com/upbound/provider-datadog/config/namespaced/downtime"
	"github.com/upbound/provider-datadog/config/namespaced/governance"
	"github.com/upbound/provider-datadog/config/namespaced/incident"
	"github.com/upbound/provider-datadog/config/namespaced/integration"
	"github.com/upbound/provider-datadog/config/namespaced/ip"
	"github.com/upbound/provider-datadog/config/namespaced/logs"
	"github.com/upbound/provider-datadog/config/namespaced/metric"
	"github.com/upbound/provider-datadog/config/namespaced/monitor"
	"github.com/upbound/provider-datadog/config/namespaced/observabilitypipeline"
	"github.com/upbound/provider-datadog/config/namespaced/oncall"
	"github.com/upbound/provider-datadog/config/namespaced/organization"
	"github.com/upbound/provider-datadog/config/namespaced/powerpack"
	"github.com/upbound/provider-datadog/config/namespaced/reference"
	"github.com/upbound/provider-datadog/config/namespaced/restriction"
	"github.com/upbound/provider-datadog/config/namespaced/role"
	"github.com/upbound/provider-datadog/config/namespaced/rum"
	"github.com/upbound/provider-datadog/config/namespaced/saml"
	"github.com/upbound/provider-datadog/config/namespaced/securityfindings"
	"github.com/upbound/provider-datadog/config/namespaced/securitymonitoring"
	"github.com/upbound/provider-datadog/config/namespaced/sensitivedatascanner"
	"github.com/upbound/provider-datadog/config/namespaced/service"
	"github.com/upbound/provider-datadog/config/namespaced/spans"
	"github.com/upbound/provider-datadog/config/namespaced/statuspage"
	"github.com/upbound/provider-datadog/config/namespaced/synthetics"
	"github.com/upbound/provider-datadog/config/namespaced/tag"
	"github.com/upbound/provider-datadog/config/namespaced/team"
	"github.com/upbound/provider-datadog/config/namespaced/user"
	"github.com/upbound/provider-datadog/config/namespaced/webhook"
	"github.com/upbound/provider-datadog/config/namespaced/workflowautomation"
)

// Configurators lists the resource configurators applied to the namespaced provider.
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
