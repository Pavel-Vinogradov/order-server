package handler

import (
	"encoding/json"
	"net/http"
	"order-server/internal/product/dto"
	"order-server/pkg/request"
	"order-server/pkg/response"
	"strconv"
	"strings"

	"order-server/internal/product/entity"
	"order-server/internal/product/service"
)

type Handler struct {
	service *service.ProductService
}

func NewRouterProduct(service *service.ProductService) http.Handler {
	mux := http.NewServeMux()
	handler := &Handler{service: service}

	mux.HandleFunc("/products", handler.handleProducts)
	mux.HandleFunc("/products/", handler.handleProductByID)

	return mux
}

func (h *Handler) handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r, id)
	case http.MethodPut:
		h.Update(w, r, id)
	case http.MethodDelete:
		h.Delete(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	response.WriteJSON(w, http.StatusOK, products)
}

// Create godoc
// @Summary Создать продукт
// @Description Создает новый продукт
// @Tags products
// @Accept  json
// @Produce  json
// @Param   product body dto.CreateProductRequest true "Продукт"
// @Success 201 {object} entity.Product
// @Failure 400 {object} response.ErrorsResponse
// @Router /products [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := request.HandleBody[dto.CreateProductRequest](w, r)
	if err != nil {
		return
	}
	toEntity := body.ToEntity()

	created, err := h.service.CreateProduct(toEntity)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}
	response.WriteJSON(w, http.StatusCreated, created)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request, id int64) {
	product, err := h.service.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	response.WriteJSON(w, http.StatusOK, product)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request, id int64) {
	var p entity.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	p.ID = id
	updated, err := h.service.UpdateProduct(p)
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}
	response.WriteJSON(w, http.StatusOK, updated)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, id int64) {
	product, err := h.service.GetProductByID(id)
	_, err = h.service.DeleteProduct(product)

	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
