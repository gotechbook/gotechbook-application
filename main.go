package main

import (
	"flag"
	"fmt"
	"github.com/gotechbook/gotechbook-application-gate/config"
	"github.com/topfreegames/pitaya/v2"
)

func main() {
	path := flag.String("conf", "config.yaml", "config path")
	flag.Parse()
	config.Viper(*path, &config.GOTECHBOOK_GATE)

	config.LoadConfig(*path, &config.GOTECHBOOK_GATE)
	pitaya.SetLogger(config.SetLogger(fmt.Sprintf("./log/%s.log", config.GOTECHBOOK_GATE.App.Name), config.GOTECHBOOK_GATE.App.LogType, config.GOTECHBOOK_GATE.App.Name))
	config.GOTECHBOOK_REDIS = config.GOTECHBOOK_GATE.Redis.Connect()

	fmt.Println(config.GOTECHBOOK_GATE)
}
