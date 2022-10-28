package config

import (
	"github.com/topfreegames/pitaya/v2/config"
)

type Concurrency struct {
	BufferAgentMessages        int `json:"buffer-agent-messages" yaml:"buffer-agent-messages"`
	BufferHandlerLocalProcess  int `json:"buffer-handler-local-process" yaml:"buffer-handler-local-process"`
	BufferHandlerRemoteProcess int `json:"buffer-handler-remote-process" yaml:"buffer-handler-remote-process"`
	ConcurrencyHandlerDispatch int `json:"concurrency-handler-dispatch" yaml:"concurrency-handler-dispatch"`
}

func (c *Concurrency) ConcurrencyConfig(conf *config.PitayaConfig) *config.PitayaConfig {
	conf.Buffer.Agent.Messages = c.BufferAgentMessages
	conf.Buffer.Handler.LocalProcess = c.BufferHandlerLocalProcess
	conf.Buffer.Handler.RemoteProcess = c.BufferHandlerRemoteProcess
	conf.Concurrency.Handler.Dispatch = c.ConcurrencyHandlerDispatch
	return conf
}
