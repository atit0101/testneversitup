package paymentHandler

import (
	"github.com/go-shop/config"
	paymentUsecase "github.com/go-shop/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHendlerService interface{}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHendlerService {
	return &paymentHttpHandler{cfg, paymentUsecase}
}
