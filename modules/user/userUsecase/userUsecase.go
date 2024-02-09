package userUsecase

import (
	"context"

	userdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/userdb/sqlc"
)

type UserUsecaseService interface {
	GetUserProfile(ctx context.Context, id int) (userdb.User, error)
}

type userUsecase struct {
	userdb.Querier
}

func NewUserUsecase() UserUsecaseService {
	return &userUsecase{}
}

func (u *userUsecase) GetUserProfile(ctx context.Context, id int) (userdb.User, error) {
	return u.GetUser(ctx, int32(id))
}
