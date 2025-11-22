package httpgoods

import (
	"encoding/json"
	"goboilerplate-domain-driven/internal/usecase/goods"
	"net/http"
	"strconv"
	"strings"
)

type GoodsHandler struct {
	goods *goods.ServiceGoods
}

func NewGoodsHandler(
	serviceGoods *goods.ServiceGoods,
) *GoodsHandler {
	return &GoodsHandler{
		goods: serviceGoods,
	}
}

func (h *GoodsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req RequestGoods

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	goodUsecase := goods.RequestGoods{
		Name: req.Name,
	}

	ctx := r.Context()

	err := h.goods.CreateGoods(ctx, goodUsecase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]any{
		"status":  "OK",
		"message": "Created",
		"data":    nil,
	})
}

func (h *GoodsHandler) GetGood(w http.ResponseWriter, r *http.Request) {

	var (
		responseData ResponseGoods
	)

	path := strings.TrimPrefix(r.URL.Path, "/goods/")
	if path == "" {
		http.Error(w, "ID not found", http.StatusBadRequest)
		return
	}

	GoodId, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "invalid ID format", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	resultGood, err := h.goods.GetGoods(ctx, GoodId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData = ResponseGoods{
		Id:   resultGood.Id,
		Name: resultGood.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]any{
		"status":  "OK",
		"message": "Successfully get data",
		"data":    responseData,
	})
}
