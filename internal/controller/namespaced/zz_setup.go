/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	retentionfilter "github.com/upbound/provider-datadog/internal/controller/namespaced/apm/retentionfilter"
	retentionfilterorder "github.com/upbound/provider-datadog/internal/controller/namespaced/apm/retentionfilterorder"
	configurationrule "github.com/upbound/provider-datadog/internal/controller/namespaced/cloud/configurationrule"
	workloadsecurityagentrule "github.com/upbound/provider-datadog/internal/controller/namespaced/cloud/workloadsecurityagentrule"
	apikey "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/apikey"
	appkey "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/appkey"
	authnmapping "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/authnmapping"
	childorganization "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/childorganization"
	dashboardjson "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/dashboardjson"
	dashboardlist "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/dashboardlist"
	downtime "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/downtime"
	downtimeschedule "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/downtimeschedule"
	ipallowlist "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/ipallowlist"
	monitor "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/monitor"
	monitorconfigpolicy "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/monitorconfigpolicy"
	monitorjson "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/monitorjson"
	organizationsettings "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/organizationsettings"
	powerpack "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/powerpack"
	restrictionpolicy "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/restrictionpolicy"
	role "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/role"
	rumapplication "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/rumapplication"
	serviceaccount "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/serviceaccount"
	serviceaccountapplicationkey "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/serviceaccountapplicationkey"
	servicedefinitionyaml "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/servicedefinitionyaml"
	servicelevelobjective "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/servicelevelobjective"
	slocorrection "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/slocorrection"
	spansmetric "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/spansmetric"
	team "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/team"
	teamlink "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/teamlink"
	teammembership "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/teammembership"
	teampermissionsetting "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/teampermissionsetting"
	user "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/user"
	webhook "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/webhook"
	webhookcustomvariable "github.com/upbound/provider-datadog/internal/controller/namespaced/datadog/webhookcustomvariable"
	aws "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/aws"
	awseventbridge "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/awseventbridge"
	awslambdaarn "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/awslambdaarn"
	awslogcollection "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/awslogcollection"
	awstagfilter "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/awstagfilter"
	azure "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/azure"
	cloudflareaccount "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/cloudflareaccount"
	confluentaccount "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/confluentaccount"
	confluentresource "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/confluentresource"
	fastlyaccount "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/fastlyaccount"
	fastlyservice "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/fastlyservice"
	gcp "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/gcp"
	gcpsts "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/gcpsts"
	opsgenieserviceobject "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/opsgenieserviceobject"
	pagerduty "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/pagerduty"
	pagerdutyserviceobject "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/pagerdutyserviceobject"
	slackchannel "github.com/upbound/provider-datadog/internal/controller/namespaced/integration/slackchannel"
	archive "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/archive"
	archiveorder "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/archiveorder"
	custompipeline "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/custompipeline"
	index "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/index"
	indexorder "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/indexorder"
	integrationpipeline "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/integrationpipeline"
	metric "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/metric"
	pipelineorder "github.com/upbound/provider-datadog/internal/controller/namespaced/logs/pipelineorder"
	metadata "github.com/upbound/provider-datadog/internal/controller/namespaced/metric/metadata"
	tagconfiguration "github.com/upbound/provider-datadog/internal/controller/namespaced/metric/tagconfiguration"
	providerconfig "github.com/upbound/provider-datadog/internal/controller/namespaced/providerconfig"
	defaultrule "github.com/upbound/provider-datadog/internal/controller/namespaced/securitymonitoring/defaultrule"
	filter "github.com/upbound/provider-datadog/internal/controller/namespaced/securitymonitoring/filter"
	rule "github.com/upbound/provider-datadog/internal/controller/namespaced/securitymonitoring/rule"
	group "github.com/upbound/provider-datadog/internal/controller/namespaced/sensitivedatascanner/group"
	grouporder "github.com/upbound/provider-datadog/internal/controller/namespaced/sensitivedatascanner/grouporder"
	rulesensitivedatascanner "github.com/upbound/provider-datadog/internal/controller/namespaced/sensitivedatascanner/rule"
	concurrencycap "github.com/upbound/provider-datadog/internal/controller/namespaced/synthetics/concurrencycap"
	globalvariable "github.com/upbound/provider-datadog/internal/controller/namespaced/synthetics/globalvariable"
	privatelocation "github.com/upbound/provider-datadog/internal/controller/namespaced/synthetics/privatelocation"
	test "github.com/upbound/provider-datadog/internal/controller/namespaced/synthetics/test"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		retentionfilter.Setup,
		retentionfilterorder.Setup,
		configurationrule.Setup,
		workloadsecurityagentrule.Setup,
		apikey.Setup,
		appkey.Setup,
		authnmapping.Setup,
		childorganization.Setup,
		dashboardjson.Setup,
		dashboardlist.Setup,
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
		tagconfiguration.Setup,
		providerconfig.Setup,
		defaultrule.Setup,
		filter.Setup,
		rule.Setup,
		group.Setup,
		grouporder.Setup,
		rulesensitivedatascanner.Setup,
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

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		retentionfilter.SetupGated,
		retentionfilterorder.SetupGated,
		configurationrule.SetupGated,
		workloadsecurityagentrule.SetupGated,
		apikey.SetupGated,
		appkey.SetupGated,
		authnmapping.SetupGated,
		childorganization.SetupGated,
		dashboardjson.SetupGated,
		dashboardlist.SetupGated,
		downtime.SetupGated,
		downtimeschedule.SetupGated,
		ipallowlist.SetupGated,
		monitor.SetupGated,
		monitorconfigpolicy.SetupGated,
		monitorjson.SetupGated,
		organizationsettings.SetupGated,
		powerpack.SetupGated,
		restrictionpolicy.SetupGated,
		role.SetupGated,
		rumapplication.SetupGated,
		serviceaccount.SetupGated,
		serviceaccountapplicationkey.SetupGated,
		servicedefinitionyaml.SetupGated,
		servicelevelobjective.SetupGated,
		slocorrection.SetupGated,
		spansmetric.SetupGated,
		team.SetupGated,
		teamlink.SetupGated,
		teammembership.SetupGated,
		teampermissionsetting.SetupGated,
		user.SetupGated,
		webhook.SetupGated,
		webhookcustomvariable.SetupGated,
		aws.SetupGated,
		awseventbridge.SetupGated,
		awslambdaarn.SetupGated,
		awslogcollection.SetupGated,
		awstagfilter.SetupGated,
		azure.SetupGated,
		cloudflareaccount.SetupGated,
		confluentaccount.SetupGated,
		confluentresource.SetupGated,
		fastlyaccount.SetupGated,
		fastlyservice.SetupGated,
		gcp.SetupGated,
		gcpsts.SetupGated,
		opsgenieserviceobject.SetupGated,
		pagerduty.SetupGated,
		pagerdutyserviceobject.SetupGated,
		slackchannel.SetupGated,
		archive.SetupGated,
		archiveorder.SetupGated,
		custompipeline.SetupGated,
		index.SetupGated,
		indexorder.SetupGated,
		integrationpipeline.SetupGated,
		metric.SetupGated,
		pipelineorder.SetupGated,
		metadata.SetupGated,
		tagconfiguration.SetupGated,
		providerconfig.SetupGated,
		defaultrule.SetupGated,
		filter.SetupGated,
		rule.SetupGated,
		group.SetupGated,
		grouporder.SetupGated,
		rulesensitivedatascanner.SetupGated,
		concurrencycap.SetupGated,
		globalvariable.SetupGated,
		privatelocation.SetupGated,
		test.SetupGated,
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
		retentionfilter.SetupWebhookWithManager,
		retentionfilterorder.SetupWebhookWithManager,
		configurationrule.SetupWebhookWithManager,
		workloadsecurityagentrule.SetupWebhookWithManager,
		apikey.SetupWebhookWithManager,
		appkey.SetupWebhookWithManager,
		authnmapping.SetupWebhookWithManager,
		childorganization.SetupWebhookWithManager,
		dashboardjson.SetupWebhookWithManager,
		dashboardlist.SetupWebhookWithManager,
		downtime.SetupWebhookWithManager,
		downtimeschedule.SetupWebhookWithManager,
		ipallowlist.SetupWebhookWithManager,
		monitor.SetupWebhookWithManager,
		monitorconfigpolicy.SetupWebhookWithManager,
		monitorjson.SetupWebhookWithManager,
		organizationsettings.SetupWebhookWithManager,
		powerpack.SetupWebhookWithManager,
		restrictionpolicy.SetupWebhookWithManager,
		role.SetupWebhookWithManager,
		rumapplication.SetupWebhookWithManager,
		serviceaccount.SetupWebhookWithManager,
		serviceaccountapplicationkey.SetupWebhookWithManager,
		servicedefinitionyaml.SetupWebhookWithManager,
		servicelevelobjective.SetupWebhookWithManager,
		slocorrection.SetupWebhookWithManager,
		spansmetric.SetupWebhookWithManager,
		team.SetupWebhookWithManager,
		teamlink.SetupWebhookWithManager,
		teammembership.SetupWebhookWithManager,
		teampermissionsetting.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		webhook.SetupWebhookWithManager,
		webhookcustomvariable.SetupWebhookWithManager,
		aws.SetupWebhookWithManager,
		awseventbridge.SetupWebhookWithManager,
		awslambdaarn.SetupWebhookWithManager,
		awslogcollection.SetupWebhookWithManager,
		awstagfilter.SetupWebhookWithManager,
		azure.SetupWebhookWithManager,
		cloudflareaccount.SetupWebhookWithManager,
		confluentaccount.SetupWebhookWithManager,
		confluentresource.SetupWebhookWithManager,
		fastlyaccount.SetupWebhookWithManager,
		fastlyservice.SetupWebhookWithManager,
		gcp.SetupWebhookWithManager,
		gcpsts.SetupWebhookWithManager,
		opsgenieserviceobject.SetupWebhookWithManager,
		pagerduty.SetupWebhookWithManager,
		pagerdutyserviceobject.SetupWebhookWithManager,
		slackchannel.SetupWebhookWithManager,
		archive.SetupWebhookWithManager,
		archiveorder.SetupWebhookWithManager,
		custompipeline.SetupWebhookWithManager,
		index.SetupWebhookWithManager,
		indexorder.SetupWebhookWithManager,
		integrationpipeline.SetupWebhookWithManager,
		metric.SetupWebhookWithManager,
		pipelineorder.SetupWebhookWithManager,
		metadata.SetupWebhookWithManager,
		tagconfiguration.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		defaultrule.SetupWebhookWithManager,
		filter.SetupWebhookWithManager,
		rule.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		grouporder.SetupWebhookWithManager,
		rulesensitivedatascanner.SetupWebhookWithManager,
		concurrencycap.SetupWebhookWithManager,
		globalvariable.SetupWebhookWithManager,
		privatelocation.SetupWebhookWithManager,
		test.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
