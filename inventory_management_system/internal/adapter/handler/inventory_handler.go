package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/RhoNit/design_patterns_in_go/hexagonal_architecture/inventory_management_system/internal/core/domain"
	"github.com/RhoNit/design_patterns_in_go/hexagonal_architecture/inventory_management_system/internal/core/service"
)

type InventoryHandler struct {
	service *service.InventoryService
}

func NewInventoryHandler(service *service.InventoryService) *InventoryHandler {
	return &InventoryHandler{
		service: service,
	}
}

func (h *InventoryHandler) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Press\n")
		fmt.Println("1. to ADD product into the inventory")
		fmt.Println("2. to GET the product")
		fmt.Println("3. to GET all products")
		fmt.Println("4. to DELETE a product from the inventory")

		fmt.Println("Enter your choice: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			h.AddProduct(scanner)
		case "2":
			h.GetProduct(scanner)
		// case "3":
		// case "4":
		default:
			return
		}

	}
}

func (h *InventoryHandler) AddProduct(scanner *bufio.Scanner) {
	fmt.Println("Enter the product name: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	fmt.Println("Enter the product price: ")
	scanner.Scan()
	price, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

	fmt.Println("Enter product stock: ")
	scanner.Scan()
	stock, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	product := domain.Product{Name: name, Price: price, Stock: stock}
	err := h.service.AddProduct(product)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Product added successfully")
	}
}

func (h *InventoryHandler) GetProduct(scanner *bufio.Scanner) {
	fmt.Println("Enter the ID of the product: ")
	scanner.Scan()
	productID, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	if foundProduct, err := h.service.GetProduct(productID); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Product found: %+v", foundProduct)
	}
}

// func (h *InventoryHandler) GetAllProducts() ([]domain.Product, error) {
// 	return nil, nil
// }

// func (h *InventoryHandler) UpdateProduct() error {
// 	return nil
// }

// func (h *InventoryHandler) DeleteProduct() error {
// 	return nil
// }
