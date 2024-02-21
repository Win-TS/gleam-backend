package authHandler

import (
	"context"
	"log"
	"net/http"

	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/modules/auth"
	"github.com/Win-TS/gleam-backend.git/modules/auth/authUsecase"
	"github.com/Win-TS/gleam-backend.git/pkg/request"
	"github.com/Win-TS/gleam-backend.git/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	AuthHttpHandlerService interface {
		RegisterUser(c echo.Context) error
		//VerifyToken(c echo.Context) error
		FindUserByEmail(c echo.Context) error
		FindUserByPhoneNo(c echo.Context) error
		FindUserByUID(c echo.Context) error
	}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{cfg, authUsecase}
}

func (h *authHttpHandler) RegisterUser(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(auth.RequestPayload)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	user, err := h.authUsecase.RegisterUserWithEmailPhoneAndPassword(ctx, req.Email, req.PhoneNumber, req.Password)
	if err != nil {
		log.Printf("Error - registering user: %v\n", err)
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, user)
}

// func (h *authHttpHandler) VerifyToken(c echo.Context) error {
// 	ctx := context.Background()

// 	wrapper := request.ContextWrapper(c)
// 	token := wrapper.GetAuthorizationHeader()

// 	if token == "" {
// 		return response.ErrResponse(c, http.StatusUnauthorized, "error: token not in header")
// 	}

// 	authToken, err := h.authUsecase.VerifyToken(ctx, token)
// 	if err != nil {
// 		log.Printf("Error - verifying token: %v\n", err)
// 		return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
// 	}

// 	return response.SuccessResponse(c, http.StatusOK, authToken)
// }

func (h *authHttpHandler) FindUserByEmail(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(auth.EmailCheck)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	user, err := h.authUsecase.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, user)
}

func (h *authHttpHandler) FindUserByPhoneNo(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(auth.PhoneCheck)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	user, err := h.authUsecase.FindUserByEmail(ctx, req.PhoneNumber)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, user)
}

func (h *authHttpHandler) FindUserByUID(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(auth.UIDCheck)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	user, err := h.authUsecase.FindUserByEmail(ctx, req.UID)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, user)
}