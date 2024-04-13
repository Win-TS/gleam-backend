package notiUsecase

import (
	"github.com/Win-TS/gleam-backend.git/modules/noti/notiRepository"
)

type (
	NotiUsecaseService interface{
	}

	notiUsecase struct{
		notiRepository notiRepository.NotiRepositoryService
	}
)

func NewNotiUsecase(notiRepository notiRepository.NotiRepositoryService) NotiUsecaseService {
	return &notiUsecase{notiRepository}
}