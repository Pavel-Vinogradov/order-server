package service

import (
	"order-server/core/repository"
	"order-server/internal/product/entity"
)

type ProductService struct {
	repo repository.Repository[entity.Product]
}

func NewProductService(repo repository.Repository[entity.Product]) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(p entity.Product) (entity.Product, error) {
	return s.repo.Create(p)
}

func (s *ProductService) UpdateProduct(p entity.Product) (entity.Product, error) {
	return s.repo.Update(p)
}

func (s *ProductService) DeleteProduct(p entity.Product) (bool, error) {
	return s.repo.Delete(p)
}

func (s *ProductService) GetAllProducts() ([]entity.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetProductByID(id int64) (entity.Product, error) {
	return s.repo.FindByID(id)
}
