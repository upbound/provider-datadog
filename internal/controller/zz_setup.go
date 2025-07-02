/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	apikey "github.com/upbound/provider-datadog/internal/controller/access/apikey"
	applicationkey "github.com/upbound/provider-datadog/internal/controller/access/applicationkey"
	retentionfilter "github.com/upbound/provider-datadog/internal/controller/apm/retentionfilter"
	retentionfilterorder "github.com/upbound/provider-datadog/internal/controller/apm/retentionfilterorder"
	authnmapping "github.com/upbound/provider-datadog/internal/controller/authentication/authnmapping"
	configurationrule "github.com/upbound/provider-datadog/internal/controller/cloud/configurationrule"
	dashboardjson "github.com/upbound/provider-datadog/internal/controller/dashboard/dashboardjson"
	dashboardlist "github.com/upbound/provider-datadog/internal/controller/dashboard/dashboardlist"
	powerpack "github.com/upbound/provider-datadog/internal/controller/dashboard/powerpack"
	restrictionpolicy "github.com/upbound/provider-datadog/internal/controller/iam/restrictionpolicy"
	role "github.com/upbound/provider-datadog/internal/controller/iam/role"
	serviceaccount "github.com/upbound/provider-datadog/internal/controller/iam/serviceaccount"
	serviceaccountapplicationkey "github.com/upbound/provider-datadog/internal/controller/iam/serviceaccountapplicationkey"
	team "github.com/upbound/provider-datadog/internal/controller/iam/team"
	teamlink "github.com/upbound/provider-datadog/internal/controller/iam/teamlink"
	teammembership "github.com/upbound/provider-datadog/internal/controller/iam/teammembership"
	teampermissionsetting "github.com/upbound/provider-datadog/internal/controller/iam/teampermissionsetting"
	user "github.com/upbound/provider-datadog/internal/controller/iam/user"
	aws "github.com/upbound/provider-datadog/internal/controller/integration/aws"
	awseventbridge "github.com/upbound/provider-datadog/internal/controller/integration/awseventbridge"
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
	opsgenieserviceobject "github.com/upbound/provider-datadog/internal/controller/integration/opsgenieserviceobject"
	pagerduty "github.com/upbound/provider-datadog/internal/controller/integration/pagerduty"
	pagerdutyserviceobject "github.com/upbound/provider-datadog/internal/controller/integration/pagerdutyserviceobject"
	slackchannel "github.com/upbound/provider-datadog/internal/controller/integration/slackchannel"
	archive "github.com/upbound/provider-datadog/internal/controller/logs/archive"
	archiveorder "github.com/upbound/provider-datadog/internal/controller/logs/archiveorder"
	custompipeline "github.com/upbound/provider-datadog/internal/controller/logs/custompipeline"
	index "github.com/upbound/provider-datadog/internal/controller/logs/index"
	indexorder "github.com/upbound/provider-datadog/internal/controller/logs/indexorder"
	integrationpipeline "github.com/upbound/provider-datadog/internal/controller/logs/integrationpipeline"
	metric "github.com/upbound/provider-datadog/internal/controller/logs/metric"
	pipelineorder "github.com/upbound/provider-datadog/internal/controller/logs/pipelineorder"
	metadata "github.com/upbound/provider-datadog/internal/controller/metric/metadata"
	spansmetric "github.com/upbound/provider-datadog/internal/controller/metric/spansmetric"
	tagconfiguration "github.com/upbound/provider-datadog/internal/controller/metric/tagconfiguration"
	configpolicy "github.com/upbound/provider-datadog/internal/controller/monitor/configpolicy"
	downtime "github.com/upbound/provider-datadog/internal/controller/monitor/downtime"
	downtimeschedule "github.com/upbound/provider-datadog/internal/controller/monitor/downtimeschedule"
	monitor "github.com/upbound/provider-datadog/internal/controller/monitor/monitor"
	monitorjson "github.com/upbound/provider-datadog/internal/controller/monitor/monitorjson"
	webhook "github.com/upbound/provider-datadog/internal/controller/notification/webhook"
	webhookcustomvariable "github.com/upbound/provider-datadog/internal/controller/notification/webhookcustomvariable"
	childorganization "github.com/upbound/provider-datadog/internal/controller/organization/childorganization"
	settings "github.com/upbound/provider-datadog/internal/controller/organization/settings"
	providerconfig "github.com/upbound/provider-datadog/internal/controller/providerconfig"
	application "github.com/upbound/provider-datadog/internal/controller/rum/application"
	ipallowlist "github.com/upbound/provider-datadog/internal/controller/security/ipallowlist"
	defaultrule "github.com/upbound/provider-datadog/internal/controller/securitymonitoring/defaultrule"
	filter "github.com/upbound/provider-datadog/internal/controller/securitymonitoring/filter"
	rule "github.com/upbound/provider-datadog/internal/controller/securitymonitoring/rule"
	scannergroup "github.com/upbound/provider-datadog/internal/controller/sensitivedata/scannergroup"
	scannergrouporder "github.com/upbound/provider-datadog/internal/controller/sensitivedata/scannergrouporder"
	scannerrule "github.com/upbound/provider-datadog/internal/controller/sensitivedata/scannerrule"
	definitionyaml "github.com/upbound/provider-datadog/internal/controller/service/definitionyaml"
	correction "github.com/upbound/provider-datadog/internal/controller/slo/correction"
	servicelevelobjective "github.com/upbound/provider-datadog/internal/controller/slo/servicelevelobjective"
	concurrencycap "github.com/upbound/provider-datadog/internal/controller/synthetics/concurrencycap"
	globalvariable "github.com/upbound/provider-datadog/internal/controller/synthetics/globalvariable"
	privatelocation "github.com/upbound/provider-datadog/internal/controller/synthetics/privatelocation"
	test "github.com/upbound/provider-datadog/internal/controller/synthetics/test"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apikey.Setup,
		applicationkey.Setup,
		retentionfilter.Setup,
		retentionfilterorder.Setup,
		authnmapping.Setup,
		configurationrule.Setup,
		dashboardjson.Setup,
		dashboardlist.Setup,
		powerpack.Setup,
		restrictionpolicy.Setup,
		role.Setup,
		serviceaccount.Setup,
		serviceaccountapplicationkey.Setup,
		team.Setup,
		teamlink.Setup,
		teammembership.Setup,
		teampermissionsetting.Setup,
		user.Setup,
		aws.Setup,
		awseventbridge.Setup,
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
		opsgenieserviceobject.Setup,
		pagerduty.Setup,
		pagerdutyserviceobject.Setup,
		slackchannel.Setup,
		archive.Setup,
		archiveorder.Setup,
		custompipeline.Setup,
		index.Setup,
		indexorder.Setup,
		integrationpipeline.Setup,
		metric.Setup,
		pipelineorder.Setup,
		metadata.Setup,
		spansmetric.Setup,
		tagconfiguration.Setup,
		configpolicy.Setup,
		downtime.Setup,
		downtimeschedule.Setup,
		monitor.Setup,
		monitorjson.Setup,
		webhook.Setup,
		webhookcustomvariable.Setup,
		childorganization.Setup,
		settings.Setup,
		providerconfig.Setup,
		application.Setup,
		ipallowlist.Setup,
		defaultrule.Setup,
		filter.Setup,
		rule.Setup,
		scannergroup.Setup,
		scannergrouporder.Setup,
		scannerrule.Setup,
		definitionyaml.Setup,
		correction.Setup,
		servicelevelobjective.Setup,
		concurrencycap.Setup,
		globalvariable.Setup,
		privatelocation.Setup,
		test.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
