package goods

import (
	"goboilerplate-domain-driven/internal/adapter/external/jsonplaceholder"
	"goboilerplate-domain-driven/internal/domain/goods/repository"
)

type ServiceGoods struct {
	repo      repository.GoodsRepository
	jsonPlace jsonplaceholder.JsonPlaceHolder
}

func NewServiceGoods(
	repo repository.GoodsRepository,
	jsonPlace jsonplaceholder.JsonPlaceHolder,
) (resp *ServiceGoods) {
	return &ServiceGoods{
		repo:      repo,
		jsonPlace: jsonPlace,
	}
}
