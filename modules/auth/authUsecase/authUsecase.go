package authUsecase

import "github.com/Win-TS/gleam-backend.git/modules/auth/authRepository"

type (
	AuthUsecaseService interface{}

	authUsecase struct{
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository}
}