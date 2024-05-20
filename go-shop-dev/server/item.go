package server

import (
	"github.com/go-shop/modules/item/itemHandler"
	itemRePository "github.com/go-shop/modules/item/itemRepo"
	"github.com/go-shop/modules/item/itemUsecase"
)

func (s *server) itemService() {
	repo := itemRePository.NewIntemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/auth_v1")

	// Health Check
	item.GET("/health", s.healthCheckService)

}
