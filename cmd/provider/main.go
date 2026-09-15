/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/alecthomas/kingpin/v2"
	changelogsv1alpha1 "github.com/crossplane/crossplane-runtime/v2/apis/changelogs/proto/v1alpha1"
	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/gate"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/customresourcesgate"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/statemetrics"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/crossplane/upjet/v2/pkg/controller/conversion"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	authv1 "k8s.io/api/authorization/v1"
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	apisCluster "github.com/upbound/provider-datadog/apis/cluster"
	apisNamespaced "github.com/upbound/provider-datadog/apis/namespaced"
	"github.com/upbound/provider-datadog/config"
	"github.com/upbound/provider-datadog/internal/clients"
	controllerCluster "github.com/upbound/provider-datadog/internal/controller/cluster"
	controllerNamespaced "github.com/upbound/provider-datadog/internal/controller/namespaced"
	"github.com/upbound/provider-datadog/internal/features"
	"github.com/upbound/provider-datadog/internal/version"
)

const (
	providerName = "provider-datadog"

	webhookTLSCertDirEnvVar = "WEBHOOK_TLS_CERT_DIR"
	tlsServerCertDirEnvVar  = "TLS_SERVER_CERTS_DIR"
	certsDirEnvVar          = "CERTS_DIR"
	tlsServerCertDir        = "/tls/server"
)

