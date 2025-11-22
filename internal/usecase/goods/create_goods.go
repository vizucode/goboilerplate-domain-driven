package goods

import (
	"context"
	"goboilerplate-domain-driven/internal/domain/goods/entity.go"
)

func (uc *ServiceGoods) CreateGoods(ctx context.Context, req RequestGoods) (err error) {
	p := entity.Goods{
		Name: req.Name,
	}

	_, err = uc.repo.CreateGoods(ctx, &p)
	if err != nil {
		return err
	}

	return nil
}
