package fff

import (
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
)

type GateBlocks struct {
	component.Base
	app pitaya.App
}

func NewGateBlocks(app pitaya.App) *GateBlocks {
	return &GateBlocks{
		app: app,
	}
}
