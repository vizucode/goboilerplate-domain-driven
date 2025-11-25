package httpgoods

import (
	"encoding/json"
	"goboilerplate-domain-driven/internal/usecase/goods"
	"goboilerplate-domain-driven/pkg/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type GoodsHandler struct {
	goods   *goods.ServiceGoods
	formVal *validator.Validate
}

func NewGoodsHandler(
	serviceGoods *goods.ServiceGoods,
	formVal *validator.Validate,
) *GoodsHandler {
	return &GoodsHandler{
		goods:   serviceGoods,
		formVal: formVal,
	}
}

func (h *GoodsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req RequestGoods
		err error
	)
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, utils.HandleError(err.Error(), http.StatusBadRequest))
		return
	}

	err = h.formVal.StructCtx(ctx, &req)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	goodUsecase := goods.RequestGoods{
		Name: req.Name,
	}

	err = h.goods.CreateGoods(ctx, goodUsecase)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteOK(w, "Success create data goods", http.StatusCreated, nil)
}

func (h *GoodsHandler) GetGood(w http.ResponseWriter, r *http.Request) {

	var (
		responseData ResponseGoods
	)

	path := strings.TrimPrefix(r.URL.Path, "/goods/")
	if path == "" {
		utils.WriteError(w, utils.HandleError("ID not found", http.StatusBadRequest))
		return
	}

	GoodId, err := strconv.Atoi(path)
	if err != nil {
		utils.WriteError(w, utils.HandleError(err.Error(), http.StatusBadRequest))
		return
	}

	ctx := r.Context()

	resultGood, err := h.goods.GetGoods(ctx, GoodId)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	responseData = ResponseGoods{
		Id:   resultGood.Id,
		Name: resultGood.Name,
	}

	utils.WriteOK(w, "Successfully get data", http.StatusOK, responseData)
}
