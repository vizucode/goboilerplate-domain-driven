package httpgoods

import (
	"encoding/json"
	"goboilerplate-domain-driven/internal/usecase/goods"
	"net/http"
	"strconv"
	"strings"
)

type goodsHandler struct {
	createGoods *goods.CreateGoods
	getAllGoods *goods.GetAllGoods
}

func NewGoodsHandler(
	createGoods *goods.CreateGoods,
	getAllGoods *goods.GetAllGoods,
) *goodsHandler {
	return &goodsHandler{
		createGoods: createGoods,
		getAllGoods: getAllGoods,
	}
}

func (h *goodsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req RequestGoods

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	goodUsecase := goods.RequestGoods{
		Id:   req.Id,
		Name: req.Name,
	}

	ctx := r.Context()

	err := h.createGoods.CreateGoods(ctx, goodUsecase)
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

func (h *goodsHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	var (
		responseData ResponseGoods
	)

	GoodIdStr := r.URL.Query().Get("id")
	if strings.EqualFold(GoodIdStr, "") {
		http.Error(w, "Query was not found", http.StatusBadRequest)
		return
	}

	GoodId, err := strconv.Atoi(GoodIdStr)
	if err != nil {
		http.Error(w, "invalid ID format", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	resultGood, err := h.getAllGoods.GetGoods(ctx, GoodId)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
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
