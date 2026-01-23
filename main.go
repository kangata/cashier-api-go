package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var categories = []Category{
	{ID: 1, Name: "Makanan Utama", Description: "Hidangan porsi lengkap"},
	{ID: 2, Name: "Minuman Berkafein", Description: "Berbagai jenis olahan kopi dan teh"},
}

var products = []Product{
	{ID: 1, Name: "Indomie Godog", Price: 3500, Stock: 10},
	{ID: 2, Name: "Vit 100ml", Price: 3500, Stock: 40},
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"data": categories,
	})
}

func storeCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Bad Request",
		})

		return
	}

	category.ID = len(categories) + 1

	categories = append(categories, category)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]any{
		"data": category,
	})
}

func showCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Invalid Category ID",
		})

		return
	}

	for _, c := range categories {
		if c.ID == id {
			json.NewEncoder(w).Encode(map[string]any{
				"data": c,
			})

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Category Not Found",
	})
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Invalid Category ID",
		})

		return
	}

	var category Category

	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Bad Request",
		})

		return
	}

	for i := range categories {
		if categories[i].ID == id {
			category.ID = id
			categories[i] = category

			json.NewEncoder(w).Encode(map[string]any{
				"data": categories[i],
			})

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Category Not Found",
	})
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Invalid Category ID",
		})

		return
	}

	for i := range categories {
		if categories[i].ID == id {
			w.Header().Set("Content-Type", "applicaion/json")

			categories = append(categories[:i], categories[i+1:]...)

			json.NewEncoder(w).Encode(map[string]string{
				"message": "Successful",
			})

			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Category Not Found",
	})
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"data": categories,
	})
}

func storeProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Bad Request",
		})

		return
	}

	product.ID = len(products) + 1

	products = append(products, product)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]any{
		"data": product,
	})
}

func showProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Invalid Product ID",
		})

		return
	}

	for _, p := range products {
		if p.ID == id {
			json.NewEncoder(w).Encode(map[string]any{
				"data": p,
			})

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Product Not Found",
	})
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Invalid Product ID",
		})

		return
	}

	var product Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Bad Request",
		})

		return
	}

	for i := range products {
		if products[i].ID == id {
			product.ID = id
			products[i] = product

			json.NewEncoder(w).Encode(map[string]any{
				"data": products[i],
			})

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Product Not Found",
	})
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Invalid Product ID",
		})

		return
	}

	for i := range products {
		if products[i].ID == id {
			w.Header().Set("Content-Type", "applicaion/json")

			products = append(products[:i], products[i+1:]...)

			json.NewEncoder(w).Encode(map[string]string{
				"message": "Successful",
			})

			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Product Not Found",
	})
}

func main() {
	http.HandleFunc("GET /api/categories", getCategories)
	http.HandleFunc("POST /api/categories", storeCategory)
	http.HandleFunc("GET /api/categories/{id}", showCategory)
	http.HandleFunc("PUT /api/categories/{id}", updateCategory)
	http.HandleFunc("DELETE /api/categories/{id}", deleteCategory)

	http.HandleFunc("GET /api/products", getProducts)
	http.HandleFunc("POST /api/products", storeProduct)
	http.HandleFunc("GET /api/products/{id}", showProduct)
	http.HandleFunc("PUT /api/products/{id}", updateProduct)
	http.HandleFunc("DELETE /api/products/{id}", deleteProduct)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"message": "API Running",
		})
	})

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed running server")
	}
}
