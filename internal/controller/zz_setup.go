// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	connection "github.com/upbound/provider-datadog/internal/controller/action/connection"
	scanningawsscanoptions "github.com/upbound/provider-datadog/internal/controller/agentless/scanningawsscanoptions"
	scanninggcpscanoptions "github.com/upbound/provider-datadog/internal/controller/agentless/scanninggcpscanoptions"
	retentionfilter "github.com/upbound/provider-datadog/internal/controller/apm/retentionfilter"
	retentionfilterorder "github.com/upbound/provider-datadog/internal/controller/apm/retentionfilterorder"
	builderapp "github.com/upbound/provider-datadog/internal/controller/app/builderapp"
	keyregistration "github.com/upbound/provider-datadog/internal/controller/app/keyregistration"
	wafcustomrule "github.com/upbound/provider-datadog/internal/controller/appsec/wafcustomrule"
	wafexclusionfilter "github.com/upbound/provider-datadog/internal/controller/appsec/wafexclusionfilter"
	curconfig "github.com/upbound/provider-datadog/internal/controller/aws/curconfig"
	ucconfig "github.com/upbound/provider-datadog/internal/controller/azure/ucconfig"
	configurationrule "github.com/upbound/provider-datadog/internal/controller/cloud/configurationrule"
	inventorysyncconfig "github.com/upbound/provider-datadog/internal/controller/cloud/inventorysyncconfig"
	workloadsecurityagentrule "github.com/upbound/provider-datadog/internal/controller/cloud/workloadsecurityagentrule"
	customframework "github.com/upbound/provider-datadog/internal/controller/compliance/customframework"
	resourceevaluationfilter "github.com/upbound/provider-datadog/internal/controller/compliance/resourceevaluationfilter"
	budget "github.com/upbound/provider-datadog/internal/controller/cost/budget"
	threatsagentrule "github.com/upbound/provider-datadog/internal/controller/csm/threatsagentrule"
	threatspolicy "github.com/upbound/provider-datadog/internal/controller/csm/threatspolicy"
	allocationrule "github.com/upbound/provider-datadog/internal/controller/custom/allocationrule"
	allocationrules "github.com/upbound/provider-datadog/internal/controller/custom/allocationrules"
	apikey "github.com/upbound/provider-datadog/internal/controller/datadog/apikey"
	appkey "github.com/upbound/provider-datadog/internal/controller/datadog/appkey"
	authnmapping "github.com/upbound/provider-datadog/internal/controller/datadog/authnmapping"
	childorganization "github.com/upbound/provider-datadog/internal/controller/datadog/childorganization"
	dashboard "github.com/upbound/provider-datadog/internal/controller/datadog/dashboard"
	dashboardjson "github.com/upbound/provider-datadog/internal/controller/datadog/dashboardjson"
	dashboardlistentry "github.com/upbound/provider-datadog/internal/controller/datadog/dashboardlistentry"
	dataset "github.com/upbound/provider-datadog/internal/controller/datadog/dataset"
	datastore "github.com/upbound/provider-datadog/internal/controller/datadog/datastore"
	downtime "github.com/upbound/provider-datadog/internal/controller/datadog/downtime"
	downtimeschedule "github.com/upbound/provider-datadog/internal/controller/datadog/downtimeschedule"
	ipallowlist "github.com/upbound/provider-datadog/internal/controller/datadog/ipallowlist"
	monitor "github.com/upbound/provider-datadog/internal/controller/datadog/monitor"
	monitorconfigpolicy "github.com/upbound/provider-datadog/internal/controller/datadog/monitorconfigpolicy"
	monitorjson "github.com/upbound/provider-datadog/internal/controller/datadog/monitorjson"
	organizationsettings "github.com/upbound/provider-datadog/internal/controller/datadog/organizationsettings"
	powerpack "github.com/upbound/provider-datadog/internal/controller/datadog/powerpack"
	restrictionpolicy "github.com/upbound/provider-datadog/internal/controller/datadog/restrictionpolicy"
	role "github.com/upbound/provider-datadog/internal/controller/datadog/role"
	rumapplication "github.com/upbound/provider-datadog/internal/controller/datadog/rumapplication"
	serviceaccount "github.com/upbound/provider-datadog/internal/controller/datadog/serviceaccount"
	serviceaccountapplicationkey "github.com/upbound/provider-datadog/internal/controller/datadog/serviceaccountapplicationkey"
	servicedefinitionyaml "github.com/upbound/provider-datadog/internal/controller/datadog/servicedefinitionyaml"
	servicelevelobjective "github.com/upbound/provider-datadog/internal/controller/datadog/servicelevelobjective"
	slocorrection "github.com/upbound/provider-datadog/internal/controller/datadog/slocorrection"
	spansmetric "github.com/upbound/provider-datadog/internal/controller/datadog/spansmetric"
	team "github.com/upbound/provider-datadog/internal/controller/datadog/team"
	teamlink "github.com/upbound/provider-datadog/internal/controller/datadog/teamlink"
	teammembership "github.com/upbound/provider-datadog/internal/controller/datadog/teammembership"
	teampermissionsetting "github.com/upbound/provider-datadog/internal/controller/datadog/teampermissionsetting"
	user "github.com/upbound/provider-datadog/internal/controller/datadog/user"
	webhook "github.com/upbound/provider-datadog/internal/controller/datadog/webhook"
	webhookcustomvariable "github.com/upbound/provider-datadog/internal/controller/datadog/webhookcustomvariable"
	item "github.com/upbound/provider-datadog/internal/controller/datastore/item"
	gate "github.com/upbound/provider-datadog/internal/controller/deployment/gate"
	allowlist "github.com/upbound/provider-datadog/internal/controller/domain/allowlist"
	ucconfiggcp "github.com/upbound/provider-datadog/internal/controller/gcp/ucconfig"
	incidenttype "github.com/upbound/provider-datadog/internal/controller/incident/incidenttype"
	notificationrule "github.com/upbound/provider-datadog/internal/controller/incident/notificationrule"
	notificationtemplate "github.com/upbound/provider-datadog/internal/controller/incident/notificationtemplate"
	aws "github.com/upbound/provider-datadog/internal/controller/integration/aws"
	awsaccount "github.com/upbound/provider-datadog/internal/controller/integration/awsaccount"
	awseventbridge "github.com/upbound/provider-datadog/internal/controller/integration/awseventbridge"
	awsexternalid "github.com/upbound/provider-datadog/internal/controller/integration/awsexternalid"
	awslambdaarn "github.com/upbound/provider-datadog/internal/controller/integration/awslambdaarn"
	awslogcollection "github.com/upbound/provider-datadog/internal/controller/integration/awslogcollection"
	awstagfilter "github.com/upbound/provider-datadog/internal/controller/integration/awstagfilter"
	azure "github.com/upbound/provider-datadog/internal/controller/integration/azure"
	cloudflareaccount "github.com/upbound/provider-datadog/internal/controller/integration/cloudflareaccount"
	confluentaccount "github.com/upbound/provider-datadog/internal/controller/integration/confluentaccount"
	confluentresource "github.com/upbound/provider-datadog/internal/controller/integration/confluentresource"
	fastlyaccount "github.com/upbound/provider-datadog/internal/controller/integration/fastlyaccount"
	fastlyservice "github.com/upbound/provider-datadog/internal/controller/integration/fastlyservice"
	gcp "github.com/upbound/provider-datadog/internal/controller/integration/gcp"
	gcpsts "github.com/upbound/provider-datadog/internal/controller/integration/gcpsts"
	msteamstenantbasedhandle "github.com/upbound/provider-datadog/internal/controller/integration/msteamstenantbasedhandle"
	msteamsworkflowswebhookhandle "github.com/upbound/provider-datadog/internal/controller/integration/msteamsworkflowswebhookhandle"
	opsgenieserviceobject "github.com/upbound/provider-datadog/internal/controller/integration/opsgenieserviceobject"
	pagerduty "github.com/upbound/provider-datadog/internal/controller/integration/pagerduty"
	pagerdutyserviceobject "github.com/upbound/provider-datadog/internal/controller/integration/pagerdutyserviceobject"
	slackchannel "github.com/upbound/provider-datadog/internal/controller/integration/slackchannel"
	archive "github.com/upbound/provider-datadog/internal/controller/logs/archive"
	archiveorder "github.com/upbound/provider-datadog/internal/controller/logs/archiveorder"
	customdestination "github.com/upbound/provider-datadog/internal/controller/logs/customdestination"
	custompipeline "github.com/upbound/provider-datadog/internal/controller/logs/custompipeline"
	index "github.com/upbound/provider-datadog/internal/controller/logs/index"
	indexorder "github.com/upbound/provider-datadog/internal/controller/logs/indexorder"
	integrationpipeline "github.com/upbound/provider-datadog/internal/controller/logs/integrationpipeline"
	metric "github.com/upbound/provider-datadog/internal/controller/logs/metric"
	pipelineorder "github.com/upbound/provider-datadog/internal/controller/logs/pipelineorder"
	restrictionquery "github.com/upbound/provider-datadog/internal/controller/logs/restrictionquery"
	metadata "github.com/upbound/provider-datadog/internal/controller/metric/metadata"
	tagconfiguration "github.com/upbound/provider-datadog/internal/controller/metric/tagconfiguration"
	notificationrulemonitor "github.com/upbound/provider-datadog/internal/controller/monitor/notificationrule"
	pipeline "github.com/upbound/provider-datadog/internal/controller/observability/pipeline"
	callescalationpolicy "github.com/upbound/provider-datadog/internal/controller/on/callescalationpolicy"
	callschedule "github.com/upbound/provider-datadog/internal/controller/on/callschedule"
	callteamroutingrules "github.com/upbound/provider-datadog/internal/controller/on/callteamroutingrules"
	callusernotificationchannel "github.com/upbound/provider-datadog/internal/controller/on/callusernotificationchannel"
	callusernotificationrule "github.com/upbound/provider-datadog/internal/controller/on/callusernotificationrule"
	api "github.com/upbound/provider-datadog/internal/controller/openapi/api"
	connectionorg "github.com/upbound/provider-datadog/internal/controller/org/connection"
	providerconfig "github.com/upbound/provider-datadog/internal/controller/providerconfig"
	table "github.com/upbound/provider-datadog/internal/controller/reference/table"
	metricrum "github.com/upbound/provider-datadog/internal/controller/rum/metric"
	retentionfilterrum "github.com/upbound/provider-datadog/internal/controller/rum/retentionfilter"
	retentionfiltersorder "github.com/upbound/provider-datadog/internal/controller/rum/retentionfiltersorder"
	monitoringcriticalasset "github.com/upbound/provider-datadog/internal/controller/security/monitoringcriticalasset"
	monitoringrulejson "github.com/upbound/provider-datadog/internal/controller/security/monitoringrulejson"
	monitoringsuppression "github.com/upbound/provider-datadog/internal/controller/security/monitoringsuppression"
	notificationrulesecurity "github.com/upbound/provider-datadog/internal/controller/security/notificationrule"
	defaultrule "github.com/upbound/provider-datadog/internal/controller/securitymonitoring/defaultrule"
	filter "github.com/upbound/provider-datadog/internal/controller/securitymonitoring/filter"
	rule "github.com/upbound/provider-datadog/internal/controller/securitymonitoring/rule"
	group "github.com/upbound/provider-datadog/internal/controller/sensitivedatascanner/group"
	grouporder "github.com/upbound/provider-datadog/internal/controller/sensitivedatascanner/grouporder"
	rulesensitivedatascanner "github.com/upbound/provider-datadog/internal/controller/sensitivedatascanner/rule"
	catalog "github.com/upbound/provider-datadog/internal/controller/software/catalog"
	concurrencycap "github.com/upbound/provider-datadog/internal/controller/synthetics/concurrencycap"
	globalvariable "github.com/upbound/provider-datadog/internal/controller/synthetics/globalvariable"
	privatelocation "github.com/upbound/provider-datadog/internal/controller/synthetics/privatelocation"
	suite "github.com/upbound/provider-datadog/internal/controller/synthetics/suite"
	test "github.com/upbound/provider-datadog/internal/controller/synthetics/test"
	pipelineruleset "github.com/upbound/provider-datadog/internal/controller/tag/pipelineruleset"
	pipelinerulesets "github.com/upbound/provider-datadog/internal/controller/tag/pipelinerulesets"
	hierarchylinks "github.com/upbound/provider-datadog/internal/controller/team/hierarchylinks"
	notificationruleteam "github.com/upbound/provider-datadog/internal/controller/team/notificationrule"
	roleuser "github.com/upbound/provider-datadog/internal/controller/user/role"
	automation "github.com/upbound/provider-datadog/internal/controller/workflow/automation"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connection.Setup,
		scanningawsscanoptions.Setup,
		scanninggcpscanoptions.Setup,
		retentionfilter.Setup,
		retentionfilterorder.Setup,
		builderapp.Setup,
		keyregistration.Setup,
		wafcustomrule.Setup,
		wafexclusionfilter.Setup,
		curconfig.Setup,
		ucconfig.Setup,
		configurationrule.Setup,
		inventorysyncconfig.Setup,
		workloadsecurityagentrule.Setup,
		customframework.Setup,
		resourceevaluationfilter.Setup,
		budget.Setup,
		threatsagentrule.Setup,
		threatspolicy.Setup,
		allocationrule.Setup,
		allocationrules.Setup,
		apikey.Setup,
		appkey.Setup,
		authnmapping.Setup,
		childorganization.Setup,
		dashboard.Setup,
		dashboardjson.Setup,
		dashboardlistentry.Setup,
		dataset.Setup,
		datastore.Setup,
		downtime.Setup,
		downtimeschedule.Setup,
		ipallowlist.Setup,
		monitor.Setup,
		monitorconfigpolicy.Setup,
		monitorjson.Setup,
		organizationsettings.Setup,
		powerpack.Setup,
		restrictionpolicy.Setup,
		role.Setup,
		rumapplication.Setup,
		serviceaccount.Setup,
		serviceaccountapplicationkey.Setup,
		servicedefinitionyaml.Setup,
		servicelevelobjective.Setup,
		slocorrection.Setup,
		spansmetric.Setup,
		team.Setup,
		teamlink.Setup,
		teammembership.Setup,
		teampermissionsetting.Setup,
		user.Setup,
		webhook.Setup,
		webhookcustomvariable.Setup,
		item.Setup,
		gate.Setup,
		allowlist.Setup,
		ucconfiggcp.Setup,
		incidenttype.Setup,
		notificationrule.Setup,
		notificationtemplate.Setup,
		aws.Setup,
		awsaccount.Setup,
		awseventbridge.Setup,
		awsexternalid.Setup,
		awslambdaarn.Setup,
		awslogcollection.Setup,
		awstagfilter.Setup,
		azure.Setup,
		cloudflareaccount.Setup,
		confluentaccount.Setup,
		confluentresource.Setup,
		fastlyaccount.Setup,
		fastlyservice.Setup,
		gcp.Setup,
		gcpsts.Setup,
		msteamstenantbasedhandle.Setup,
		msteamsworkflowswebhookhandle.Setup,
		opsgenieserviceobject.Setup,
		pagerduty.Setup,
		pagerdutyserviceobject.Setup,
		slackchannel.Setup,
		archive.Setup,
		archiveorder.Setup,
		customdestination.Setup,
		custompipeline.Setup,
		index.Setup,
		indexorder.Setup,
		integrationpipeline.Setup,
		metric.Setup,
		pipelineorder.Setup,
		restrictionquery.Setup,
		metadata.Setup,
		tagconfiguration.Setup,
		notificationrulemonitor.Setup,
		pipeline.Setup,
		callescalationpolicy.Setup,
		callschedule.Setup,
		callteamroutingrules.Setup,
		callusernotificationchannel.Setup,
		callusernotificationrule.Setup,
		api.Setup,
		connectionorg.Setup,
		providerconfig.Setup,
		table.Setup,
		metricrum.Setup,
		retentionfilterrum.Setup,
		retentionfiltersorder.Setup,
		monitoringcriticalasset.Setup,
		monitoringrulejson.Setup,
		monitoringsuppression.Setup,
		notificationrulesecurity.Setup,
		defaultrule.Setup,
		filter.Setup,
		rule.Setup,
		group.Setup,
		grouporder.Setup,
		rulesensitivedatascanner.Setup,
		catalog.Setup,
		concurrencycap.Setup,
		globalvariable.Setup,
		privatelocation.Setup,
		suite.Setup,
		test.Setup,
		pipelineruleset.Setup,
		pipelinerulesets.Setup,
		hierarchylinks.Setup,
		notificationruleteam.Setup,
		roleuser.Setup,
		automation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
