package server

import (
	"github.com/go-shop/modules/auth/authHandler"
	authRepository "github.com/go-shop/modules/auth/authRepo"
	"github.com/go-shop/modules/auth/authUsecase"
)

func (s *server) authServiec() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("/health", s.healthCheckService)

}
