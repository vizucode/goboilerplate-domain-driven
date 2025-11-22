package httpgoods

type RequestGoods struct {
	Name string `json:"name" validate:"required"`
}

type ResponseGoods struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
