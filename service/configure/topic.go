package configure

import (
	"context"
	"github.com/gotechbook/gotechbook-application/service"
	"github.com/topfreegames/pitaya/v2/component"
)

var _ service.Service = (*Topic)(nil)

type Topic struct {
	component.Base
}

func (t Topic) Create(ctx context.Context, msg interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (t Topic) Delete(ctx context.Context, ids []string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t Topic) DeleteByIds(ctx context.Context, msg interface{}) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t Topic) Update(ctx context.Context, msg interface{}) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t Topic) FindOne(ctx context.Context, msg interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (t Topic) FindPage(ctx context.Context, msg interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (t Topic) FindAll(ctx context.Context, msg interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func NewTopic() *Topic {
	return &Topic{}
}
