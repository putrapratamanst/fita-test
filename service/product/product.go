package product

import (
	"fita-test/interface/product"
)

// Service interface
type Service struct {
	repo product.IProductRepository
}

// NewService create new use case
func NewService(repo product.IProductRepository) *Service {
	return &Service{
		repo: repo,
	}
}
