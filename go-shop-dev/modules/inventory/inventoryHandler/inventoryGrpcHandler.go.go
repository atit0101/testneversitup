package inventoryHandler

import inventoryUsecase "github.com/go-shop/modules/inventory/inventotyUsecase"


type (
    inventoryGrpcHandler struct {
      inventoryUsecase inventoryUsecase.InventoryUsecaseService
    }
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandler  {
 return &inventoryGrpcHandler{inventoryUsecase} 
}
