package goods

import (
	"context"
	"goboilerplate-domain-driven/internal/domain/goods/entity.go"
)

func (uc *goodsRepository) CreateGoods(ctx context.Context, p *entity.Goods) (resp *entity.Goods, err error) {

	resp = &entity.Goods{}

	query := `INSERT INTO goods (name) VALUES ($1, $2)`

	err = uc.db.QueryRowContext(ctx, query, p.Name).Scan(&resp.ID, &resp.Name)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
