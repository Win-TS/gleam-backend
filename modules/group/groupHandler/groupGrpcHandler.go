package groupHandler

import (
	"context"

	groupPb "github.com/Win-TS/gleam-backend.git/modules/group/groupPb"
	"github.com/Win-TS/gleam-backend.git/modules/group/groupUsecase"
	"google.golang.org/protobuf/types/known/emptypb"
)

type groupGrpcHandler struct {
	groupPb.UnimplementedGroupGrpcServiceServer
	groupUsecase groupUsecase.GroupUsecaseService
}

func NewGroupGrpcHandler(groupUsecase groupUsecase.GroupUsecaseService) *groupGrpcHandler {
	return &groupGrpcHandler{groupUsecase: groupUsecase}
}

func (h *groupGrpcHandler) DeleteUserData(ctx context.Context, req *groupPb.DeleteUserDataReq) (*emptypb.Empty, error) {
	userID := req.GetUserId()

	err := h.groupUsecase.DeleteUserData(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}