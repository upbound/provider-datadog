/*
Copyright 2026 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	connection "github.com/upbound/provider-datadog/internal/controller/cluster/action/connection"
	awsscanoptions "github.com/upbound/provider-datadog/internal/controller/cluster/agentlessscanning/awsscanoptions"
	azurescanoptions "github.com/upbound/provider-datadog/internal/controller/cluster/agentlessscanning/azurescanoptions"
	gcpscanoptions "github.com/upbound/provider-datadog/internal/controller/cluster/agentlessscanning/gcpscanoptions"
	softwarecatalog "github.com/upbound/provider-datadog/internal/controller/cluster/apicatalog/softwarecatalog"
	retentionfilter "github.com/upbound/provider-datadog/internal/controller/cluster/apm/retentionfilter"
	retentionfilterorder "github.com/upbound/provider-datadog/internal/controller/cluster/apm/retentionfilterorder"
	app "github.com/upbound/provider-datadog/internal/controller/cluster/appbuilder/app"
	wafcustomrule "github.com/upbound/provider-datadog/internal/controller/cluster/appsec/wafcustomrule"
	wafexclusionfilter "github.com/upbound/provider-datadog/internal/controller/cluster/appsec/wafexclusionfilter"
	cloudinventorysyncconfig "github.com/upbound/provider-datadog/internal/controller/cluster/cloud/cloudinventorysyncconfig"
	configurationrule "github.com/upbound/provider-datadog/internal/controller/cluster/cloud/configurationrule"
	workloadsecurityagentrule "github.com/upbound/provider-datadog/internal/controller/cluster/cloud/workloadsecurityagentrule"
	allocationrule "github.com/upbound/provider-datadog/internal/controller/cluster/cloudcost/allocationrule"
	awscurconfig "github.com/upbound/provider-datadog/internal/controller/cluster/cloudcost/awscurconfig"
	azureucconfig "github.com/upbound/provider-datadog/internal/controller/cluster/cloudcost/azureucconfig"
	costbudget "github.com/upbound/provider-datadog/internal/controller/cluster/cloudcost/costbudget"
	costcustomforecast "github.com/upbound/provider-datadog/internal/controller/cluster/cloudcost/costcustomforecast"
	customallocationrules "github.com/upbound/provider-datadog/internal/controller/cluster/cloudcost/customallocationrules"
	gcpucconfig "github.com/upbound/provider-datadog/internal/controller/cluster/cloudcost/gcpucconfig"
	customframework "github.com/upbound/provider-datadog/internal/controller/cluster/compliance/customframework"
	resourceevaluationfilter "github.com/upbound/provider-datadog/internal/controller/cluster/compliance/resourceevaluationfilter"
	agentrule "github.com/upbound/provider-datadog/internal/controller/cluster/csmthreats/agentrule"
	policy "github.com/upbound/provider-datadog/internal/controller/cluster/csmthreats/policy"
	dataset "github.com/upbound/provider-datadog/internal/controller/cluster/data/dataset"
	datastore "github.com/upbound/provider-datadog/internal/controller/cluster/data/datastore"
	datastoreitem "github.com/upbound/provider-datadog/internal/controller/cluster/data/datastoreitem"
	apikey "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/apikey"
	appkey "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/appkey"
	appkeyregistration "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/appkeyregistration"
	authnmapping "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/authnmapping"
	childorganization "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/childorganization"
	dashboardjson "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/dashboardjson"
	dashboardlist "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/dashboardlist"
	domainallowlist "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/domainallowlist"
	downtime "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/downtime"
	downtimeschedule "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/downtimeschedule"
	ipallowlist "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/ipallowlist"
	monitor "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/monitor"
	monitorconfigpolicy "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/monitorconfigpolicy"
	monitorjson "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/monitorjson"
	monitornotificationrule "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/monitornotificationrule"
	organizationsettings "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/organizationsettings"
	orgconnection "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/orgconnection"
	orggroup "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/orggroup"
	orggroupmembership "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/orggroupmembership"
	orggrouppolicy "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/orggrouppolicy"
	orggrouppolicyoverride "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/orggrouppolicyoverride"
	powerpack "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/powerpack"
	powerpackv2 "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/powerpackv2"
	restrictionpolicy "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/restrictionpolicy"
	role "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/role"
	rumapplication "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/rumapplication"
	rumexclusionfilter "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/rumexclusionfilter"
	rummetric "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/rummetric"
	rumretentionfilter "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/rumretentionfilter"
	rumretentionquota "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/rumretentionquota"
	secureembeddashboard "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/secureembeddashboard"
	serviceaccesstoken "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/serviceaccesstoken"
	serviceaccount "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/serviceaccount"
	serviceaccountapplicationkey "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/serviceaccountapplicationkey"
	servicedefinitionyaml "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/servicedefinitionyaml"
	servicelevelobjective "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/servicelevelobjective"
	slocorrection "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/slocorrection"
	spansmetric "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/spansmetric"
	team "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/team"
	teamconnection "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/teamconnection"
	teamhierarchylinks "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/teamhierarchylinks"
	teamlink "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/teamlink"
	teammembership "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/teammembership"
	teamnotificationrule "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/teamnotificationrule"
	teampermissionsetting "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/teampermissionsetting"
	teamsync "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/teamsync"
	user "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/user"
	userrole "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/userrole"
	webhook "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/webhook"
	webhookcustomvariable "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/webhookcustomvariable"
	webhookoauth2clientcredentials "github.com/upbound/provider-datadog/internal/controller/cluster/datadog/webhookoauth2clientcredentials"
	control "github.com/upbound/provider-datadog/internal/controller/cluster/governance/control"
	incidenttype "github.com/upbound/provider-datadog/internal/controller/cluster/incident/incidenttype"
	notificationrule "github.com/upbound/provider-datadog/internal/controller/cluster/incident/notificationrule"
	notificationtemplate "github.com/upbound/provider-datadog/internal/controller/cluster/incident/notificationtemplate"
	userdefinedfield "github.com/upbound/provider-datadog/internal/controller/cluster/incident/userdefinedfield"
	userdefinedrole "github.com/upbound/provider-datadog/internal/controller/cluster/incident/userdefinedrole"
	awsaccount "github.com/upbound/provider-datadog/internal/controller/cluster/integration/awsaccount"
	awsaccountccmconfig "github.com/upbound/provider-datadog/internal/controller/cluster/integration/awsaccountccmconfig"
	awseventbridge "github.com/upbound/provider-datadog/internal/controller/cluster/integration/awseventbridge"
	awsexternalid "github.com/upbound/provider-datadog/internal/controller/cluster/integration/awsexternalid"
	azure "github.com/upbound/provider-datadog/internal/controller/cluster/integration/azure"
	cloudflareaccount "github.com/upbound/provider-datadog/internal/controller/cluster/integration/cloudflareaccount"
	confluentaccount "github.com/upbound/provider-datadog/internal/controller/cluster/integration/confluentaccount"
	confluentresource "github.com/upbound/provider-datadog/internal/controller/cluster/integration/confluentresource"
	fastlyaccount "github.com/upbound/provider-datadog/internal/controller/cluster/integration/fastlyaccount"
	fastlyservice "github.com/upbound/provider-datadog/internal/controller/cluster/integration/fastlyservice"
	gcp "github.com/upbound/provider-datadog/internal/controller/cluster/integration/gcp"
	gcpsts "github.com/upbound/provider-datadog/internal/controller/cluster/integration/gcpsts"
	msteamstenantbasedhandle "github.com/upbound/provider-datadog/internal/controller/cluster/integration/msteamstenantbasedhandle"
	msteamsworkflowswebhookhandle "github.com/upbound/provider-datadog/internal/controller/cluster/integration/msteamsworkflowswebhookhandle"
	opsgenieserviceobject "github.com/upbound/provider-datadog/internal/controller/cluster/integration/opsgenieserviceobject"
	pagerduty "github.com/upbound/provider-datadog/internal/controller/cluster/integration/pagerduty"
	pagerdutyserviceobject "github.com/upbound/provider-datadog/internal/controller/cluster/integration/pagerdutyserviceobject"
	slackchannel "github.com/upbound/provider-datadog/internal/controller/cluster/integration/slackchannel"
	archive "github.com/upbound/provider-datadog/internal/controller/cluster/logs/archive"
	archiveorder "github.com/upbound/provider-datadog/internal/controller/cluster/logs/archiveorder"
	customdestination "github.com/upbound/provider-datadog/internal/controller/cluster/logs/customdestination"
	custompipeline "github.com/upbound/provider-datadog/internal/controller/cluster/logs/custompipeline"
	index "github.com/upbound/provider-datadog/internal/controller/cluster/logs/index"
	indexorder "github.com/upbound/provider-datadog/internal/controller/cluster/logs/indexorder"
	integrationpipeline "github.com/upbound/provider-datadog/internal/controller/cluster/logs/integrationpipeline"
	metric "github.com/upbound/provider-datadog/internal/controller/cluster/logs/metric"
	pipelineorder "github.com/upbound/provider-datadog/internal/controller/cluster/logs/pipelineorder"
	restrictionquery "github.com/upbound/provider-datadog/internal/controller/cluster/logs/restrictionquery"
	metadata "github.com/upbound/provider-datadog/internal/controller/cluster/metric/metadata"
	tagconfiguration "github.com/upbound/provider-datadog/internal/controller/cluster/metric/tagconfiguration"
	observabilitypipeline "github.com/upbound/provider-datadog/internal/controller/cluster/observabilitypipeline/observabilitypipeline"
	escalationpolicy "github.com/upbound/provider-datadog/internal/controller/cluster/oncall/escalationpolicy"
	schedule "github.com/upbound/provider-datadog/internal/controller/cluster/oncall/schedule"
	teamroutingrules "github.com/upbound/provider-datadog/internal/controller/cluster/oncall/teamroutingrules"
	usernotificationchannel "github.com/upbound/provider-datadog/internal/controller/cluster/oncall/usernotificationchannel"
	usernotificationrule "github.com/upbound/provider-datadog/internal/controller/cluster/oncall/usernotificationrule"
	providerconfig "github.com/upbound/provider-datadog/internal/controller/cluster/providerconfig"
	table "github.com/upbound/provider-datadog/internal/controller/cluster/reference/table"
	idpmetadata "github.com/upbound/provider-datadog/internal/controller/cluster/saml/idpmetadata"
	duedaterule "github.com/upbound/provider-datadog/internal/controller/cluster/securityfindings/duedaterule"
	muterule "github.com/upbound/provider-datadog/internal/controller/cluster/securityfindings/muterule"
	ticketcreationrule "github.com/upbound/provider-datadog/internal/controller/cluster/securityfindings/ticketcreationrule"
	criticalasset "github.com/upbound/provider-datadog/internal/controller/cluster/securitymonitoring/criticalasset"
	defaultrule "github.com/upbound/provider-datadog/internal/controller/cluster/securitymonitoring/defaultrule"
	filter "github.com/upbound/provider-datadog/internal/controller/cluster/securitymonitoring/filter"
	notificationrulesecuritymonitoring "github.com/upbound/provider-datadog/internal/controller/cluster/securitymonitoring/notificationrule"
	rule "github.com/upbound/provider-datadog/internal/controller/cluster/securitymonitoring/rule"
	rulejson "github.com/upbound/provider-datadog/internal/controller/cluster/securitymonitoring/rulejson"
	suppression "github.com/upbound/provider-datadog/internal/controller/cluster/securitymonitoring/suppression"
	group "github.com/upbound/provider-datadog/internal/controller/cluster/sensitivedatascanner/group"
	grouporder "github.com/upbound/provider-datadog/internal/controller/cluster/sensitivedatascanner/grouporder"
	rulesensitivedatascanner "github.com/upbound/provider-datadog/internal/controller/cluster/sensitivedatascanner/rule"
	page "github.com/upbound/provider-datadog/internal/controller/cluster/statuspage/page"
	pagecomponent "github.com/upbound/provider-datadog/internal/controller/cluster/statuspage/pagecomponent"
	pagedegradationtemplate "github.com/upbound/provider-datadog/internal/controller/cluster/statuspage/pagedegradationtemplate"
	pagemaintenancetemplate "github.com/upbound/provider-datadog/internal/controller/cluster/statuspage/pagemaintenancetemplate"
	concurrencycap "github.com/upbound/provider-datadog/internal/controller/cluster/synthetics/concurrencycap"
	globalvariable "github.com/upbound/provider-datadog/internal/controller/cluster/synthetics/globalvariable"
	privatelocation "github.com/upbound/provider-datadog/internal/controller/cluster/synthetics/privatelocation"
	syntheticssuite "github.com/upbound/provider-datadog/internal/controller/cluster/synthetics/syntheticssuite"
	test "github.com/upbound/provider-datadog/internal/controller/cluster/synthetics/test"
	pipelinerulesets "github.com/upbound/provider-datadog/internal/controller/cluster/tag/pipelinerulesets"
	ruletag "github.com/upbound/provider-datadog/internal/controller/cluster/tag/rule"
	tagpipelineruleset "github.com/upbound/provider-datadog/internal/controller/cluster/tag/tagpipelineruleset"
	workflowautomation "github.com/upbound/provider-datadog/internal/controller/cluster/workflowautomation/workflowautomation"
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
