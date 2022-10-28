package config

type App struct {
	Name     string `json:"name" mapstructure:"name"`
	Port     int    `json:"port" mapstructure:"port"`
	RpcPort  int    `json:"rpc-port" mapstructure:"rpc-port"`
	GrpcHost string `json:"grpc-host" mapstructure:"grpc-host"`
	LogType  string `json:"log-type" mapstructure:"log-type"`
}
