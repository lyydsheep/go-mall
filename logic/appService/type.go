package appService

import "context"

type DemoAppService interface {
	GetAllIdentities(c context.Context) ([]int64, error)
}
