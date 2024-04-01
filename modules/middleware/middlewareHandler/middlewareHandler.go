package middlewareHandler

import (
	"net/http"

	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/modules/middleware/middlewareUsecase"
	"github.com/Win-TS/gleam-backend.git/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	MiddlewareHandlerService interface{
		FirebaseAuthorization(next echo.HandlerFunc) echo.HandlerFunc
	}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{cfg, middlewareUsecase}
}

func (h *middlewareHandler) FirebaseAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return response.ErrResponse(c, http.StatusUnauthorized, "Token is required")
		}

		newCtx, err := h.middlewareUsecase.FirebaseAuthorization(c, h.cfg, token)
		if err != nil {
			return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
		}

		return next(newCtx)
	}
}