package domainService

import (
	"context"
	"github.com/faiz/go-mall/common/errcode"
	"github.com/faiz/go-mall/dal/dao"
	"github.com/faiz/go-mall/logic/domain"
)

type DemoDomainServiceV1 struct {
	Dao *dao.DemoDAO
}

func NewDemoDomainServiceV1(d *dao.DemoDAO) *DemoDomainServiceV1 {
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
