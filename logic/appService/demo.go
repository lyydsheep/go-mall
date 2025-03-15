package appService

import (
	"context"
	"github.com/faiz/go-mall/logic/domainService"
)

type DemoAppServiceV1 struct {
	ds domainService.DemoDomainService
}

func NewDemoAppServiceV1(domainSvc domainService.DemoDomainService) *DemoAppServiceV1 {
	return &DemoAppServiceV1{
		ds: domainSvc,
	}
}

func (as *DemoAppServiceV1) GetAllIdentities(c context.Context) ([]int64, error) {
	domains, err := as.ds.GetDemos(c)
	if err != nil {
		return nil, err
	}
	res := make([]int64, 0, len(domains))
	for i := range domains {
		res = append(res, domains[i].Id)
	}
	return res, nil
}
