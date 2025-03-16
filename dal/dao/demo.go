package dao

import (
	"context"
	"github.com/faiz/go-mall/common/errcode"
	"github.com/faiz/go-mall/common/util"
	"github.com/faiz/go-mall/dal/model"
	"github.com/faiz/go-mall/logic/domain"
)

type DemoDAOV1 struct {
}

func NewDemoDAO() *DemoDAOV1 {
	return &DemoDAOV1{}
}

func (dao *DemoDAOV1) FindAllDemo(c context.Context) ([]model.DemoOrder, error) {
	var res []model.DemoOrder
	err := DB().WithContext(c).Model(&model.DemoOrder{}).Find(&res).Error
	if err != nil {
		err = errcode.Wrap("query entity error", err)
	}
	return res, err
}

func (dao *DemoDAOV1) CreateDemoOrder(c context.Context, order *domain.DemoOrder) (*model.DemoOrder, error) {
	mod := new(model.DemoOrder)
	err := util.Convert(mod, order)
	if err != nil {
		return nil, err
	}
	return mod, DB().WithContext(c).Model(&model.DemoOrder{}).Create(mod).Error
}
