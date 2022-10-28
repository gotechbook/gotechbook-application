package service

import (
	"context"
)

const (
	CHAT   = "chat"
	CONFIG = "config"
)

type Service interface {
	Create(ctx context.Context, msg interface{}) (interface{}, error)
	Delete(ctx context.Context, id string) (int, error)
	DeleteByIds(ctx context.Context, ids []string) (int, error)
	Update(ctx context.Context, msg interface{}) (int, error)
	FindOne(ctx context.Context, msg interface{}) (interface{}, error)
	FindPage(ctx context.Context, msg interface{}) (interface{}, error)
	FindAll(ctx context.Context, msg interface{}) (interface{}, error)
}
