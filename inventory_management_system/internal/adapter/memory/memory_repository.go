package memory

import (
	"errors"

	"github.com/RhoNit/design_patterns_in_go/hexagonal_architecture/inventory_management_system/internal/core/domain"
)

type MemoryRepository struct {
	products map[int]domain.Product
	nextID   int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		products: make(map[int]domain.Product),
		nextID:   1,
	}
}

func (m *MemoryRepository) AddProduct(product domain.Product) error {
	product.ID = m.nextID
	m.products[product.ID] = product
	m.nextID++

	return nil
}

func (m *MemoryRepository) GetProduct(id int) (domain.Product, error) {
	product, exists := m.products[id]
	if !exists {
		return domain.Product{}, errors.New("product not found")
	}

	return product, nil
}

func (m *MemoryRepository) GetAllProducts() ([]domain.Product, error) {
	var newList []domain.Product

	for _, product := range m.products {
		newList = append([]domain.Product{product}, newList...)
	}

	return newList, nil
}

func (m *MemoryRepository) UpdateProduct(id int, product domain.Product) error {
	_, exists := m.products[id]
	if !exists {
		return errors.New("product not found")
	}

	// product.ID = id
	m.products[id] = product

	return nil
}

func (m *MemoryRepository) DeleteProduct(id int) error {
	if _, exists := m.products[id]; !exists {
		return errors.New("product not found")
	}

	delete(m.products, id)

	return nil
}
