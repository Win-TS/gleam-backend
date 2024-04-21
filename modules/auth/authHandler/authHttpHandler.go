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
		FindUserByEmail(c echo.Context) error
		FindUserByPhoneNo(c echo.Context) error
		FindUserByUID(c echo.Context) error
		DeleteUser(c echo.Context) error
		UpdatePassword(c echo.Context) error
		VerifyToken(c echo.Context) error
		ManualFirebaseSignup(c echo.Context) error
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

	user, err := h.authUsecase.RegisterUserWithEmailPhoneAndPassword(ctx, req)
	if err != nil {
		log.Printf("Error - registering user: %v\n", err)
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, user)
}

func (h *authHttpHandler) FindUserByEmail(c echo.Context) error {
	ctx := context.Background()
	email := c.QueryParam("email")

	user, err := h.authUsecase.FindUserByEmail(ctx, email)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, user)
}

func (h *authHttpHandler) FindUserByPhoneNo(c echo.Context) error {
	ctx := context.Background()
	phoneNumber := ("+" + c.QueryParam("phone_no"))

	user, err := h.authUsecase.FindUserByPhoneNo(ctx, phoneNumber)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, user)
}

func (h *authHttpHandler) FindUserByUID(c echo.Context) error {
	ctx := context.Background()
	uid := c.QueryParam("uid")

	user, err := h.authUsecase.FindUserByUID(ctx, uid)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, user)
}

func (h *authHttpHandler) DeleteUser(c echo.Context) error {
	ctx := context.Background()
	uid := c.QueryParam("uid")

	err := h.authUsecase.DeleteUser(ctx, uid)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "user deleted")
}

func (h *authHttpHandler) UpdatePassword(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(auth.UpdatePasswordReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUsecase.UpdatePassword(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *authHttpHandler) VerifyToken(c echo.Context) error {
	ctx := context.Background()
	token := c.Request().Header.Get("Authorization")

	authToken, err := h.authUsecase.VerifyToken(ctx, token)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, authToken)
}

func (h *authHttpHandler) ManualFirebaseSignup(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(auth.RequestPayload)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUsecase.RegisterUserWithEmailPhoneAndPassword(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}