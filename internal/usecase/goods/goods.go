package goods

import (
	"goboilerplate-domain-driven/internal/domain/goods/repository"
)

type ServiceGoods struct {
	repo repository.GoodsRepository
}

func NewServiceGoods(
	repo repository.GoodsRepository,
) (resp *ServiceGoods) {
	return &ServiceGoods{
		repo: repo,
	}
}
