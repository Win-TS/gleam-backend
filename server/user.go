package server

import (
	"github.com/Win-TS/gleam-backend.git/modules/user/userHandler"
	"github.com/Win-TS/gleam-backend.git/modules/user/userUsecase"
)

func (s *server) userService() {
	usecase := userUsecase.NewUserUsecase()
	httpHandler := userHandler.NewUserHttpHandler(s.cfg, usecase)
	grpcHandler := userHandler.NewUserGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	user := s.app.Group("/user_v1")

	// Health Check
	user.GET("", s.healthCheckService)
}
