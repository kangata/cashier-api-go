package main

import (
	"cashier-api/database"
	"cashier-api/handlers"
	"cashier-api/repositories"
	"cashier-api/services"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
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

func initConfig() {
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigFile(".env")

	_ = viper.ReadInConfig()

	viper.BindEnv("APP_PORT")
	viper.BindEnv("DB_URL")

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
