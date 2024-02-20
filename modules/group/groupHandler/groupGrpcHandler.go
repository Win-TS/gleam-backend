package groupHandler

import (
	"github.com/Win-TS/gleam-backend.git/modules/group/groupUsecase"
)

type (
	GroupGrpcHandlerService interface{}

	groupGrpcHandler struct {
		groupUsecase groupUsecase.GroupUsecaseService
	}
)


func NewGroupGrpcHandler(groupUsecase groupUsecase.GroupUsecaseService) GroupGrpcHandlerService {
	return &groupGrpcHandler{groupUsecase}
}