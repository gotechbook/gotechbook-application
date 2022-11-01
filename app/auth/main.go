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
	path := flag.String("conf", "./auth.yaml", "config path")
	flag.Parse()
	config.Viper(*path, &config.GOTECHBOOK_AUTH)

	config.LoadConfig(*path, &config.GOTECHBOOK_AUTH)
	pitaya.SetLogger(config.SetLogger(fmt.Sprintf("./log/%s.log", config.GOTECHBOOK_AUTH.App.Name), config.GOTECHBOOK_AUTH.App.LogType, config.GOTECHBOOK_AUTH.App.Name))
	config.GOTECHBOOK_REDIS = config.GOTECHBOOK_AUTH.Redis.Connect()
	config.GOTECHBOOK_MONGO, _ = config.GOTECHBOOK_AUTH.Mongo.MongoConfig(context.Background())
	app, bs := createApp()
	defer app.Shutdown()
	app.RegisterModule(bs, fmt.Sprintf("%s-storage", config.GOTECHBOOK_AUTH.App.Name))
	router.Configure(app)
	app.Start()
}
func createApp() (pitaya.Pitaya, *modules.ETCDBindingStorage) {
	builderConfig := c.NewDefaultBuilderConfig()
	builderConfig.Pitaya = *config.GOTECHBOOK_AUTH.Connection.ConnectionConfig()
	customMetrics := c.NewDefaultCustomMetricsSpec()
	prometheusConfig := c.NewDefaultPrometheusConfig()
	statsdConfig := c.NewDefaultStatsdConfig()
	etcdSDConfig := config.GOTECHBOOK_AUTH.Discovery.EtcdDiscoveryConfig()
	natsRPCServerConfig := c.NewDefaultNatsRPCServerConfig()
	natsRPCClientConfig := c.NewDefaultNatsRPCClientConfig()
	workerConfig := config.GOTECHBOOK_AUTH.Redis.WorkerConfig()
	enqueueOpts := c.NewDefaultEnqueueOpts()
	groupServiceConfig := c.NewDefaultMemoryGroupConfig()
	builder := pitaya.NewBuilder(false,
		config.GOTECHBOOK_AUTH.App.Name,
		pitaya.Cluster,
		map[string]string{
			constants.GRPCHostKey: config.GOTECHBOOK_AUTH.App.GrpcHost,
			constants.GRPCPortKey: strconv.Itoa(config.GOTECHBOOK_AUTH.App.RpcPort),
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
	grpcServerConfig.Port = config.GOTECHBOOK_AUTH.App.RpcPort
	gs, err := cluster.NewGRPCServer(*grpcServerConfig, builder.Server, builder.MetricsReporters)
	if err != nil {
		panic(err)
	}
	builder.RPCServer = gs
	builder.Groups = groups.NewMemoryGroupService(*c.NewDefaultMemoryGroupConfig())

	bs := modules.NewETCDBindingStorage(builder.Server, builder.SessionPool, *config.GOTECHBOOK_AUTH.Modules.ETCDBindingConfig())
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
