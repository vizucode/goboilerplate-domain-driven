package goods

import (
	"context"
	"goboilerplate-domain-driven/internal/domain/goods/entity.go"
	"goboilerplate-domain-driven/pkg/utils"

	"go.opentelemetry.io/otel"
)

func (uc *ServiceGoods) CreateGoods(ctx context.Context, req RequestGoods) (err error) {
	tracer := otel.Tracer("usecase")
	ctx, span := tracer.Start(ctx, "Usecase:CreateGoods")
	defer span.End()

	p := entity.Goods{
		Name: req.Name,
	}

	_, err = uc.repo.CreateGoods(ctx, &p)
	if err != nil {
		return err
	}

	utils.AddLogDebug(ctx, "Tataglia")

	return nil
}
