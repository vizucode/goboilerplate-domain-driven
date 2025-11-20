package service

import "goboilerplate-domain-driven/internal/domain/order/entity.go"

type OrderCustomer struct{}

func (s *OrderCustomer) CustomerOrderMax(max int, lenOrder []*entity.Order) (err error) {
	if max >= len(lenOrder) {
		return err
	}
	return nil
}
