package service

import (
	"context"

	"inventory-dashboard/models"
	"inventory-dashboard/repo"
)

type ProductService interface {
	ListProducts(ctx context.Context) ([]models.Product, error)
}

type productService struct {
	repo repo.ProductRepository
}

func NewProductService(repository repo.ProductRepository) ProductService {
	return &productService{repo: repository}
}

func (s *productService) ListProducts(ctx context.Context) ([]models.Product, error) {
	return s.repo.GetAll(ctx)
}
