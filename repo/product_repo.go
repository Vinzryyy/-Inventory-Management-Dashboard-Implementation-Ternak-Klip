package repo

import (
	"context"
	"time"

	"inventory-dashboard/database"
	"inventory-dashboard/models"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]models.Product, error)
}

type productRepository struct {
	db *database.Postgres
}

func NewProductRepository(db *database.Postgres) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := r.db.DB.QueryContext(ctx, `
		SELECT id, name, stock_count, last_updated
		FROM products
		ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.StockCount, &product.LastUpdated); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
