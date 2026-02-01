package repositories

import (
	"cashier-api/models"
	"database/sql"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAllProducts() ([]models.Product, error) {
	query := "SELECT id, name, price, stock FROM products"

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]models.Product, 0)

	for rows.Next() {
		var p models.Product

		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepository) CreateProduct(product *models.Product) error {
	query := "INSERT INTO products (category_id, name, price, stock) VALUES($1, $2, $3, $4) RETURNING id"

	err := repo.db.QueryRow(query, product.CategoryID, product.Name, product.Price, product.Stock).Scan(&product.ID)

	return err
}

func (repo *ProductRepository) FindProduct(id int) (*models.Product, error) {
	query := "SELECT id, category_id, name, price, stock FROM products WHERE id=$1"

	var p models.Product

	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.CategoryID, &p.Name, &p.Price, &p.Stock)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *ProductRepository) UpdateProduct(product *models.Product) error {
	query := "UPDATE products SET category_id=$1, name=$2, price=$3, stock=$4 WHERE id=$5"

	result, err := repo.db.Exec(query, product.CategoryID, product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (repo *ProductRepository) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id=$1"

	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}
