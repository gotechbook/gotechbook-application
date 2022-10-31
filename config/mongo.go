package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Uri           string `mapstructure:"uri" json:"uri" yaml:"uri"`
	MaxPoolSize   int64  `mapstructure:"max-pool-size" json:"max-pool-size" yaml:"max-pool-size"`
	MinPoolSize   int64  `mapstructure:"min-pool-size" json:"min-pool-size" yaml:"min-pool-size"`
	MaxConnecting int64  `mapstructure:"max-connecting" json:"max-connecting" yaml:"max-connecting"`
}

func (m *Mongo) MongoConfig(ctx context.Context) (client *mongo.Client, err error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(m.Uri).
		SetMaxPoolSize(uint64(m.MaxPoolSize)).
		SetMinPoolSize(uint64(m.MinPoolSize)).
		SetMaxConnecting(uint64(m.MaxConnecting)))
}
