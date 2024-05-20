package playerHandler

import "github.com/go-shop/modules/player/playerUsecase"

type (
  playerGrpcHandlerService struct{
    playerUsecase playerUsecase.PlayerUsecaseService
  }
)

func NewPlayerGrpcHandler(playerUsecase playerUsecase.PlayerUsecaseService) playerGrpcHandlerService  {
    return playerGrpcHandlerService{playerUsecase} 
}
