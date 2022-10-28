package main

import (
	"flag"
	"fmt"
	"github.com/gotechbook/gotechbook-application-gate/config"
	"github.com/gotechbook/gotechbook-application-gate/route"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/cluster"
	c "github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/constants"
	"github.com/topfreegames/pitaya/v2/groups"
	"github.com/topfreegames/pitaya/v2/modules"
	"strconv"
)

func main() {
	path := flag.String("conf", "config.yaml", "config path")
	flag.Parse()
	config.Viper(*path, &config.GOTECHBOOK_GATE)

	config.LoadConfig(*path, &config.GOTECHBOOK_GATE)
	pitaya.SetLogger(config.SetLogger(fmt.Sprintf("./log/%s.log", config.GOTECHBOOK_GATE.App.Name), config.GOTECHBOOK_GATE.App.LogType, config.GOTECHBOOK_GATE.App.Name))
	config.GOTECHBOOK_REDIS = config.GOTECHBOOK_GATE.Redis.Connect()

	app, bs := createApp()
	defer app.Shutdown()
	app.RegisterModule(bs, fmt.Sprintf("%s-storage", config.GOTECHBOOK_GATE.App.Name))
	route.Register(app)
	app.Start()
}
func createApp() (pitaya.Pitaya, *modules.ETCDBindingStorage) {
	builderConfig := c.NewDefaultBuilderConfig()
	builderConfig.Pitaya = *config.GOTECHBOOK_GATE.Connection.ConnectionConfig()
	builderConfig.Metrics.Prometheus.Enabled = true

	customMetrics := c.NewDefaultCustomMetricsSpec()
	prometheusConfig := c.NewDefaultPrometheusConfig()
	statsdConfig := c.NewDefaultStatsdConfig()
	etcdSDConfig := config.GOTECHBOOK_GATE.Discovery.EtcdDiscoveryConfig()
	natsRPCServerConfig := c.NewDefaultNatsRPCServerConfig()
	natsRPCClientConfig := c.NewDefaultNatsRPCClientConfig()
	workerConfig := config.GOTECHBOOK_GATE.Redis.WorkerConfig()
	enqueueOpts := c.NewDefaultEnqueueOpts()
	groupServiceConfig := c.NewDefaultMemoryGroupConfig()
	builder := pitaya.NewBuilder(true,
		config.GOTECHBOOK_GATE.App.Name,
		pitaya.Cluster,
		map[string]string{
			constants.GRPCHostKey: config.GOTECHBOOK_GATE.App.GrpcHost,
			constants.GRPCPortKey: strconv.Itoa(config.GOTECHBOOK_GATE.App.RpcPort),
		},
		*builderConfig,
		*customMetrics,
		*prometheusConfig,
		*statsdConfig,
		*etcdSDConfig,
		*natsRPCServerConfig,
		*natsRPCClientConfig,
		*workerConfig,
		*enqueueOpts,
		*groupServiceConfig,
	)

	grpcServerConfig := c.NewDefaultGRPCServerConfig()
	grpcServerConfig.Port = config.GOTECHBOOK_GATE.App.RpcPort
	gs, err := cluster.NewGRPCServer(*grpcServerConfig, builder.Server, builder.MetricsReporters)
	if err != nil {
		panic(err)
	}
	builder.RPCServer = gs
	builder.Groups = groups.NewMemoryGroupService(*c.NewDefaultMemoryGroupConfig())

	bs := modules.NewETCDBindingStorage(builder.Server, builder.SessionPool, *config.GOTECHBOOK_GATE.Modules.ETCDBindingConfig())
	gc, err := cluster.NewGRPCClient(
		*c.NewDefaultGRPCClientConfig(),
		builder.Server,
		builder.MetricsReporters,
		bs,
		cluster.NewInfoRetriever(*c.NewDefaultInfoRetrieverConfig()),
	)
	if err != nil {
		panic(err)
	}
	builder.RPCClient = gc
	builder.AddAcceptor(acceptor.NewWSAcceptor(fmt.Sprintf(":%d", config.GOTECHBOOK_GATE.App.Port)))
	return builder.Build(), bs
}
