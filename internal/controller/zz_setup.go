// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	retentionfilter "github.com/upbound/provider-datadog/internal/controller/apm/retentionfilter"
	retentionfilterorder "github.com/upbound/provider-datadog/internal/controller/apm/retentionfilterorder"
	configurationrule "github.com/upbound/provider-datadog/internal/controller/cloud/configurationrule"
	workloadsecurityagentrule "github.com/upbound/provider-datadog/internal/controller/cloud/workloadsecurityagentrule"
	apikey "github.com/upbound/provider-datadog/internal/controller/datadog/apikey"
	appkey "github.com/upbound/provider-datadog/internal/controller/datadog/appkey"
	authnmapping "github.com/upbound/provider-datadog/internal/controller/datadog/authnmapping"
	childorganization "github.com/upbound/provider-datadog/internal/controller/datadog/childorganization"
	dashboardjson "github.com/upbound/provider-datadog/internal/controller/datadog/dashboardjson"
	dashboardlist "github.com/upbound/provider-datadog/internal/controller/datadog/dashboardlist"
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
	tagconfiguration "github.com/upbound/provider-datadog/internal/controller/metric/tagconfiguration"
	providerconfig "github.com/upbound/provider-datadog/internal/controller/providerconfig"
	defaultrule "github.com/upbound/provider-datadog/internal/controller/securitymonitoring/defaultrule"
	filter "github.com/upbound/provider-datadog/internal/controller/securitymonitoring/filter"
	rule "github.com/upbound/provider-datadog/internal/controller/securitymonitoring/rule"
	group "github.com/upbound/provider-datadog/internal/controller/sensitivedatascanner/group"
	grouporder "github.com/upbound/provider-datadog/internal/controller/sensitivedatascanner/grouporder"
	rulesensitivedatascanner "github.com/upbound/provider-datadog/internal/controller/sensitivedatascanner/rule"
	concurrencycap "github.com/upbound/provider-datadog/internal/controller/synthetics/concurrencycap"
	globalvariable "github.com/upbound/provider-datadog/internal/controller/synthetics/globalvariable"
	privatelocation "github.com/upbound/provider-datadog/internal/controller/synthetics/privatelocation"
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
