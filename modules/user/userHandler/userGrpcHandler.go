package userHandler

import (
	"github.com/Win-TS/gleam-backend.git/modules/user/userUsecase"
)

type (
	UserGrpcHandlerService interface{}

	userGrpcHandler struct {
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserGrpcHandler(userUsecase userUsecase.UserUsecaseService) UserGrpcHandlerService {
	return userGrpcHandler{userUsecase}
}