func main() {
	var (
		app                     = kingpin.New(filepath.Base(os.Args[0]), "Terraform based Crossplane provider for Datadog").DefaultEnvars()
		debug                   = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncPeriod              = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		pollInterval            = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		pollStateMetricInterval = app.Flag("poll-state-metric", "State metric recording interval").Default("5s").Duration()
		leaderElection          = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		maxReconcileRate        = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may be checked for drift from the desired state.").Default("10").Int()

		webhookPort            = app.Flag("webhook-port", "The port the webhook listens on").Default("9443").Envar("WEBHOOK_PORT").Int()
		metricsBindAddress     = app.Flag("metrics-bind-address", "The address the metrics server listens on").Default(":8080").Envar("METRICS_BIND_ADDRESS").String()
		healthProbeBindAddress = app.Flag("health-probe-bind-addr", "The address the health/readiness probe server listens on").Default(":8081").Envar("HEALTH_PROBE_BIND_ADDRESS").String()
		changelogsSocketPath   = app.Flag("changelogs-socket-path", "Path for changelogs socket (if enabled)").Default("/var/run/changelogs/changelogs.sock").Envar("CHANGELOGS_SOCKET_PATH").String()

		terraformVersion = app.Flag("terraform-version", "Terraform version.").Required().Envar("TERRAFORM_VERSION").String()
		providerSource   = app.Flag("terraform-provider-source", "Terraform provider source.").Required().Envar("TERRAFORM_PROVIDER_SOURCE").String()
		providerVersion  = app.Flag("terraform-provider-version", "Terraform provider version.").Required().Envar("TERRAFORM_PROVIDER_VERSION").String()

		enableManagementPolicies = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()
		enableChangeLogs         = app.Flag("enable-changelogs", "Enable support for capturing change logs during reconciliation.").Default("false").Envar("ENABLE_CHANGE_LOGS").Bool()
		enableSecretCache        = app.Flag("enable-secret-cache", "Enable caching of Secrets in the controller manager's informer cache. Reduces API server load at the cost of memory.").Default("true").Envar("ENABLE_SECRET_CACHE").Bool()

		certsDirSet = false
		// we record whether the command-line option "--certs-dir" was supplied
		// in the registered PreAction for the flag.
		certsDir = app.Flag("certs-dir", "The directory that contains the server key and certificate.").Default(tlsServerCertDir).Envar(certsDirEnvVar).PreAction(func(_ *kingpin.ParseContext) error {
			certsDirSet = true
			return nil
		}).String()
	)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName(providerName))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	log.Debug("Starting", "sync-period", syncPeriod.String(), "poll-interval", pollInterval.String(), "max-reconcile-rate", *maxReconcileRate)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	*certsDir = resolveCertsDir(certsDirSet, *certsDir)

	scheme := runtime.NewScheme()
	kingpin.FatalIfError(clientgoscheme.AddToScheme(scheme), "Cannot add client-go APIs to scheme")
	kingpin.FatalIfError(apisCluster.AddToScheme(scheme), "Cannot add cluster-scoped Datadog APIs to scheme")
	kingpin.FatalIfError(apisNamespaced.AddToScheme(scheme), "Cannot add namespaced Datadog APIs to scheme")
	kingpin.FatalIfError(apiextensionsv1.AddToScheme(scheme), "Cannot add api-extensions APIs to scheme")
	kingpin.FatalIfError(authv1.AddToScheme(scheme), "Cannot add k8s authorization APIs to scheme")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme:           scheme,
		Client:           clientOptions(*enableSecretCache),
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-provider-datadog",
		Cache: cache.Options{
			SyncPeriod: syncPeriod,
			ByObject: map[client.Object]cache.ByObject{
				&apiextensionsv1.CustomResourceDefinition{}: {
					Transform: customresourcesgate.TransformStripCRDSchema,
				},
			},
		},
		Metrics: metricsserver.Options{
			BindAddress: *metricsBindAddress,
		},
		HealthProbeBindAddress: *healthProbeBindAddress,
		WebhookServer: webhook.NewServer(webhook.Options{
			CertDir: *certsDir,
			Port:    *webhookPort,
		}),
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              new(60 * time.Second),
		RenewDeadline:              new(50 * time.Second),
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")
	if *certsDir != "" {
		kingpin.FatalIfError(mgr.AddReadyzCheck("webhook", mgr.GetWebhookServer().StartedChecker()), "Cannot add webhook server readyz checker to controller manager")
	}

	metricRecorder := managed.NewMRMetricRecorder()
	stateMetrics := statemetrics.NewMRStateMetrics()
	metrics.Registry.MustRegister(metricRecorder)
	metrics.Registry.MustRegister(stateMetrics)

	shared := sharedControllerOptions{
		log:                     log,
		pollInterval:            *pollInterval,
		pollStateMetricInterval: *pollStateMetricInterval,
		maxReconcileRate:        *maxReconcileRate,
		metricRecorder:          metricRecorder,
		stateMetrics:            stateMetrics,
		setupFn:                 clients.TerraformSetupBuilder(*terraformVersion, *providerSource, *providerVersion),
	}
	clusterOpts := shared.controllerOptions(config.GetProvider())
	namespacedOpts := shared.controllerOptions(config.GetProviderNamespaced())

	if *enableManagementPolicies {
		clusterOpts.Features.Enable(features.EnableBetaManagementPolicies)
		namespacedOpts.Features.Enable(features.EnableBetaManagementPolicies)
		log.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	if *enableChangeLogs {
		clo, err := changeLogOptions(*changelogsSocketPath)
		kingpin.FatalIfError(err, "Cannot enable change logs")
		clusterOpts.Features.Enable(feature.EnableAlphaChangeLogs)
		namespacedOpts.Features.Enable(feature.EnableAlphaChangeLogs)
		clusterOpts.ChangeLogOptions = clo
		namespacedOpts.ChangeLogOptions = clo
		log.Info("Alpha feature enabled", "flag", feature.EnableAlphaChangeLogs)
	}

	// Conversion webhooks are registered on every replica so that followers
	// can serve conversion requests too. Reconcilers only run on the leader.
	if *certsDir != "" {
		kingpin.FatalIfError(controllerCluster.SetupWebhookWithManager(mgr), "Cannot setup cluster-scoped Datadog webhooks")
		kingpin.FatalIfError(controllerNamespaced.SetupWebhookWithManager(mgr), "Cannot setup namespaced Datadog webhooks")
	}

	kingpin.FatalIfError(setupControllers(mgr, log, clusterOpts, namespacedOpts), "Cannot setup Datadog controllers")
	kingpin.FatalIfError(conversion.RegisterConversions(clusterOpts.Provider, namespacedOpts.Provider, mgr.GetScheme()), "Cannot initialize the webhook conversion registry")
	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}

// sharedControllerOptions holds the runtime settings that are identical for
// the cluster-scoped and the namespaced controller sets.
type sharedControllerOptions struct {
	log                     logging.Logger
	pollInterval            time.Duration
	pollStateMetricInterval time.Duration
	maxReconcileRate        int
	metricRecorder          *managed.MRMetricRecorder
	stateMetrics            *statemetrics.MRStateMetrics
	setupFn                 terraform.SetupFn
}

