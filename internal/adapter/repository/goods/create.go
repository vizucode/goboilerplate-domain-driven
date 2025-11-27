package goods

import (
	"context"
	"goboilerplate-domain-driven/internal/domain/goods/entity.go"

	"go.opentelemetry.io/otel"
)

func (uc *goodsRepository) CreateGoods(ctx context.Context, p *entity.Goods) (resp *entity.Goods, err error) {
	tracer := otel.Tracer("repository")
	ctx, span := tracer.Start(ctx, "Repository:CreateGoods")
	defer span.End()

	resp = &entity.Goods{}

	query := `INSERT INTO public.goods (name)
              VALUES ($1)
              RETURNING id, name`

	err = uc.db.QueryRowContext(ctx, query, p.Name).Scan(&resp.ID, &resp.Name)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
