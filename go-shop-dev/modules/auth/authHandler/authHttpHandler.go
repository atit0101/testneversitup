package authHandler

import (
	"github.com/go-shop/config"
	"github.com/go-shop/modules/auth/authUsecase"
)

type (
	AuthHttpHendlerService interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHendlerService {
	return &authHttpHandler{cfg, authUsecase}
}
