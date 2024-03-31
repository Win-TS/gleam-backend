package authUsecase

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/Win-TS/gleam-backend.git/modules/auth/authRepository"
)

type (
	AuthUsecaseService interface{
		RegisterUserWithEmailPhoneAndPassword(pctx context.Context, email, phoneNumber, password string) (*auth.UserRecord, error)
		VerifyToken(pctx context.Context, token string) (*auth.Token, error)
		FindUserByEmail(pctx context.Context, email string) (*auth.UserRecord, error)
		FindUserByUID(pctx context.Context, uid string) (*auth.UserRecord, error)
		FindUserByPhoneNo(pctx context.Context, uid string) (*auth.UserRecord, error)
		DeleteUser(pctx context.Context, uid string) error
	}

	authUsecase struct{
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository}
}

func (u *authUsecase) RegisterUserWithEmailPhoneAndPassword(pctx context.Context, email, phoneNumber, password string) (*auth.UserRecord, error) {
	return u.authRepository.CreateUserWithEmailPhoneAndPassword(pctx, email, phoneNumber, password)
}

func (u *authUsecase) VerifyToken(pctx context.Context, token string) (*auth.Token, error) {
	return u.authRepository.VerifyToken(pctx, token)
}

func (u *authUsecase) FindUserByEmail(pctx context.Context, email string) (*auth.UserRecord, error) {
	return u.authRepository.FindUserByEmail(pctx, email)
}

func (u *authUsecase) FindUserByUID(pctx context.Context, uid string) (*auth.UserRecord, error) {
	return u.authRepository.FindUserByUID(pctx, uid)
}

func (u *authUsecase) FindUserByPhoneNo(pctx context.Context, phoneNo string) (*auth.UserRecord, error) {
	return u.authRepository.FindUserByPhoneNo(pctx, phoneNo)
}

func (u *authUsecase) DeleteUser(pctx context.Context, uid string) error {
	return u.authRepository.DeleteUser(pctx, uid)
}