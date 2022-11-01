package config

import (
	"github.com/FinanceFutureFactory/3fcoin/core/ethclient"
	"github.com/topfreegames/pitaya/v2/logger"
	"go.uber.org/zap"
)

type Chain struct {
	Id   string `json:"id" mapstructure:"id"`
	Type string `json:"type" mapstructure:"type"`
	Name string `json:"name" mapstructure:"name"`
	Node string `json:"node" mapstructure:"node"`
}

func (r *Chain) Connect() *ethclient.Client {
	if r.Type == "ETH" && r.Name == "FFF" {
		client, err := ethclient.Dial(r.Node)
		if err != nil {
			logger.Log.Error("eth client connect failed, err:", zap.Error(err))
			return nil
		}
		logger.Log.Info("eth client connect success")
		return client
	}
	logger.Log.Error("eth client error")
	return nil
}
