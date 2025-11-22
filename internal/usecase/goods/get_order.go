package goods

import (
	"context"
	"goboilerplate-domain-driven/internal/domain/goods/repository"
)

type GetAllGoods struct {
	repo repository.GoodsRepository
}

func NewGetAllOrder(
	repo repository.GoodsRepository,
) (resp *GetAllGoods) {
	return &GetAllGoods{
		repo: repo,
	}
}

func (uc *GetAllGoods) GetGoods(ctx context.Context, id int) (resp ResponseGoods, err error) {
	resultGoods, err := uc.repo.GetGoods(ctx, id)
	if err != nil {
		return resp, err
	}

	resp = ResponseGoods{
		Id:   uint(resultGoods.ID),
		Name: resultGoods.Name,
	}

	return resp, err
}
