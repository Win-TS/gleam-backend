package userHandler

import (
	"context"
	"errors"

	userPb "github.com/Win-TS/gleam-backend.git/modules/user/userPb"
	"github.com/Win-TS/gleam-backend.git/modules/user/userUsecase"
)

type (
	userGrpcHandler struct {
		userPb.UnimplementedUserGrpcServiceServer
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserGrpcHandler(userUsecase userUsecase.UserUsecaseService) *userGrpcHandler {
	return &userGrpcHandler{userUsecase: userUsecase}
}

func (g *userGrpcHandler) SearchUser(ctx context.Context, req *userPb.SearchUserReq) (*userPb.SearchUserRes, error) {
	_, err := g.userUsecase.GetUserInfo(ctx, int(req.UserId))
	if err != nil {
		return &userPb.SearchUserRes{
			UserId: req.UserId,
			Valid:  false,
		}, errors.New("error: userId not found")
	}

	return &userPb.SearchUserRes{
		UserId: req.UserId,
		Valid:  true,
	}, nil
}
