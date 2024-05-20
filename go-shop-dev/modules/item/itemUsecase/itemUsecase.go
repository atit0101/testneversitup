package itemUsecase

import itemRePository "github.com/go-shop/modules/item/itemRepo"

type (
	ItemUsecaseService interface{}

	itemUsecase struct {
		itemRepository itemRePository.IntemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemRePository.IntemRepositoryService) ItemUsecaseService {

	return &itemUsecase{itemRepository}
}
