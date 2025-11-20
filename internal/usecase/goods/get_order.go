package goods

import (
	"context"
	"fmt"
	"goboilerplate-domain-driven/internal/domain/goods/repository"
)

type GetAllGoods struct {
	repoOrder repository.GoodsRepository
}

func NewGetAllOrder() (resp *GetAllGoods) {
	return &GetAllGoods{}
}

func (uc *GetAllGoods) GetGoods(ctx context.Context, id int) (resp ResponseGoods, err error) {
	resultGoods, err := uc.repoOrder.GetGoods(ctx, id)
	if err != nil {
		return resp, err
	}

	resp = ResponseGoods{
		Id:   uint(resultGoods.ID),
		Name: resultGoods.Name,
	}

	fmt.Println(resp)
	return resp, err
}
