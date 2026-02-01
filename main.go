package main

import (
	"cashier-api/database"
	"cashier-api/handlers"
	"cashier-api/models"
	"cashier-api/repositories"
	"cashier-api/services"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"APP_PORT"`
	DatabaseURL string `mapstructure:"DB_URL"`
}

var (
	config Config
	db     *sql.DB
)

var categories = []models.Category{
	{ID: 1, Name: "Makanan Utama", Description: "Hidangan porsi lengkap"},
	{ID: 2, Name: "Minuman Berkafein", Description: "Berbagai jenis olahan kopi dan teh"},
}

var products = []models.Product{
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

	var category models.Category

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

	var category models.Category

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

	var product models.Product

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

func initConfig() {
	viper.SetConfigFile(".env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed read config: %w", err))
	}

	viper.Unmarshal(&config)
}

func initDB() {
	var err error

	db, err = database.NewConnection(config.DatabaseURL)
	if err != nil {
		panic(fmt.Errorf("failed connect to database: %w", err))
	}
}

func main() {
	initConfig()
	initDB()

	defer db.Close()

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService, categoryService)

	http.HandleFunc("GET /api/categories", categoryHandler.GetAllCategories)
	http.HandleFunc("POST /api/categories", categoryHandler.StoreCategory)
	http.HandleFunc("GET /api/categories/{id}", categoryHandler.ShowCategory)
	http.HandleFunc("PUT /api/categories/{id}", categoryHandler.UpdateCategory)
	http.HandleFunc("DELETE /api/categories/{id}", categoryHandler.DeleteCategory)

	http.HandleFunc("GET /api/products", productHandler.GetAllProducts)
	http.HandleFunc("POST /api/products", productHandler.StoreProduct)
	http.HandleFunc("GET /api/products/{id}", productHandler.ShowProduct)
	http.HandleFunc("PUT /api/products/{id}", productHandler.UpdateProduct)
	http.HandleFunc("DELETE /api/products/{id}", productHandler.DeleteProduct)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"message": "API Running",
		})
	})

	addr := "0.0.0.0:" + config.Port

	fmt.Printf("Server running on %s \n", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Failed running server")
	}
}
