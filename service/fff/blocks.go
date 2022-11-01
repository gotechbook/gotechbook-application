package fff

import (
	"context"
	"github.com/gotechbook/gotechbook-application/config"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/timer"
	"math/big"
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
	ctx := context.Background()
	b.timer = pitaya.NewTimer(time.Second*1, func() {
		// 同步高度
		//chainId, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.ChainID(context.Background())
		blockNumber, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.BlockNumber(ctx)
		block, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
		logger.Log.Info(block.Header(), err)
	})
}
