package configure

import (
	"context"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
)

type GateTopic struct {
	component.Base
	app pitaya.Pitaya
}

func NewGateTopic(app pitaya.Pitaya) *GateTopic {
	return &GateTopic{
		app: app,
	}
}

func (gate *GateTopic) Create(ctx context.Context) {
	logger.Log.Info("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}
