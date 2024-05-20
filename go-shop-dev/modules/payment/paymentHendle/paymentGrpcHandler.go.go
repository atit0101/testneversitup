package paymentHandler

import paymentUsacase "github.com/go-shop/modules/payment/paymentUsecase"

type (
  paymentGrpcHandler struct {
    paymentUsecase paymentUsacase.PaymentUsecaseService
  }
)

func NewPaymentGrpcHandler(paymentUsecase paymentUsacase.PaymentUsecaseService) *paymentGrpcHandler{
  return &paymentGrpcHandler{paymentUsecase}
}
