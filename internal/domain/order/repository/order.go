package repository

import "goboilerplate-domain-driven/internal/domain/order/entity.go"

type OrderRepository interface {
	CreateOrder(p *entity.Order) (resp *entity.Order, err error)
	GetOrder(id int) (resp *entity.Order, err error)
}
