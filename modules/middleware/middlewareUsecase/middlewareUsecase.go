package middlewareUsecase

import (
	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/modules/middleware/middlewareRepository"
	"github.com/labstack/echo/v4"
)

type (
	MiddlewareUsecaseService interface{
		FirebaseAuthorization(c echo.Context, cfg *config.Config, token string) (echo.Context, error)
	}

	middlewareUsecase struct {
		middlewareRepository middlewareRepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
	return &middlewareUsecase{middlewareRepository}
}

func (u *middlewareUsecase) FirebaseAuthorization(c echo.Context, cfg *config.Config, token string) (echo.Context, error) {
	ctx := c.Request().Context()

	if err := u.middlewareRepository.VerifyToken(ctx, cfg.Grpc.AuthUrl, token); err != nil {
		return nil, err
	}

	return c, nil
}