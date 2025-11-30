package goods

import (
	"context"
)

func (uc *ServiceGoods) GetGoods(ctx context.Context, id int) (resp ResponseGoods, err error) {
	resultGoods, err := uc.repo.GetGoods(ctx, id)
	if err != nil {
		return resp, err
	}

	uc.jsonPlace.FetchExternal(ctx)

	resp = ResponseGoods{
		Id:   uint(resultGoods.ID),
		Name: resultGoods.Name,
	}

	return resp, err
}
