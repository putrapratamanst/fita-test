package product

import (
	"fita-test/interface/product"
	"fita-test/model"
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


func (s *Service)Create(input model.Add){
	s.repo.Insert(input)
}

func (s *Service)AddToCart(input model.Cart){
	s.repo.SaveCart(input)
}
