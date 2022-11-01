package router

import (
	"github.com/gotechbook/gotechbook-application/service/fff"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"strings"
)

func FFF(app pitaya.Pitaya) {
	blocks := fff.NewBlocks()
	app.RegisterRemote(blocks, component.WithName(strings.ToLower("RemoveBlocks")), component.WithNameFunc(strings.ToLower))
}

func GateFFF(app pitaya.Pitaya) {
	//topic := configure.NewGateTopic(app)
	//app.Register(topic, component.WithName(strings.ToLower("topic")), component.WithNameFunc(strings.ToLower))
}
