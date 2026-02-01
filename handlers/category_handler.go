package handlers

import (
	"cashier-api/models"
	"cashier-api/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := h.service.GetAllCategories()
	if err != nil {
		log.Printf("Service Error: %s \n", err)

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Internal Server Error",
		})

		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"data": categories,
	})
}

func (h *CategoryHandler) StoreCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Bad Request",
		})

		return
	}

	h.service.CreateCategory(&category)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]any{
		"data": category,
	})
}

func (h *CategoryHandler) ShowCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid category id",
		})

		return
	}

	category, err := h.service.FindCategory(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Category not found",
		})

		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"data": category,
	})
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid category id",
		})

		return
	}

	var category models.Category

	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Bad request",
		})

		return
	}

	category.ID = id

	err = h.service.UpdateCategory(&category)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Category not found",
		})

		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"data": category,
	})
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid category id",
		})

		return
	}

	err = h.service.DeleteCategory(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Category not found",
		})

		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successful",
	})
}
