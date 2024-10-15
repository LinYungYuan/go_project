// internal/repository/product_repository.go
package repository

import (
	"database/sql"
	"go_project/internal/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *model.Product) error {
	query := `INSERT INTO products (name, description, price, stock, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err := r.db.QueryRow(query, product.Name, product.Description, product.Price,
		product.Stock, product.CreatedAt, product.UpdatedAt).Scan(&product.ID)

	return err
}

func (r *ProductRepository) GetByID(id int64) (*model.Product, error) {
	query := `SELECT id, name, description, price, stock, created_at, updated_at 
              FROM products WHERE id = $1`

	var product model.Product
	err := r.db.QueryRow(query, id).Scan(
		&product.ID, &product.Name, &product.Description, &product.Price,
		&product.Stock, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) List() ([]*model.Product, error) {
	query := `SELECT id, name, description, price, stock, created_at, updated_at 
              FROM products`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price,
			&p.Stock, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}
