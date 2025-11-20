package order

import (
	"fmt"
	"goboilerplate-domain-driven/internal/domain/order/entity.go"
	"goboilerplate-domain-driven/internal/domain/order/repository"
)

type createOrder struct {
	repoOrder repository.OrderRepository
}

func NewCreateOrder() (resp *createOrder) {
	return &createOrder{}
}

func (uc *createOrder) CreateOrder() (err error) {
	p := entity.Order{}
	orders, err := uc.repoOrder.CreateOrder(&p)
	if err != nil {
		return err
	}

	fmt.Println(orders.ID)
	return nil
}
