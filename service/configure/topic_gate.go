package configure

import (
	"context"
	"fmt"
	"github.com/gotechbook/gotechbook-application/protos"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"strings"
)

type GateTopic struct {
	component.Base
	app pitaya.Pitaya
}

type TopicCreateRequest struct {
	Name     string `json:"name,omitempty"`
	Describe string `json:"describe,omitempty"`
	Code     int64  `json:"code,omitempty"`
}

func NewGateTopic(app pitaya.Pitaya) *GateTopic {
	return &GateTopic{
		app: app,
	}
}

func (gate *GateTopic) Create(ctx context.Context, msg *TopicCreateRequest) (ret *protos.TopicCreateResponse, err error) {
	ret = &protos.TopicCreateResponse{}
	if err := gate.app.RPC(ctx, fmt.Sprintf("%s.%s.%s", "configure", strings.ToLower("RemoveTopic"), strings.ToLower("Create")), ret,
		&protos.TopicCreateRequest{
			Name:     msg.Name,
			Describe: msg.Describe,
			Code:     msg.Code,
		}); err != nil {
	}
	return ret, nil
}
