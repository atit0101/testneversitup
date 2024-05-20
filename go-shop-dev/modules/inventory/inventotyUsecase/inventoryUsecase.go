package inventoryUsecase

import inventoryRepository "github.com/go-shop/modules/inventory/inventoryRepo"

type(
  InventoryUsecaseService interface{}
  inventoryUsecase struct {
    inventoryRepository inventoryRepository.InventoryRepositoryService
  }
)

func NewInventoryUsecase(inventoryRepository inventoryRepository.InventoryRepositoryService) InventoryUsecaseService{
 return &inventoryUsecase{inventoryRepository} 
}
