package fff

import (
	"context"
	"github.com/gotechbook/gotechbook-application/config"
	"github.com/gotechbook/gotechbook-application/db/redis"
	"github.com/gotechbook/gotechbook-application/db/redis/store"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
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
	ctx := context.Background()
	blockChainStore := store.NewBlockChainDataStore()
	blockChainStore.UseWithCtx(ctx)

	b.timer = pitaya.NewTimer(time.Second*1, func() {

		//m := make(map[string]redis.BlockChainData)
		//blockChainStore.Set(m)

		chainId, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.ChainID(ctx)
		blockNumber, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.BlockNumber(ctx)
		data := redis.BlockChain{
			ChainId:           chainId.String(),
			Name:              "FFF",
			CurrentHeight:     int64(blockNumber),
			SynchronizeHeight: 0,
		}

		//// 同步高度
		////
		//blockNumber, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.BlockNumber(ctx)
		//block, err := config.GOTECHBOOK_FFF_CHAIN_CLIENT.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
		//logger.Log.Info(block.Header(), err)
	})
}
