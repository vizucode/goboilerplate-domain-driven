package repository

import (
	"context"
	"goboilerplate-domain-driven/internal/domain/goods/entity.go"
)

type GoodsRepository interface {
	CreateGoods(ctx context.Context, p *entity.Goods) (resp *entity.Goods, err error)
	GetGoods(ctx context.Context, id int) (resp *entity.Goods, err error)
}
