package inventoryHandler

import (
	"github.com/go-shop/config"
	inventoryUsecase "github.com/go-shop/modules/inventory/inventotyUsecase"
)

type (
    InventoryQueueHandlerService interface{}

    inventoryQueueHandler struct {
      cfg *config.Config
      inventoryUsecase inventoryUsecase.InventoryUsecaseService
    }
)

func NewInventoryQueueHandler(cfg *config.Config , inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryQueueHandlerService  {
 return &inventoryQueueHandler{cfg , inventoryUsecase} 
}
