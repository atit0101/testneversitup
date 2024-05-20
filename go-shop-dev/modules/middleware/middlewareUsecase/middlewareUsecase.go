package middlewareUsecase

import "github.com/go-shop/modules/middleware/middlewareRepository"

type (
  MiddlewareUsecaseService interface{}

  middlewareUsecase struct {
    middlewareRepository middlewareRepository.MiddlewareRepositoryService
  }
)

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
  return &middlewareUsecase{middlewareRepository}
}
