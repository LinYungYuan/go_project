package model

import (
	"time"
)

type Order struct {
	ID         int64       `json:"id" db:"id"`
	UserID     int64       `json:"user_id" db:"user_id"`
	TotalPrice float64     `json:"total_price" db:"total_price"`
	Status     string      `json:"status" db:"status"`
	CreatedAt  time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at" db:"updated_at"`
	Items      []OrderItem `json:"items" db:"-"` // 這個字段不直接映射到數據庫
}

type OrderItem struct {
	ID        int64   `json:"id" db:"id"`
	OrderID   int64   `json:"order_id" db:"order_id"`
	ProductID int64   `json:"product_id" db:"product_id"`
	Quantity  int     `json:"quantity" db:"quantity"`
	Price     float64 `json:"price" db:"price"`
}
