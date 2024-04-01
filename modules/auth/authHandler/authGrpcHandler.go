package authHandler

import (
	"context"

	authPb "github.com/Win-TS/gleam-backend.git/modules/auth/authPb"
	"github.com/Win-TS/gleam-backend.git/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandler {
	return &authGrpcHandler{authUsecase: authUsecase}
}

func (g *authGrpcHandler) DeleteUser(ctx context.Context, req *authPb.DeleteUserReq) (*authPb.DeleteUserRes, error) {
	err := g.authUsecase.DeleteUser(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &authPb.DeleteUserRes{
		Uid: req.Uid,
		Success: true,
	}, nil
}

func (g *authGrpcHandler) GetUidFromEmail(ctx context.Context, req *authPb.GetUidFromEmailReq) (*authPb.GetUidFromEmailRes, error) {
	user, err := g.authUsecase.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	
	return &authPb.GetUidFromEmailRes{
		Uid: user.UserInfo.UID,
		Email: user.UserInfo.Email,
	}, nil
}

func (g *authGrpcHandler) VerifyToken(ctx context.Context, req *authPb.VerifyTokenReq) (*authPb.VerifyTokenRes, error) {
	token, err := g.authUsecase.VerifyToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	
	return &authPb.VerifyTokenRes{
		Uid: token.UID,
		Success: true,
	}, nil
}