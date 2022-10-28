package config

type GateConfig struct {
	App         `json:"app" mapstructure:"app"`
	Connection  `json:"connection" mapstructure:"connection"`
	Concurrency `json:"concurrency" mapstructure:"concurrency"`
	Discovery   `json:"discovery" mapstructure:"discovery" `
	Redis       `json:"redis" mapstructure:"redis"`
	Modules     `json:"modules" mapstructure:"modules"`
}
