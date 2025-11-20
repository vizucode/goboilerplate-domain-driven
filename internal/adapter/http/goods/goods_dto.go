package httpgoods

type RequestGoods struct {
	Id   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type ResponseGoods struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
