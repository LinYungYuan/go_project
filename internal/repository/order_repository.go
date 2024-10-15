// internal/repository/order_repository.go
package repository

import (
	"database/sql"
	"go_project/internal/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *model.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO orders (user_id, total_price, status, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = tx.QueryRow(query, order.UserID, order.TotalPrice, order.Status,
		order.CreatedAt, order.UpdatedAt).Scan(&order.ID)
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		query = `INSERT INTO order_items (order_id, product_id, quantity, price)
                 VALUES ($1, $2, $3, $4)`
		_, err = tx.Exec(query, order.ID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *OrderRepository) GetByID(id int64, userID int64) (*model.Order, error) {
	query := `SELECT id, user_id, total_price, status, created_at, updated_at 
              FROM orders WHERE id = $1 AND user_id = $2`

	var order model.Order
	err := r.db.QueryRow(query, id, userID).Scan(
		&order.ID, &order.UserID, &order.TotalPrice, &order.Status,
		&order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		return nil, err
	}

	query = `SELECT id, product_id, quantity, price 
             FROM order_items WHERE order_id = $1`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.OrderItem
		err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}
		order.Items = append(order.Items, item)
	}

	return &order, nil
}
