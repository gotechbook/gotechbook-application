package config

type FFFConfig struct {
	App
	Connection  `mapstructure:"connection" json:"connection"`
	Concurrency `mapstructure:"concurrency" json:"concurrency"`
	Discovery   `mapstructure:"discovery" json:"discovery"`
	Redis       `mapstructure:"redis" json:"redis"`
	Modules     `mapstructure:"modules" json:"modules"`
	Mongo       `mapstructure:"mongo" json:"mongo"`
	Chain       `mapstructure:"chain" json:"chain"`
}
