package server

import (
	"github.com/go-shop/modules/inventory/inventoryHandler"
	inventoryRepository "github.com/go-shop/modules/inventory/inventoryRepo"
	inventoryUsecase "github.com/go-shop/modules/inventory/inventotyUsecase"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	inventory := s.app.Group("/auth_v1")

	// Health Check
	inventory.GET("/health", s.healthCheckService)

}
