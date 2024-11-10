package service

import (
	"github.com/RhoNit/design_patterns_in_go/hexagonal_architecture/inventory_management_system/internal/core/domain"
	"github.com/RhoNit/design_patterns_in_go/hexagonal_architecture/inventory_management_system/internal/port"
)

type InventoryService struct {
	repo port.ProductRepository
}

func NewInventoryService(repo port.ProductRepository) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}

func (s *InventoryService) AddProduct(product domain.Product) error {
	return s.repo.AddProduct(product)
}

func (s *InventoryService) GetProduct(id int) (domain.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *InventoryService) GetAllProducts() ([]domain.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *InventoryService) UpdateProduct(id int, product domain.Product) error {
	return s.repo.UpdateProduct(id, product)
}

func (s *InventoryService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
