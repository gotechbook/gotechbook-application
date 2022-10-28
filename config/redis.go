package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/logger"
	"go.uber.org/zap"
)

type Redis struct {
	ServerURL   string `json:"server-url" yaml:"server-url"`
	Pool        string `json:"pool" yaml:"pool"`
	Password    string `json:"password" yaml:"password"`
	Concurrency int    `json:"concurrency" yaml:"concurrency"`
}

func (r *Redis) WorkerConfig() *config.WorkerConfig {
	return &config.WorkerConfig{
		Redis: struct {
			ServerURL string
			Pool      string
			Password  string
		}{
			ServerURL: r.ServerURL,
			Pool:      r.Pool,
			Password:  r.Password,
		},
		Concurrency: r.Concurrency,
	}
}

func (r *Redis) Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     r.ServerURL,
		Password: r.Password,    // no password set
		DB:       r.Concurrency, // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Log.Error("redis connect ping failed, err:", zap.Error(err))
		return nil
	} else {
		logger.Log.Info("redis connect ping response:", zap.String("pong", pong))
	}
	return client
}
