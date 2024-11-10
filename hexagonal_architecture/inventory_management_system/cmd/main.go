package main

import (
	"github.com/RhoNit/design_patterns_in_go/hexagonal_architecture/inventory_management_system/internal/adapter/handler"
	"github.com/RhoNit/design_patterns_in_go/hexagonal_architecture/inventory_management_system/internal/adapter/memory"
	"github.com/RhoNit/design_patterns_in_go/hexagonal_architecture/inventory_management_system/internal/core/service"
)

func main() {
	memoryRepo := memory.NewMemoryRepository()
	inventoryService := service.NewInventoryService(memoryRepo)
	inventoryHandler := handler.NewInventoryHandler(inventoryService)
	inventoryHandler.Run()
}
