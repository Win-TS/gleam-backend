package server

import (
	"log"

	firebase "firebase.google.com/go"
	"github.com/Win-TS/gleam-backend.git/modules/auth/authHandler"
	"github.com/Win-TS/gleam-backend.git/modules/auth/authRepository"
	"github.com/Win-TS/gleam-backend.git/modules/auth/authUsecase"
)

func (s *server) authService() {
	firebaseDB, ok := s.db.(*firebase.App)
	if !ok {
		log.Fatal("Unsupported database type")
		return
	}
	repo := authRepository.NewAuthRepository(firebaseDB)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	// // gRPC
	// go func() {
	// 	grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)
	// 	authPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)
	// 	log.Printf("Player gRPC server listening on %s", s.cfg.Grpc.AuthUrl)
	// 	grpcServer.Serve(lis)
	// }()

	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	auth.GET("", s.healthCheckService)
	auth.POST("/signup", httpHandler.RegisterUser)
	auth.GET("/verify", httpHandler.VerifyToken)
	auth.POST("/find/email", httpHandler.FindUserByEmail)
	auth.POST("/find/phone", httpHandler.FindUserByPhoneNo)
	auth.POST("/find/uid", httpHandler.FindUserByUID)

}
