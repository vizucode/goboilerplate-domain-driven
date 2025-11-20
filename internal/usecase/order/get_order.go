package order

import (
	"fmt"
	"goboilerplate-domain-driven/internal/domain/order/repository"
)

type getAllOrder struct {
	repoOrder repository.OrderRepository
}

func NewGetAllOrder() (resp *createOrder) {
	return &createOrder{}
}

func (uc *createOrder) GetAllOrder() (err error) {
	resp, err := uc.repoOrder.GetOrder(1)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}
