package inventoryHandler

import (
	"github.com/go-shop/config"
	inventoryUsecase "github.com/go-shop/modules/inventory/inventotyUsecase"
)

type (
    InventoryHttpHandlerService interface{}

    inventoryHttpHandler struct {
      cfg *config.Config
      inventoryUsecase inventoryUsecase.InventoryUsecaseService
    }
)

func NewInventoryHttpHandler(cfg *config.Config , inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryHttpHandlerService  {
 return &inventoryHttpHandler{cfg , inventoryUsecase} 
}
