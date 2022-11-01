package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gotechbook/gotechbook-application/config"
	"github.com/gotechbook/gotechbook-application/router"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/cluster"
	c "github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/constants"
	"github.com/topfreegames/pitaya/v2/groups"
	"github.com/topfreegames/pitaya/v2/modules"
	"strconv"
)

func main() {
	path := flag.String("conf", "./fff.yaml", "config path")
	flag.Parse()
	config.Viper(*path, &config.GOTECHBOOK_FFF)

	config.LoadConfig(*path, &config.GOTECHBOOK_FFF)
	pitaya.SetLogger(config.SetLogger(fmt.Sprintf("./log/%s.log", config.GOTECHBOOK_FFF.App.Name), config.GOTECHBOOK_FFF.App.LogType, config.GOTECHBOOK_FFF.App.Name))
	config.GOTECHBOOK_REDIS = config.GOTECHBOOK_FFF.Redis.Connect()
	config.GOTECHBOOK_MONGO, _ = config.GOTECHBOOK_FFF.Mongo.MongoConfig(context.Background())
	config.GOTECHBOOK_FFF_CHAIN_CLIENT = config.GOTECHBOOK_FFF.Chain.Connect()

	app, bs := createApp()
	defer app.Shutdown()
	app.RegisterModule(bs, fmt.Sprintf("%s-storage", config.GOTECHBOOK_FFF.App.Name))
	router.FFF(app)
	app.Start()
}
func createApp() (pitaya.Pitaya, *modules.ETCDBindingStorage) {
	builderConfig := c.NewDefaultBuilderConfig()
	builderConfig.Pitaya = *config.GOTECHBOOK_FFF.Connection.ConnectionConfig()
	customMetrics := c.NewDefaultCustomMetricsSpec()
	prometheusConfig := c.NewDefaultPrometheusConfig()
	statsdConfig := c.NewDefaultStatsdConfig()
	etcdSDConfig := config.GOTECHBOOK_FFF.Discovery.EtcdDiscoveryConfig()
	natsRPCServerConfig := c.NewDefaultNatsRPCServerConfig()
	natsRPCClientConfig := c.NewDefaultNatsRPCClientConfig()
	workerConfig := config.GOTECHBOOK_FFF.Redis.WorkerConfig()
	enqueueOpts := c.NewDefaultEnqueueOpts()
	groupServiceConfig := c.NewDefaultMemoryGroupConfig()
	builder := pitaya.NewBuilder(false,
		config.GOTECHBOOK_FFF.App.Name,
		pitaya.Cluster,
		map[string]string{
			constants.GRPCHostKey: config.GOTECHBOOK_FFF.App.GrpcHost,
			constants.GRPCPortKey: strconv.Itoa(config.GOTECHBOOK_FFF.App.RpcPort),
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
	grpcServerConfig.Port = config.GOTECHBOOK_FFF.App.RpcPort
	gs, err := cluster.NewGRPCServer(*grpcServerConfig, builder.Server, builder.MetricsReporters)
	if err != nil {
		panic(err)
	}
	builder.RPCServer = gs
	builder.Groups = groups.NewMemoryGroupService(*c.NewDefaultMemoryGroupConfig())

	bs := modules.NewETCDBindingStorage(builder.Server, builder.SessionPool, *config.GOTECHBOOK_FFF.Modules.ETCDBindingConfig())
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
	return builder.Build(), bs
}
