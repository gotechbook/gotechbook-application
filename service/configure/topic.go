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

var (
	_ service.Service = (*Topic)(nil)
)

type Topic struct {
	component.Base
}

func (t Topic) Create(ctx context.Context, msg interface{}) (interface{}, error) {
	logger.Log.Info("receive msg => ", msg)
	ret := &protos.TopicCreateResponse{Code: service.SUCCESS}
	data := msg.(mongo.TbConfigureTopic)
	data.ID = utils.NewSnowFlake().GenerateID()
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	var _, err = mongo.Insert(ctx, config.GOTECHBOOK_MONGO, mongo.DB_CONFIGURE, mongo.TbConfigureTopicName, &data)
	if err != nil {
		ret.Code = service.FAIL
		return ret, err
	}
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

func NewTopic() *Topic {
	return &Topic{}
}
