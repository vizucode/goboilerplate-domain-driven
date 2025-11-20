package goods

import (
	"context"
	"goboilerplate-domain-driven/internal/domain/goods/entity.go"
)

func (uc *goodsRepository) GetGoods(ctx context.Context, id int) (resp *entity.Goods, err error) {
	resp = &entity.Goods{}

	row, err := uc.db.QueryContext(ctx, "SELECT * FROM goods WHERE id = $1", id)
	if err != nil {
		return resp, err
	}

	err = row.Scan(resp.ID, resp.Name)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
