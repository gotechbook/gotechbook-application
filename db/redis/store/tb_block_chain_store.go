package store

import (
	"context"
	"encoding/json"
	"github.com/gotechbook/gotechbook-application/config"
	"github.com/gotechbook/gotechbook-application/db/redis"
	"github.com/topfreegames/pitaya/v2/logger"
	"go.uber.org/zap"
	"time"
)

type BlockChainDataStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func NewBlockChainDataStore() *BlockChainDataStore {
	return &BlockChainDataStore{
		PreKey: redis.TbBlockChainKey,
	}
}

func (rs *BlockChainDataStore) UseWithCtx(ctx context.Context) *BlockChainDataStore {
	rs.Context = ctx
	return rs
}

func (rs *BlockChainDataStore) Set(value map[string]interface{}) error {
	err := config.GOTECHBOOK_REDIS.HSet(rs.Context, rs.PreKey, value).Err()
	if err != nil {
		logger.Log.Error("LogicTeamDetailStore Set Error!", zap.Error(err))
	}
	return err
}

func (rs *BlockChainDataStore) GetByKey(id string) (rst redis.BlockChainData, err error) {
	val, err := config.GOTECHBOOK_REDIS.HMGet(rs.Context, rs.PreKey, id).Result()
	if err != nil {
		logger.Log.Error("LogicTeamDetailStore GetByKey Error!", zap.Error(err))
		return rst, err
	}
	json.Unmarshal([]byte(val[0].(string)), &rst)
	logger.Log.Info("LogicTeamDetailStore GetKey val!", rst)
	return rst, err
}
