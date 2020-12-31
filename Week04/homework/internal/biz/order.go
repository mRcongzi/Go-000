// 业务逻辑处理层

package biz

// 定义Order的数据结构，即数据model
type Order struct {
	Item string
}

// 定义要一个Order的抽象类
// 在data层实现该接口
type OrderRepo interface {
	SaveOrder(*Order)
}

func NewOrderUsecase(repo OrderRepo) *OrderUsecase {
	return &OrderUsecase{repo: repo}
}

// 定义一个订单使用案例
type OrderUsecase struct {
	repo OrderRepo
}

// 订单案例拥有购买（Buy）的功能，在此实现业务逻辑
func (uc *OrderUsecase) Buy(o *Order) {
	uc.repo.SaveOrder(o)
}
