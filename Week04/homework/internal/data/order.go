package data

import (
	"fmt"
	"github.com/mRcongzi/Go-000/Week04/homework/internal/biz"
)

// 检验是否实现接口（语法糖）
var _ biz.OrderRepo = (*orderRepo)(nil)

func NewOrderRepo() biz.OrderRepo {
	return new(orderRepo)
}

type orderRepo struct{}

func (or *orderRepo) SaveOrder(o *biz.Order) {
	//
	fmt.Printf("%+v", o)
}
