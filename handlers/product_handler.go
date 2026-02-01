package handlers

import (
	"cashier-api/models"
	"cashier-api/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service         *services.ProductService
	categoryService *services.CategoryService
}

func NewProductHandler(service *services.ProductService, categoryService *services.CategoryService) *ProductHandler {
	return &ProductHandler{service: service, categoryService: categoryService}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.service.GetAllProducts()
	if err != nil {
		log.Printf("Service Error: %s \n", err)

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Internal Server Error",
		})

		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"data": products,
	})
}

func (h *ProductHandler) StoreProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Bad Request",
		})

		return
	}

	h.service.CreateProduct(&product)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]any{
		"data": product,
	})
}

func (h *ProductHandler) ShowProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid product id",
		})

		return
	}

	product, err := h.service.FindProduct(id)
	if err != nil {
		log.Printf("Failed to find product: %s \n", err)

		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Product not found",
		})

		return
	}

	if product.CategoryID != nil {
		category, err := h.categoryService.FindCategory(*product.CategoryID)
		if err == nil {
			product.Category = category
		}
	}

	json.NewEncoder(w).Encode(map[string]any{
		"data": product,
	})
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid product id",
		})

		return
	}

	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Bad request",
		})

		return
	}

	product.ID = id

	err = h.service.UpdateProduct(&product)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Product not found",
		})

		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"data": product,
	})
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid product id",
		})

		return
	}

	err = h.service.DeleteProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Product not found",
		})

		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successful",
	})
}
