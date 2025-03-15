package dao

import (
	"context"
	"github.com/faiz/go-mall/common/errcode"
	"github.com/faiz/go-mall/dal/model"
)

type DemoDAO struct {
}

func NewDemoDAO() *DemoDAO {
	return &DemoDAO{}
}

func (dao *DemoDAO) FindAllDemo(c context.Context) ([]model.DemoOrder, error) {
	var res []model.DemoOrder
	err := DB().WithContext(c).Model(&model.DemoOrder{}).Find(&res).Error
	if err != nil {
		err = errcode.Wrap("query entity error", err)
	}
	return res, err
}
