package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/DataDog/datadog-operator/pkg/apis"
	"github.com/DataDog/datadog-operator/pkg/controller"
	"github.com/DataDog/datadog-operator/version"

	"github.com/heptiolabs/healthcheck"
	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	"github.com/operator-framework/operator-sdk/pkg/leader"
	"github.com/operator-framework/operator-sdk/pkg/log/zap"
	"github.com/operator-framework/operator-sdk/pkg/restmapper"
	"github.com/spf13/pflag"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

// Change below variables to serve metrics on different host or port.
var (
	metricsHost              = "0.0.0.0"
	metricsPort        int32 = 8383
	printVersionArg    bool
	printVersionFormat string
	log                = logf.Log.WithName("cmd")
)

func main() {
	// Add the zap logger flag set to the CLI. The flag set must
	// be added before calling pflag.Parse().
	pflag.CommandLine.AddFlagSet(zap.FlagSet())

	// Add flags registered by imported packages (e.g. glog and
	// controller-runtime)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.BoolVarP(&printVersionArg, "version", "v", printVersionArg, "print version")
	pflag.StringVarP(&printVersionFormat, "version-format", "f", "text", "version output format ('text', 'json')")

	pflag.Parse()

	// Use a zap logr.Logger implementation. If none of the zap
	// flags are configured (or if the zap flag set is not being
	// used), this defaults to a production zap logger.
	//
	// The logger instantiated here can be changed to any logger
	// implementing the logr.Logger interface. This logger will
	// be propagated through the whole operator, generating
	// uniform and structured logs.
	logf.SetLogger(zap.Logger())

	if printVersionArg {
		version.PrintVersionWriter(os.Stdout, printVersionFormat)
		os.Exit(0)
	}

	version.PrintVersionLogs(log)

	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		log.Error(err, "Failed to get watch namespace")
		os.Exit(1)
	}

	// Get a config to talk to the apiserver
	cfg, err := config.GetConfig()
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	ctx := context.TODO()

	// Serve /live and /ready before leader election
	health := healthcheck.NewHandler()

	// Add a liveness check to detect go routine leaks. If this fails, we want to be restarted.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(200))

	go func() {
		log.Error(http.ListenAndServe(":8080", health),
			"Failed to listen on the probe endpoint. We’ll be killed by k8s.")
	}()

	// Become the leader before proceeding
	err = leader.Become(ctx, "datadog-operator-lock")
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Now that we’re a leader, the availability of /metrics becomes part of liveness check
	health.AddLivenessCheck("metrics", healthcheck.Async(healthcheck.HTTPGetCheck(fmt.Sprintf("http://localhost:%d/metrics", metricsPort), 5*time.Second), 10*time.Second))

	// Create a new Cmd to provide shared dependencies and start components
	mgr, err := manager.New(cfg, manager.Options{
		Namespace:          namespace,
		MapperProvider:     restmapper.NewDynamicRESTMapper,
		MetricsBindAddress: fmt.Sprintf("%s:%d", metricsHost, metricsPort),
	})
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	log.Info("Registering Components.")

	// Setup Scheme for all resources
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Setup all Controllers
	if err := controller.AddToManager(mgr); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	log.Info("Starting the Cmd.")

	// Start the Cmd
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		log.Error(err, "Manager exited non-zero")
		os.Exit(1)
	}
}
