package server

import (
	"log"

	firebase "firebase.google.com/go"
	authPb "github.com/Win-TS/gleam-backend.git/modules/auth/authPb"
	"github.com/Win-TS/gleam-backend.git/modules/auth/authHandler"
	"github.com/Win-TS/gleam-backend.git/modules/auth/authRepository"
	"github.com/Win-TS/gleam-backend.git/modules/auth/authUsecase"
	"github.com/Win-TS/gleam-backend.git/pkg/grpcconn"
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

	// gRPC
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(s.cfg, s.cfg.Grpc.AuthUrl)
		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)
		log.Printf("Player gRPC server listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	auth.GET("", s.healthCheckService)
	auth.POST("/signup", httpHandler.RegisterUser)
	auth.GET("/find/email", httpHandler.FindUserByEmail)
	auth.GET("/find/phone", httpHandler.FindUserByPhoneNo)
	auth.GET("/find/uid", httpHandler.FindUserByUID)
	auth.DELETE("/delete", httpHandler.DeleteUser)
	auth.PUT("/update-password", httpHandler.UpdatePassword)
}
