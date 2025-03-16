package appService

import (
	"context"
	"github.com/faiz/go-mall/api/reply"
	"github.com/faiz/go-mall/api/request"
)

type DemoAppService interface {
	GetAllIdentities(c context.Context) ([]int64, error)
	CreateDemoOrder(c context.Context, order *request.DemoOrderReq) (*reply.DemoOrder, error)
}
