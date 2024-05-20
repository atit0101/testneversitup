package paymentUsecase

import paymentRepository "github.com/go-shop/modules/payment/paymentRepo"

type (
  PaymentUsecaseService interface{}

  paymentUsecase struct {
    paymentRepository paymentRepository.PaymentRepositoryService
  }
)

func NewPaymentUsecase(paymentRepository paymentRepository.PaymentRepositoryService) PaymentUsecaseService  {
 return &paymentUsecase{paymentRepository} 
}
