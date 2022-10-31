package configure

import (
	"context"
	"github.com/gotechbook/gotechbook-application/config"
	"github.com/gotechbook/gotechbook-application/db/mongo"
	"github.com/gotechbook/gotechbook-application/protos"
	"github.com/gotechbook/gotechbook-application/service"
	"github.com/gotechbook/gotechbook-application/utils"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
	"time"
)

type Topic struct {
	component.Base
}

func NewTopic() *Topic {
	return &Topic{}
}
func (t Topic) Create(ctx context.Context, msg *protos.TopicCreateRequest) (*protos.TopicCreateResponse, error) {
	logger.Log.Info("receive msg => ", msg)
	ret := &protos.TopicCreateResponse{}
	ret.Code = service.FAIL
	data := mongo.TbConfigureTopic{
		ID:        utils.NewSnowFlake().GenerateID(),
		Code:      msg.Code,
		Name:      msg.Name,
		Describe:  msg.Describe,
		IsRead:    false,
		IsDel:     false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := mongo.Insert(ctx, config.GOTECHBOOK_MONGO, mongo.DB_CONFIGURE, mongo.TbConfigureTopicName, &data)
	if err != nil {
		return ret, err
	}
	ret.Code = service.SUCCESS
	return ret, nil
}

func (t Topic) Delete(ctx context.Context, id string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (t Topic) DeleteByIds(ctx context.Context, ids []string) (int, error) {
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
