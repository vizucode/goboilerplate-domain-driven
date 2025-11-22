package goods

import (
	"context"
	"goboilerplate-domain-driven/internal/domain/goods/entity.go"
	"goboilerplate-domain-driven/internal/domain/goods/repository"
)

type CreateGoods struct {
	repo repository.GoodsRepository
}

func NewCreateGoods(
	repo repository.GoodsRepository,
) (resp *CreateGoods) {
	return &CreateGoods{
		repo: repo,
	}
}

func (uc *CreateGoods) CreateGoods(ctx context.Context, req RequestGoods) (err error) {
	p := entity.Goods{}

	_, err = uc.repo.CreateGoods(ctx, &p)
	if err != nil {
		return err
	}

	return nil
}
