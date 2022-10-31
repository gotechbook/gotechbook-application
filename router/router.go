package router

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/cluster"
	"github.com/topfreegames/pitaya/v2/route"
)

func Register(app pitaya.Pitaya) {

}

func register(app pitaya.Pitaya, srvType string) {
	err := app.AddRoute(srvType, func(ctx context.Context, route *route.Route, payload []byte, servers map[string]*cluster.Server) (*cluster.Server, error) {
		for k := range servers {
			return servers[k], nil
		}
		return nil, nil
	})
	if err != nil {
		fmt.Printf("error adding router %s\n", err.Error())
	}
}
