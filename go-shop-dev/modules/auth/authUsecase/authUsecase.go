package authUsecase

import authRepository "github.com/go-shop/modules/auth/authRepo"


type (
  AuthUsecaseService interface{}

  authUsecase struct {
    authRepository authRepository.AuthRepositoryService 
  }

)

func NewAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
  return &authUsecase{authRepository} 
}
