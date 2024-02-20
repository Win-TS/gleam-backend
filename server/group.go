package server

import (
	"log"

	"github.com/Win-TS/gleam-backend.git/modules/group/groupHandler"
	"github.com/Win-TS/gleam-backend.git/modules/group/groupUsecase"
	groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
)

func (s *server) groupService() {
	postgresDB, ok := s.db.(groupdb.Store)
	if !ok {
		log.Fatal("Unsupported database type")
		return
	}
	usecase := groupUsecase.NewGroupUsecase(postgresDB, s.storage)
	httpHandler := groupHandler.NewGroupHttpHandler(s.cfg, usecase)
	grpcHandler := groupHandler.NewGroupGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	group := s.app.Group("/group_v1")

	// Health Check
	group.GET("", s.healthCheckService)

}
