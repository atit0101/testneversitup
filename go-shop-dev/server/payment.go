package server

import (
	paymentHandler "github.com/go-shop/modules/payment/paymentHendle"
	paymentRepository "github.com/go-shop/modules/payment/paymentRepo"
	"github.com/go-shop/modules/payment/paymentUsecase"
)

func (s *server) paymetnService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	grpcHandler := paymentHandler.NewPaymentGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	payment := s.app.Group("/auth_v1")

	// Health Check
	payment.GET("/health", s.healthCheckService)

}