func (s sharedControllerOptions) controllerOptions(provider *ujconfig.Provider) tjcontroller.Options {
	return tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  s.log,
			GlobalRateLimiter:       ratelimiter.NewGlobal(s.maxReconcileRate),
			PollInterval:            s.pollInterval,
			MaxConcurrentReconciles: s.maxReconcileRate,
			Features:                &feature.Flags{},
			MetricOptions: &xpcontroller.MetricOptions{
				PollStateMetricInterval: s.pollStateMetricInterval,
				MRMetrics:               s.metricRecorder,
				MRStateMetrics:          s.stateMetrics,
			},
		},
		Provider: provider,
		// use the following WorkspaceStoreOption to enable the shared gRPC mode
		// terraform.WithProviderRunner(terraform.NewSharedProvider(log, os.Getenv("TERRAFORM_NATIVE_PROVIDER_PATH"), terraform.WithNativeProviderArgs("-debuggable")))
		WorkspaceStore: terraform.NewWorkspaceStore(s.log),
		SetupFn:        s.setupFn,
	}
}

// setupControllers starts both controller sets, gated on their CRDs becoming
// watchable when the provider is allowed to watch CRDs (SafeStart), and
// unconditionally otherwise.
func setupControllers(mgr manager.Manager, log logging.Logger, clusterOpts, namespacedOpts tjcontroller.Options) error {
	canSafeStart, err := canWatchCRD(context.TODO(), mgr)
	if err != nil {
		return errors.Wrap(err, "SafeStart precheck failed")
	}
	if !canSafeStart {
		log.Info("Provider has missing RBAC permissions for watching CRDs, controller SafeStart capability will be disabled")
		if err := controllerCluster.Setup(mgr, clusterOpts); err != nil {
			return errors.Wrap(err, "cannot setup cluster-scoped controllers")
		}
		return errors.Wrap(controllerNamespaced.Setup(mgr, namespacedOpts), "cannot setup namespaced controllers")
	}

	crdGate := new(gate.Gate[schema.GroupVersionKind])
	clusterOpts.Gate = crdGate
	namespacedOpts.Gate = crdGate
	if err := customresourcesgate.Setup(mgr, xpcontroller.Options{
		Logger:                  log,
		Gate:                    crdGate,
		MaxConcurrentReconciles: 1,
	}); err != nil {
		return errors.Wrap(err, "cannot setup CRD gate")
	}
	if err := controllerCluster.SetupGated(mgr, clusterOpts); err != nil {
		return errors.Wrap(err, "cannot setup gated cluster-scoped controllers")
	}
	return errors.Wrap(controllerNamespaced.SetupGated(mgr, namespacedOpts), "cannot setup gated namespaced controllers")
}

// clientOptions bypasses the informer cache for Secrets when secret caching is
// disabled, so every Secret read goes to the API server instead of memory.
func clientOptions(enableSecretCache bool) client.Options {
	if enableSecretCache {
		return client.Options{}
	}
	return client.Options{
		Cache: &client.CacheOptions{DisableFor: []client.Object{&corev1.Secret{}}},
	}
}

func changeLogOptions(socketPath string) (*xpcontroller.ChangeLogOptions, error) {
	conn, err := grpc.NewClient("unix://"+socketPath, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create change logs client connection at %s", socketPath)
	}
	return &xpcontroller.ChangeLogOptions{
		ChangeLogger: managed.NewGRPCChangeLogger(
			changelogsv1alpha1.NewChangeLogServiceClient(conn),
			managed.WithProviderVersion(fmt.Sprintf("%s:%s", providerName, version.Version))),
	}, nil
}

// resolveCertsDir returns the TLS certs directory to use. When no explicit
// certs dir was supplied via the command-line options, the environment
// variables set by Crossplane are used instead: WEBHOOK_TLS_CERT_DIR in older
// Crossplane versions, TLS_SERVER_CERTS_DIR in newer ones.
func resolveCertsDir(certsDirSet bool, certsDir string) string {
	if certsDirSet {
		return certsDir
	}
	for _, envVar := range []string{certsDirEnvVar, tlsServerCertDirEnvVar, webhookTLSCertDirEnvVar} {
		if dir := os.Getenv(envVar); dir != "" {
			return dir
		}
	}
	return certsDir
}

func canWatchCRD(ctx context.Context, mgr manager.Manager) (bool, error) {
	verbs := []string{"get", "list", "watch"}
	for _, verb := range verbs {
		sar := &authv1.SelfSubjectAccessReview{
			Spec: authv1.SelfSubjectAccessReviewSpec{
				ResourceAttributes: &authv1.ResourceAttributes{
					Group:    "apiextensions.k8s.io",
					Resource: "customresourcedefinitions",
					Verb:     verb,
				},
			},
		}
		if err := mgr.GetClient().Create(ctx, sar); err != nil {
			return false, errors.Wrapf(err, "unable to perform RBAC check for verb %s on CustomResourceDefinitions", verb)
		}
		if !sar.Status.Allowed {
			return false, nil
		}
	}
	return true, nil
}
