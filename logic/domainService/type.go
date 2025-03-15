package domainService

import (
	"context"
	"github.com/faiz/go-mall/logic/domain"
)

// DemoDomainService 保持依赖注入风格
type DemoDomainService interface {
	GetDemos(context.Context) ([]domain.DemoOrder, error)
}
