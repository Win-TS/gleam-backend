package userHandler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/modules/user"
	"github.com/Win-TS/gleam-backend.git/modules/user/userUsecase"
	"github.com/Win-TS/gleam-backend.git/pkg/request"
	"github.com/Win-TS/gleam-backend.git/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	UserHttpHandlerService interface {
		GetUserProfile(c echo.Context) error
		RegisterNewUser(c echo.Context) error
		UploadProfilePhoto(c echo.Context) error
	}

	userHttpHandler struct {
		cfg         *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserHttpHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserHttpHandlerService {
	return &userHttpHandler{cfg, userUsecase}
}

func (h *userHttpHandler) GetUserProfile(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return err
	}

	res, err := h.userUsecase.GetUserProfile(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *userHttpHandler) RegisterNewUser(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(user.NewUserReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.RegisterNewUser(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *userHttpHandler) UploadProfilePhoto(c echo.Context) error {
	file, err := c.FormFile("photo")
	if err != nil {
		return c.String(http.StatusBadRequest, "Error: retrieving the file")
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: opening the file")
	}
	defer src.Close()

	bucketName := h.cfg.Firebase.StorageBucket
	objectPath := "userprofile"

	url, err := h.userUsecase.SaveToFirebaseStorage(c.Request().Context(), bucketName, objectPath, file.Filename, src)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: saving the file to Firebase Storage")
	}

	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully to Firebase Storage: %s", file.Filename, url))
}