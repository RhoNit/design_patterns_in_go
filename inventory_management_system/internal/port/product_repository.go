package port

import "github.com/RhoNit/design_patterns_in_go/hexagonal_architecture/inventory_management_system/internal/core/domain"

type ProductRepository interface {
	AddProduct(product domain.Product) error
	GetProduct(id int) (domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	UpdateProduct(id int, product domain.Product) error
	DeleteProduct(id int) error
}
