package fff

import (
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/timer"
	"time"
)

type Blocks struct {
	component.Base
	timer *timer.Timer
}

func NewBlocks() *Blocks {
	return &Blocks{}
}

func (b *Blocks) AfterInit() {
	b.timer = pitaya.NewTimer(time.Second, func() {
		logger.Log.Infof(time.Now().String())
	})
}
