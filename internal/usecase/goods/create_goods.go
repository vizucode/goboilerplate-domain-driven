package goods

import (
	"context"
	"fmt"
	"goboilerplate-domain-driven/internal/domain/goods/entity.go"
	"goboilerplate-domain-driven/internal/domain/goods/repository"
)

type CreateGoods struct {
	repoOrder repository.GoodsRepository
}

func NewCreateGoods() (resp *CreateGoods) {
	return &CreateGoods{}
}

func (uc *CreateGoods) CreateGoods(ctx context.Context, req RequestGoods) (err error) {
	p := entity.Goods{}

	orders, err := uc.repoOrder.CreateGoods(ctx, &p)
	if err != nil {
		return err
	}

	fmt.Println(orders.ID)
	return nil
}
