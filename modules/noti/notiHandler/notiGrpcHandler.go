package notiHandler

import "github.com/Win-TS/gleam-backend.git/modules/noti/notiUsecase"

type (
	notiGrpcHandler struct {
		//notiPb.UnimplementedNotiGrpcServiceServer
		notiUsecase notiUsecase.NotiUsecaseService
	}
)

func NewNotiGrpcHandler(notiUsecase notiUsecase.NotiUsecaseService) *notiGrpcHandler {
	return &notiGrpcHandler{notiUsecase: notiUsecase}
}
