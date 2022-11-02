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

type FFFBlockChainDataStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func NewBlockChainDataStore() *FFFBlockChainDataStore {
	return &FFFBlockChainDataStore{
		PreKey: redis.TbBlockChainKey,
	}
}

func (rs *FFFBlockChainDataStore) UseWithCtx(ctx context.Context) *FFFBlockChainDataStore {
	rs.Context = ctx
	return rs
}

func (rs *FFFBlockChainDataStore) Set(value redis.FFFBlockChain) error {
	err := config.GOTECHBOOK_REDIS.Set(rs.Context, rs.PreKey, value, 0).Err()
	if err != nil {
		logger.Log.Error("BlockChainDataStore Set Error!", zap.Error(err))
	}
	return err
}

func (rs *FFFBlockChainDataStore) Get() (rst redis.FFFBlockChain, err error) {
	val, err := config.GOTECHBOOK_REDIS.Get(rs.Context, rs.PreKey).Result()
	if err != nil {
		logger.Log.Error("BlockChainDataStore GetByKey Error!", zap.Error(err))
		return rst, err
	}
	json.Unmarshal([]byte(val), &rst)
	logger.Log.Info("BlockChainDataStore GetKey val!", rst)
	return rst, err
}
