package groupHandler

import (
	//"context"

	//groupPb "github.com/Win-TS/gleam-backend.git/modules/group/groupPb"
	"github.com/Win-TS/gleam-backend.git/modules/group/groupUsecase"
)

type (
	groupGrpcHandler struct {
		//groupPb.UnimplementedGroupServiceServer
		groupUsecase groupUsecase.GroupUsecaseService
	}
)

func NewGroupGrpcHandler(groupUsecase groupUsecase.GroupUsecaseService) *groupGrpcHandler {
	return &groupGrpcHandler{groupUsecase}
}
