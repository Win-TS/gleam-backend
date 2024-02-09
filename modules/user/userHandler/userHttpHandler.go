package userHandler

import (
	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/modules/user/userUsecase"
)

type (
	UserHttpHandlerService interface{}

	userHttpHandler struct {
		cfg         *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserHttpHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserHttpHandlerService {
	return userHttpHandler{cfg, userUsecase}
}
