package config

type App struct {
	Name     string `json:"name" yaml:"name"`
	Port     int    `json:"port" yaml:"port"`
	RpcPort  int    `json:"rpc-port" yaml:"rpc-port"`
	GrpcHost string `json:"grpc-host" yaml:"grpc-host"`
	LogType  string `json:"log-type" yaml:"log-type"`
}
