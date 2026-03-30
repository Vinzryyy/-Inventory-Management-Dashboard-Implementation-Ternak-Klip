package models

import "time"

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	StockCount  int       `json:"stock_count"`
	LastUpdated time.Time `json:"last_updated"`
}

type ProductListResponse struct {
	Data []Product `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type HealthResponse struct {
	Status string `json:"status"`
}
