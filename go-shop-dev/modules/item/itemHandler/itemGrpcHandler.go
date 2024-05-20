package itemHandler

import "github.com/go-shop/modules/item/itemUsecase"

type (
	itemGrpcHandler struct {
		itemGrpcUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemGrpcHandler(itemGrpcUsecase itemUsecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{itemGrpcUsecase}
}
