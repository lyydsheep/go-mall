package dao

import (
	"context"
	"github.com/faiz/go-mall/dal/model"
	"github.com/faiz/go-mall/logic/domain"
)

type DemoDAO interface {
	FindAllDemo(c context.Context) ([]model.DemoOrder, error)
	CreateDemoOrder(c context.Context, order *domain.DemoOrder) (*model.DemoOrder, error)
}
