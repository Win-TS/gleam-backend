package authHandler

import (
	"github.com/Win-TS/gleam-backend.git/modules/auth/authUsecase"
)

type (
	AuthGrpcHandlerService interface{}

	authGrpcHandler struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) AuthGrpcHandlerService {
	return &authGrpcHandler{authUsecase}
}