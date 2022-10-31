package router

import (
	"github.com/gotechbook/gotechbook-application/service/configure"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"strings"
)

func Configure(app pitaya.Pitaya) {
	topic := configure.NewTopic()
	app.RegisterRemote(topic, component.WithName(strings.ToLower("RemoveTopic")), component.WithNameFunc(strings.ToLower))
}

func GateConfigure(app pitaya.Pitaya) {
	topic := configure.NewGateTopic(app)
	app.Register(topic, component.WithName(strings.ToLower("topic")), component.WithNameFunc(strings.ToLower))
}
