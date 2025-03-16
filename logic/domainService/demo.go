package domainService

import (
	"context"
	"github.com/faiz/go-mall/common/errcode"
	"github.com/faiz/go-mall/common/util"
	"github.com/faiz/go-mall/dal/dao"
	"github.com/faiz/go-mall/logic/domain"
)

type DemoDomainServiceV1 struct {
	Dao dao.DemoDAO
}

func NewDemoDomainServiceV1(d dao.DemoDAO) *DemoDomainServiceV1 {
	return &DemoDomainServiceV1{
		Dao: d,
	}
}

func (ds *DemoDomainServiceV1) GetDemos(c context.Context) ([]domain.DemoOrder, error) {
	demos, err := ds.Dao.FindAllDemo(c)
	if err != nil {
		err = errcode.Wrap("query entity error", err)
	}
	res := make([]domain.DemoOrder, 0, len(demos))
	for i := range demos {
		res = append(res, domain.DemoOrder{
			Id:           demos[i].Id,
			OrderId:      demos[i].OrderId,
			UserId:       demos[i].UserId,
			OrderGoodsId: demos[i].OrderGoodsId,
			BillMoney:    demos[i].BillMoney,
			State:        demos[i].State,
			PaidAt:       demos[i].PaidAt,
		})
	}
	return res, err
}

// CreateDemoOrder 创建订单
// 核心业务逻辑
func (ds *DemoDomainServiceV1) CreateDemoOrder(c context.Context, order *domain.DemoOrder) (*domain.DemoOrder, error) {
	order.OrderId = "this is random Id"
	demo, err := ds.Dao.CreateDemoOrder(c, order)
	if err != nil {
		return nil, errcode.Wrap("create entity error", err)
	}
	// do something...
	if err = util.Convert(order, demo); err != nil {
		return nil, errcode.Wrap("copy entity error", err)
	}
	return order, nil
}
