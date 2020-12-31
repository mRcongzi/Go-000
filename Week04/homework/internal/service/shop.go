package service

import (
	"context"
	"github.com/mRcongzi/Go-000/Week04/homework/internal/biz"
)

type ShopService struct {
	ouc *biz.OrderUsecase
}

func NewShopService(ouc *biz.OrderUsecase) *ShopService {
	return &ShopService{ouc: ouc}
}

func (svr *ShopService) CreateOrder(ctx context.Context) error {
	o := new(biz.Order)
	o.Item = "item1"

	svr.ouc.Buy(o)
	return nil
}
