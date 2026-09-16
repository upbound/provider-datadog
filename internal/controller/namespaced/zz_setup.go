/*
Copyright 2026 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	connection "github.com/upbound/provider-datadog/internal/controller/namespaced/action/connection"
	awsscanoptions "github.com/upbound/provider-datadog/internal/controller/namespaced/agentlessscanning/awsscanoptions"
	azurescanoptions "github.com/upbound/provider-datadog/internal/controller/namespaced/agentlessscanning/azurescanoptions"
	gcpscanoptions "github.com/upbound/provider-datadog/internal/controller/namespaced/agentlessscanning/gcpscanoptions"
	softwarecatalog "github.com/upbound/provider-datadog/internal/controller/namespaced/apicatalog/softwarecatalog"
	retentionfilter "github.com/upbound/provider-datadog/internal/controller/namespaced/apm/retentionfilter"
	retentionfilterorder "github.com/upbound/provider-datadog/internal/controller/namespaced/apm/retentionfilterorder"
	app "github.com/upbound/provider-datadog/internal/controller/namespaced/appbuilder/app"
	wafcustomrule "github.com/upbound/provider-datadog/internal/controller/namespaced/appsec/wafcustomrule"
	wafexclusionfilter "github.com/upbound/provider-datadog/internal/controller/namespaced/appsec/wafexclusionfilter"
	cloudinventorysyncconfig "github.com/upbound/provider-datadog/internal/controller/namespaced/cloud/cloudinventorysyncconfig"
	configurationrule "github.com/upbound/provider-datadog/internal/controller/namespaced/cloud/configurationrule"
	workloadsecurityagentrule "github.com/upbound/provider-datadog/internal/controller/namespaced/cloud/workloadsecurityagentrule"
	allocationrule "github.com/upbound/provider-datadog/internal/controller/namespaced/cloudcost/allocationrule"
	awscurconfig "github.com/upbound/provider-datadog/internal/controller/namespaced/cloudcost/awscurconfig"
	azureucconfig "github.com/upbound/provider-datadog/internal/controller/namespaced/cloudcost/azureucconfig"
	costbudget "github.com/upbound/provider-datadog/internal/controller/namespaced/cloudcost/costbudget"
	costcustomforecast "github.com/upbound/provider-datadog/internal/controller/namespaced/cloudcost/costcustomforecast"
	customallocationrules "github.com/upbound/provider-datadog/internal/controller/namespaced/cloudcost/customallocationrules"
	gcpucconfig "github.com/upbound/provider-datadog/internal/controller/namespaced/cloudcost/gcpucconfig"
	customframework "github.com/upbound/provider-datadog/internal/controller/namespaced/compliance/customframework"
	resourceevaluationfilter "github.com/upbound/provider-datadog/internal/controller/namespaced/compliance/resourceevaluationfilter"
	agentrule "github.com/upbound/provider-datadog/internal/controller/namespaced/csmthreats/agentrule"
	policy "github.com/upbound/provider-datadog/internal/controller/namespaced/csmthreats/policy"
	dataset "github.com/upbound/provider-datadog/internal/controller/namespaced/data/dataset"
	datastore "github.com/upbound/provider-datadog/internal/controller/namespaced/data/datastore"
	datastoreitem "github.com/upbound/provider-datadog/internal/controller/namespaced/data/datastoreitem"
	apikey "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/apikey"
	appkey "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/appkey"
	appkeyregistration "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/appkeyregistration"
	authnmapping "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/authnmapping"
	childorganization "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/childorganization"
	dashboardjson "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/dashboardjson"
	dashboardlist "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/dashboardlist"
	domainallowlist "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/domainallowlist"
	downtime "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/downtime"
	downtimeschedule "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/downtimeschedule"
	ipallowlist "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/ipallowlist"
	monitor "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/monitor"
	monitorconfigpolicy "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/monitorconfigpolicy"
	monitorjson "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/monitorjson"
	monitornotificationrule "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/monitornotificationrule"
	organizationsettings "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/organizationsettings"
	orgconnection "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/orgconnection"
	orggroup "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/orggroup"
	orggroupmembership "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/orggroupmembership"
	orggrouppolicy "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/orggrouppolicy"
	orggrouppolicyoverride "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/orggrouppolicyoverride"
	powerpack "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/powerpack"
	powerpackv2 "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/powerpackv2"
	restrictionpolicy "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/restrictionpolicy"
	role "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/role"
	rumapplication "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/rumapplication"
	rumexclusionfilter "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/rumexclusionfilter"
	rummetric "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/rummetric"
	rumretentionfilter "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/rumretentionfilter"
	rumretentionquota "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/rumretentionquota"
	secureembeddashboard "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/secureembeddashboard"
	serviceaccesstoken "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/serviceaccesstoken"
	serviceaccount "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/serviceaccount"
	serviceaccountapplicationkey "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/serviceaccountapplicationkey"
	servicedefinitionyaml "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/servicedefinitionyaml"
	servicelevelobjective "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/servicelevelobjective"
	slocorrection "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/slocorrection"
	spansmetric "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/spansmetric"
	team "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/team"
	teamconnection "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/teamconnection"
	teamhierarchylinks "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/teamhierarchylinks"
	teamlink "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/teamlink"
	teammembership "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/teammembership"
	teamnotificationrule "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/teamnotificationrule"
	teampermissionsetting "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/teampermissionsetting"
	teamsync "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/teamsync"
	user "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/user"
	userrole "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/userrole"
	webhook "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/webhook"
	webhookcustomvariable "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/webhookcustomvariable"
	webhookoauth2clientcredentials "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/webhookoauth2clientcredentials"
	control "github.com/upbound/provider-datadog/internal/controller/namespaced/governance/control"
	incidenttype "github.com/upbound/provider-datadog/internal/controller/namespaced/incident/incidenttype"
	notificationrule "github.com/upbound/provider-datadog/internal/controller/namespaced/incident/notificationrule"
	notificationtemplate "github.com/upbound/provider-datadog/internal/controller/namespaced/incident/notificationtemplate"
	userdefinedfield "github.com/upbound/provider-datadog/internal/controller/namespaced/incident/userdefinedfield"
	userdefinedrole "github.com/upbound/provider-datadog/internal/controller/namespaced/incident/userdefinedrole"
	awsaccount "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/awsaccount"
	awsaccountccmconfig "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/awsaccountccmconfig"
	awseventbridge "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/awseventbridge"
	awsexternalid "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/awsexternalid"
	azure "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/azure"
	cloudflareaccount "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/cloudflareaccount"
	confluentaccount "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/confluentaccount"
	confluentresource "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/confluentresource"
	fastlyaccount "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/fastlyaccount"
	fastlyservice "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/fastlyservice"
	gcp "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/gcp"
	gcpsts "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/gcpsts"
	msteamstenantbasedhandle "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/msteamstenantbasedhandle"
	msteamsworkflowswebhookhandle "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/msteamsworkflowswebhookhandle"
	opsgenieserviceobject "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/opsgenieserviceobject"
	pagerduty "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/pagerduty"
	pagerdutyserviceobject "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/pagerdutyserviceobject"
	slackchannel "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/slackchannel"
	archive "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/archive"
	archiveorder "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/archiveorder"
	customdestination "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/customdestination"
	custompipeline "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/custompipeline"
	index "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/index"
	indexorder "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/indexorder"
	integrationpipeline "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/integrationpipeline"
	metric "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/metric"
	pipelineorder "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/pipelineorder"
	restrictionquery "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/restrictionquery"
	metadata "github.com/upbound/provider-datadog/internal/controller/namespaced/metric/metadata"
	tagconfiguration "github.com/upbound/provider-datadog/internal/controller/namespaced/metric/tagconfiguration"
	observabilitypipeline "github.com/upbound/provider-datadog/internal/controller/namespaced/observabilitypipeline/observabilitypipeline"
	escalationpolicy "github.com/upbound/provider-datadog/internal/controller/namespaced/oncall/escalationpolicy"
	schedule "github.com/upbound/provider-datadog/internal/controller/namespaced/oncall/schedule"
	teamroutingrules "github.com/upbound/provider-datadog/internal/controller/namespaced/oncall/teamroutingrules"
	usernotificationchannel "github.com/upbound/provider-datadog/internal/controller/namespaced/oncall/usernotificationchannel"
	usernotificationrule "github.com/upbound/provider-datadog/internal/controller/namespaced/oncall/usernotificationrule"
	providerconfig "github.com/upbound/provider-datadog/internal/controller/namespaced/providerconfig"
	table "github.com/upbound/provider-datadog/internal/controller/namespaced/reference/table"
	idpmetadata "github.com/upbound/provider-datadog/internal/controller/namespaced/saml/idpmetadata"
	duedaterule "github.com/upbound/provider-datadog/internal/controller/namespaced/securityfindings/duedaterule"
	muterule "github.com/upbound/provider-datadog/internal/controller/namespaced/securityfindings/muterule"
	ticketcreationrule "github.com/upbound/provider-datadog/internal/controller/namespaced/securityfindings/ticketcreationrule"
	criticalasset "github.com/upbound/provider-datadog/internal/controller/namespaced/securitymonitoring/criticalasset"
	defaultrule "github.com/upbound/provider-datadog/internal/controller/namespaced/securitymonitoring/defaultrule"
	filter "github.com/upbound/provider-datadog/internal/controller/namespaced/securitymonitoring/filter"
	notificationrulesecuritymonitoring "github.com/upbound/provider-datadog/internal/controller/namespaced/securitymonitoring/notificationrule"
	rule "github.com/upbound/provider-datadog/internal/controller/namespaced/securitymonitoring/rule"
	rulejson "github.com/upbound/provider-datadog/internal/controller/namespaced/securitymonitoring/rulejson"
	suppression "github.com/upbound/provider-datadog/internal/controller/namespaced/securitymonitoring/suppression"
	group "github.com/upbound/provider-datadog/internal/controller/namespaced/sensitivedatascanner/group"
	grouporder "github.com/upbound/provider-datadog/internal/controller/namespaced/sensitivedatascanner/grouporder"
	rulesensitivedatascanner "github.com/upbound/provider-datadog/internal/controller/namespaced/sensitivedatascanner/rule"
	page "github.com/upbound/provider-datadog/internal/controller/namespaced/statuspage/page"
	pagecomponent "github.com/upbound/provider-datadog/internal/controller/namespaced/statuspage/pagecomponent"
	pagedegradationtemplate "github.com/upbound/provider-datadog/internal/controller/namespaced/statuspage/pagedegradationtemplate"
	pagemaintenancetemplate "github.com/upbound/provider-datadog/internal/controller/namespaced/statuspage/pagemaintenancetemplate"
	concurrencycap "github.com/upbound/provider-datadog/internal/controller/namespaced/synthetics/concurrencycap"
	globalvariable "github.com/upbound/provider-datadog/internal/controller/namespaced/synthetics/globalvariable"
	privatelocation "github.com/upbound/provider-datadog/internal/controller/namespaced/synthetics/privatelocation"
	syntheticssuite "github.com/upbound/provider-datadog/internal/controller/namespaced/synthetics/syntheticssuite"
	test "github.com/upbound/provider-datadog/internal/controller/namespaced/synthetics/test"
	pipelinerulesets "github.com/upbound/provider-datadog/internal/controller/namespaced/tag/pipelinerulesets"
	ruletag "github.com/upbound/provider-datadog/internal/controller/namespaced/tag/rule"
	tagpipelineruleset "github.com/upbound/provider-datadog/internal/controller/namespaced/tag/tagpipelineruleset"
	workflowautomation "github.com/upbound/provider-datadog/internal/controller/namespaced/workflowautomation/workflowautomation"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connection.Setup,
		awsscanoptions.Setup,
		azurescanoptions.Setup,
		gcpscanoptions.Setup,
		softwarecatalog.Setup,
		retentionfilter.Setup,
		retentionfilterorder.Setup,
		app.Setup,
		wafcustomrule.Setup,
		wafexclusionfilter.Setup,
		cloudinventorysyncconfig.Setup,
		configurationrule.Setup,
		workloadsecurityagentrule.Setup,
		allocationrule.Setup,
		awscurconfig.Setup,
		azureucconfig.Setup,
		costbudget.Setup,
		costcustomforecast.Setup,
		customallocationrules.Setup,
		gcpucconfig.Setup,
		customframework.Setup,
		resourceevaluationfilter.Setup,
		agentrule.Setup,
		policy.Setup,
		dataset.Setup,
		datastore.Setup,
		datastoreitem.Setup,
		apikey.Setup,
		appkey.Setup,
		appkeyregistration.Setup,
		authnmapping.Setup,
		childorganization.Setup,
		dashboardjson.Setup,
		dashboardlist.Setup,
		domainallowlist.Setup,
		downtime.Setup,
		downtimeschedule.Setup,
		ipallowlist.Setup,
		monitor.Setup,
		monitorconfigpolicy.Setup,
		monitorjson.Setup,
		monitornotificationrule.Setup,
		organizationsettings.Setup,
		orgconnection.Setup,
		orggroup.Setup,
		orggroupmembership.Setup,
		orggrouppolicy.Setup,
		orggrouppolicyoverride.Setup,
		powerpack.Setup,
		powerpackv2.Setup,
		restrictionpolicy.Setup,
		role.Setup,
		rumapplication.Setup,
		rumexclusionfilter.Setup,
		rummetric.Setup,
		rumretentionfilter.Setup,
		rumretentionquota.Setup,
		secureembeddashboard.Setup,
		serviceaccesstoken.Setup,
		serviceaccount.Setup,
		serviceaccountapplicationkey.Setup,
		servicedefinitionyaml.Setup,
		servicelevelobjective.Setup,
		slocorrection.Setup,
		spansmetric.Setup,
		team.Setup,
		teamconnection.Setup,
		teamhierarchylinks.Setup,
		teamlink.Setup,
		teammembership.Setup,
		teamnotificationrule.Setup,
		teampermissionsetting.Setup,
		teamsync.Setup,
		user.Setup,
		userrole.Setup,
		webhook.Setup,
		webhookcustomvariable.Setup,
		webhookoauth2clientcredentials.Setup,
		control.Setup,
		incidenttype.Setup,
		notificationrule.Setup,
		notificationtemplate.Setup,
		userdefinedfield.Setup,
		userdefinedrole.Setup,
		awsaccount.Setup,
		awsaccountccmconfig.Setup,
		awseventbridge.Setup,
		awsexternalid.Setup,
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
		observabilitypipeline.Setup,
		escalationpolicy.Setup,
		schedule.Setup,
		teamroutingrules.Setup,
		usernotificationchannel.Setup,
		usernotificationrule.Setup,
		providerconfig.Setup,
		table.Setup,
		idpmetadata.Setup,
		duedaterule.Setup,
		muterule.Setup,
		ticketcreationrule.Setup,
		criticalasset.Setup,
		defaultrule.Setup,
		filter.Setup,
		notificationrulesecuritymonitoring.Setup,
		rule.Setup,
		rulejson.Setup,
		suppression.Setup,
		group.Setup,
		grouporder.Setup,
		rulesensitivedatascanner.Setup,
		page.Setup,
		pagecomponent.Setup,
		pagedegradationtemplate.Setup,
		pagemaintenancetemplate.Setup,
		concurrencycap.Setup,
		globalvariable.Setup,
		privatelocation.Setup,
		syntheticssuite.Setup,
		test.Setup,
		pipelinerulesets.Setup,
		ruletag.Setup,
		tagpipelineruleset.Setup,
		workflowautomation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connection.SetupGated,
		awsscanoptions.SetupGated,
		azurescanoptions.SetupGated,
		gcpscanoptions.SetupGated,
		softwarecatalog.SetupGated,
		retentionfilter.SetupGated,
		retentionfilterorder.SetupGated,
		app.SetupGated,
		wafcustomrule.SetupGated,
		wafexclusionfilter.SetupGated,
		cloudinventorysyncconfig.SetupGated,
		configurationrule.SetupGated,
		workloadsecurityagentrule.SetupGated,
		allocationrule.SetupGated,
		awscurconfig.SetupGated,
		azureucconfig.SetupGated,
		costbudget.SetupGated,
		costcustomforecast.SetupGated,
		customallocationrules.SetupGated,
		gcpucconfig.SetupGated,
		customframework.SetupGated,
		resourceevaluationfilter.SetupGated,
		agentrule.SetupGated,
		policy.SetupGated,
		dataset.SetupGated,
		datastore.SetupGated,
		datastoreitem.SetupGated,
		apikey.SetupGated,
		appkey.SetupGated,
		appkeyregistration.SetupGated,
		authnmapping.SetupGated,
		childorganization.SetupGated,
		dashboardjson.SetupGated,
		dashboardlist.SetupGated,
		domainallowlist.SetupGated,
		downtime.SetupGated,
		downtimeschedule.SetupGated,
		ipallowlist.SetupGated,
		monitor.SetupGated,
		monitorconfigpolicy.SetupGated,
		monitorjson.SetupGated,
		monitornotificationrule.SetupGated,
		organizationsettings.SetupGated,
		orgconnection.SetupGated,
		orggroup.SetupGated,
		orggroupmembership.SetupGated,
		orggrouppolicy.SetupGated,
		orggrouppolicyoverride.SetupGated,
		powerpack.SetupGated,
		powerpackv2.SetupGated,
		restrictionpolicy.SetupGated,
		role.SetupGated,
		rumapplication.SetupGated,
		rumexclusionfilter.SetupGated,
		rummetric.SetupGated,
		rumretentionfilter.SetupGated,
		rumretentionquota.SetupGated,
		secureembeddashboard.SetupGated,
		serviceaccesstoken.SetupGated,
		serviceaccount.SetupGated,
		serviceaccountapplicationkey.SetupGated,
		servicedefinitionyaml.SetupGated,
		servicelevelobjective.SetupGated,
		slocorrection.SetupGated,
		spansmetric.SetupGated,
		team.SetupGated,
		teamconnection.SetupGated,
		teamhierarchylinks.SetupGated,
		teamlink.SetupGated,
		teammembership.SetupGated,
		teamnotificationrule.SetupGated,
		teampermissionsetting.SetupGated,
		teamsync.SetupGated,
		user.SetupGated,
		userrole.SetupGated,
		webhook.SetupGated,
		webhookcustomvariable.SetupGated,
		webhookoauth2clientcredentials.SetupGated,
		control.SetupGated,
		incidenttype.SetupGated,
		notificationrule.SetupGated,
		notificationtemplate.SetupGated,
		userdefinedfield.SetupGated,
		userdefinedrole.SetupGated,
		awsaccount.SetupGated,
		awsaccountccmconfig.SetupGated,
		awseventbridge.SetupGated,
		awsexternalid.SetupGated,
		azure.SetupGated,
		cloudflareaccount.SetupGated,
		confluentaccount.SetupGated,
		confluentresource.SetupGated,
		fastlyaccount.SetupGated,
		fastlyservice.SetupGated,
		gcp.SetupGated,
		gcpsts.SetupGated,
		msteamstenantbasedhandle.SetupGated,
		msteamsworkflowswebhookhandle.SetupGated,
		opsgenieserviceobject.SetupGated,
		pagerduty.SetupGated,
		pagerdutyserviceobject.SetupGated,
		slackchannel.SetupGated,
		archive.SetupGated,
		archiveorder.SetupGated,
		customdestination.SetupGated,
		custompipeline.SetupGated,
		index.SetupGated,
		indexorder.SetupGated,
		integrationpipeline.SetupGated,
		metric.SetupGated,
		pipelineorder.SetupGated,
		restrictionquery.SetupGated,
		metadata.SetupGated,
		tagconfiguration.SetupGated,
		observabilitypipeline.SetupGated,
		escalationpolicy.SetupGated,
		schedule.SetupGated,
		teamroutingrules.SetupGated,
		usernotificationchannel.SetupGated,
		usernotificationrule.SetupGated,
		providerconfig.SetupGated,
		table.SetupGated,
		idpmetadata.SetupGated,
		duedaterule.SetupGated,
		muterule.SetupGated,
		ticketcreationrule.SetupGated,
		criticalasset.SetupGated,
		defaultrule.SetupGated,
		filter.SetupGated,
		notificationrulesecuritymonitoring.SetupGated,
		rule.SetupGated,
		rulejson.SetupGated,
		suppression.SetupGated,
		group.SetupGated,
		grouporder.SetupGated,
		rulesensitivedatascanner.SetupGated,
		page.SetupGated,
		pagecomponent.SetupGated,
		pagedegradationtemplate.SetupGated,
		pagemaintenancetemplate.SetupGated,
		concurrencycap.SetupGated,
		globalvariable.SetupGated,
		privatelocation.SetupGated,
		syntheticssuite.SetupGated,
		test.SetupGated,
		pipelinerulesets.SetupGated,
		ruletag.SetupGated,
		tagpipelineruleset.SetupGated,
		workflowautomation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		connection.SetupWebhookWithManager,
		awsscanoptions.SetupWebhookWithManager,
		azurescanoptions.SetupWebhookWithManager,
		gcpscanoptions.SetupWebhookWithManager,
		softwarecatalog.SetupWebhookWithManager,
		retentionfilter.SetupWebhookWithManager,
		retentionfilterorder.SetupWebhookWithManager,
		app.SetupWebhookWithManager,
		wafcustomrule.SetupWebhookWithManager,
		wafexclusionfilter.SetupWebhookWithManager,
		cloudinventorysyncconfig.SetupWebhookWithManager,
		configurationrule.SetupWebhookWithManager,
		workloadsecurityagentrule.SetupWebhookWithManager,
		allocationrule.SetupWebhookWithManager,
		awscurconfig.SetupWebhookWithManager,
		azureucconfig.SetupWebhookWithManager,
		costbudget.SetupWebhookWithManager,
		costcustomforecast.SetupWebhookWithManager,
		customallocationrules.SetupWebhookWithManager,
		gcpucconfig.SetupWebhookWithManager,
		customframework.SetupWebhookWithManager,
		resourceevaluationfilter.SetupWebhookWithManager,
		agentrule.SetupWebhookWithManager,
		policy.SetupWebhookWithManager,
		dataset.SetupWebhookWithManager,
		datastore.SetupWebhookWithManager,
		datastoreitem.SetupWebhookWithManager,
		apikey.SetupWebhookWithManager,
		appkey.SetupWebhookWithManager,
		appkeyregistration.SetupWebhookWithManager,
		authnmapping.SetupWebhookWithManager,
		childorganization.SetupWebhookWithManager,
		dashboardjson.SetupWebhookWithManager,
		dashboardlist.SetupWebhookWithManager,
		domainallowlist.SetupWebhookWithManager,
		downtime.SetupWebhookWithManager,
		downtimeschedule.SetupWebhookWithManager,
		ipallowlist.SetupWebhookWithManager,
		monitor.SetupWebhookWithManager,
		monitorconfigpolicy.SetupWebhookWithManager,
		monitorjson.SetupWebhookWithManager,
		monitornotificationrule.SetupWebhookWithManager,
		organizationsettings.SetupWebhookWithManager,
		orgconnection.SetupWebhookWithManager,
		orggroup.SetupWebhookWithManager,
		orggroupmembership.SetupWebhookWithManager,
		orggrouppolicy.SetupWebhookWithManager,
		orggrouppolicyoverride.SetupWebhookWithManager,
		powerpack.SetupWebhookWithManager,
		powerpackv2.SetupWebhookWithManager,
		restrictionpolicy.SetupWebhookWithManager,
		role.SetupWebhookWithManager,
		rumapplication.SetupWebhookWithManager,
		rumexclusionfilter.SetupWebhookWithManager,
		rummetric.SetupWebhookWithManager,
		rumretentionfilter.SetupWebhookWithManager,
		rumretentionquota.SetupWebhookWithManager,
		secureembeddashboard.SetupWebhookWithManager,
		serviceaccesstoken.SetupWebhookWithManager,
		serviceaccount.SetupWebhookWithManager,
		serviceaccountapplicationkey.SetupWebhookWithManager,
		servicedefinitionyaml.SetupWebhookWithManager,
		servicelevelobjective.SetupWebhookWithManager,
		slocorrection.SetupWebhookWithManager,
		spansmetric.SetupWebhookWithManager,
		team.SetupWebhookWithManager,
		teamconnection.SetupWebhookWithManager,
		teamhierarchylinks.SetupWebhookWithManager,
		teamlink.SetupWebhookWithManager,
		teammembership.SetupWebhookWithManager,
		teamnotificationrule.SetupWebhookWithManager,
		teampermissionsetting.SetupWebhookWithManager,
		teamsync.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		userrole.SetupWebhookWithManager,
		webhook.SetupWebhookWithManager,
		webhookcustomvariable.SetupWebhookWithManager,
		webhookoauth2clientcredentials.SetupWebhookWithManager,
		control.SetupWebhookWithManager,
		incidenttype.SetupWebhookWithManager,
		notificationrule.SetupWebhookWithManager,
		notificationtemplate.SetupWebhookWithManager,
		userdefinedfield.SetupWebhookWithManager,
		userdefinedrole.SetupWebhookWithManager,
		awsaccount.SetupWebhookWithManager,
		awsaccountccmconfig.SetupWebhookWithManager,
		awseventbridge.SetupWebhookWithManager,
		awsexternalid.SetupWebhookWithManager,
		azure.SetupWebhookWithManager,
		cloudflareaccount.SetupWebhookWithManager,
		confluentaccount.SetupWebhookWithManager,
		confluentresource.SetupWebhookWithManager,
		fastlyaccount.SetupWebhookWithManager,
		fastlyservice.SetupWebhookWithManager,
		gcp.SetupWebhookWithManager,
		gcpsts.SetupWebhookWithManager,
		msteamstenantbasedhandle.SetupWebhookWithManager,
		msteamsworkflowswebhookhandle.SetupWebhookWithManager,
		opsgenieserviceobject.SetupWebhookWithManager,
		pagerduty.SetupWebhookWithManager,
		pagerdutyserviceobject.SetupWebhookWithManager,
		slackchannel.SetupWebhookWithManager,
		archive.SetupWebhookWithManager,
		archiveorder.SetupWebhookWithManager,
		customdestination.SetupWebhookWithManager,
		custompipeline.SetupWebhookWithManager,
		index.SetupWebhookWithManager,
		indexorder.SetupWebhookWithManager,
		integrationpipeline.SetupWebhookWithManager,
		metric.SetupWebhookWithManager,
		pipelineorder.SetupWebhookWithManager,
		restrictionquery.SetupWebhookWithManager,
		metadata.SetupWebhookWithManager,
		tagconfiguration.SetupWebhookWithManager,
		observabilitypipeline.SetupWebhookWithManager,
		escalationpolicy.SetupWebhookWithManager,
		schedule.SetupWebhookWithManager,
		teamroutingrules.SetupWebhookWithManager,
		usernotificationchannel.SetupWebhookWithManager,
		usernotificationrule.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		table.SetupWebhookWithManager,
		idpmetadata.SetupWebhookWithManager,
		duedaterule.SetupWebhookWithManager,
		muterule.SetupWebhookWithManager,
		ticketcreationrule.SetupWebhookWithManager,
		criticalasset.SetupWebhookWithManager,
		defaultrule.SetupWebhookWithManager,
		filter.SetupWebhookWithManager,
		notificationrulesecuritymonitoring.SetupWebhookWithManager,
		rule.SetupWebhookWithManager,
		rulejson.SetupWebhookWithManager,
		suppression.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		grouporder.SetupWebhookWithManager,
		rulesensitivedatascanner.SetupWebhookWithManager,
		page.SetupWebhookWithManager,
		pagecomponent.SetupWebhookWithManager,
		pagedegradationtemplate.SetupWebhookWithManager,
		pagemaintenancetemplate.SetupWebhookWithManager,
		concurrencycap.SetupWebhookWithManager,
		globalvariable.SetupWebhookWithManager,
		privatelocation.SetupWebhookWithManager,
		syntheticssuite.SetupWebhookWithManager,
		test.SetupWebhookWithManager,
		pipelinerulesets.SetupWebhookWithManager,
		ruletag.SetupWebhookWithManager,
		tagpipelineruleset.SetupWebhookWithManager,
		workflowautomation.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
