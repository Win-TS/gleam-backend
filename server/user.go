package server

import (
	"log"

	"github.com/Win-TS/gleam-backend.git/modules/user/userHandler"
	"github.com/Win-TS/gleam-backend.git/modules/user/userUsecase"
	userdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/userdb/sqlc"
)

func (s *server) userService() {
	postgresDB, ok := s.db.(userdb.Store)
	if !ok {
		log.Fatal("Unsupported database type")
		return
	}
	usecase := userUsecase.NewUserUsecase(postgresDB, s.storage)
	httpHandler := userHandler.NewUserHttpHandler(s.cfg, usecase)
	grpcHandler := userHandler.NewUserGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	user := s.app.Group("/user_v1")

	// Health Check
	user.GET("", s.healthCheckService)

	user.POST("/createuser", httpHandler.RegisterNewUser)
	user.GET("/userprofile", httpHandler.GetUserProfile)
	user.POST("/uploaduserphoto", httpHandler.UploadProfilePhoto)
}
