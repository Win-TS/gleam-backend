package notiHandler

import (
	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/modules/noti/notiUsecase"
)

type (
	NotiHttpHandlerService interface {
	}

	notiHttpHandler struct {
		cfg         *config.Config
		notiUsecase notiUsecase.NotiUsecaseService
	}
)

func NewNotiHttpHandler(cfg *config.Config, notiUsecase notiUsecase.NotiUsecaseService) NotiHttpHandlerService {
	return &notiHttpHandler{cfg, notiUsecase}
}
