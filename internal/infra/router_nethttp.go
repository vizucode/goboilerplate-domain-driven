package infra

import httpgoods "goboilerplate-domain-driven/internal/adapter/http/goods"

func (s *server) RouteNetHttp(goodsAdapter httpgoods.GoodsHandler) {
	s.POST("/goods", goodsAdapter.Create)
	s.GET("/goods/:id", goodsAdapter.GetGood)
}
