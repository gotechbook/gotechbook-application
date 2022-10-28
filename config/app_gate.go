package config

import "github.com/go-redis/redis/v8"

var (
	GOTECHBOOK_GATE       GateConfig
	GOTECHBOOK_GATE_REDIS *redis.Client
)

type GateConfig struct {
	App         `json:"app" mapstructure:"app"`
	Connection  `json:"connection" mapstructure:"connection"`
	Concurrency `json:"concurrency" mapstructure:"concurrency"`
	Discovery   `json:"discovery" mapstructure:"discovery" `
	Redis       `json:"redis" mapstructure:"redis"`
	Modules     `json:"modules" mapstructure:"modules"`
}
