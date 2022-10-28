package config

import (
	"github.com/topfreegames/pitaya/v2/config"
	"time"
)

type Connection struct {
	HandlerMessagesCompression   bool `json:"handler-messages-compression" mapstructure:"handler-messages-compression"`
	HeartbeatInterval            int  `json:"heartbeat-interval" mapstructure:"heartbeat-interval"`
	ConnRateLimitingInterval     int  `json:"conn-rate-limiting-interval" mapstructure:"conn-rate-limiting-interval"`
	ConnRateLimitingLimit        int  `json:"conn-rate-limiting-limit" mapstructure:"conn-rate-limiting-limit"`
	ConnRateLimitingForceDisable bool `json:"conn-rate-limiting-force-disable" mapstructure:"conn-rate-limiting-force-disable"`
}

func (c *Connection) RateLimitingConfig() *config.RateLimitingConfig {
	return &config.RateLimitingConfig{
		Limit:        c.ConnRateLimitingLimit,
		Interval:     time.Duration(c.ConnRateLimitingInterval),
		ForceDisable: c.ConnRateLimitingForceDisable,
	}
}

func (c *Connection) ConnectionConfig() *config.PitayaConfig {
	conf := config.NewDefaultPitayaConfig()
	conf.Heartbeat.Interval = time.Duration(c.HeartbeatInterval) * time.Second
	conf.Handler.Messages.Compression = c.HandlerMessagesCompression
	return conf
}
