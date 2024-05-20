package server

import (
	playerHandler "github.com/go-shop/modules/player/playerHendle"
	playerRepository "github.com/go-shop/modules/player/playerRepo"
	"github.com/go-shop/modules/player/playerUsecase"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/auth_v1")

	// Health Check
	player.GET("/health", s.healthCheckService)

}
