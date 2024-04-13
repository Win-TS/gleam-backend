package server

import (
	"log"

	firebase "firebase.google.com/go"
	"github.com/Win-TS/gleam-backend.git/modules/noti/notiHandler"
	"github.com/Win-TS/gleam-backend.git/modules/noti/notiRepository"
	"github.com/Win-TS/gleam-backend.git/modules/noti/notiUsecase"
)

func (s *server) notiService() {
	firebaseDB, ok := s.db.(*firebase.App)
	if !ok {
		log.Fatal("Unsupported database type")
		return
	}
	repo := notiRepository.NewNotiRepository(firebaseDB)
	usecase := notiUsecase.NewNotiUsecase(repo)
	httpHandler := notiHandler.NewNotiHttpHandler(s.cfg, usecase)
	grpcHandler := notiHandler.NewNotiGrpcHandler(usecase)

	// gRPC
	// go func() {
	// 	grpcServer, lis := grpcconn.NewGrpcServer(s.cfg, s.cfg.Grpc.NotiUrl)
	// 	notiPb.RegisterNotiGrpcServiceServer(grpcServer, grpcHandler)
	// 	log.Printf("Player gRPC server listening on %s", s.cfg.Grpc.NotiUrl)
	// 	grpcServer.Serve(lis)
	// }()

	_ = httpHandler
	_ = grpcHandler

	noti := s.app.Group("/noti_v1")

	noti.GET("", s.healthCheckService)
}