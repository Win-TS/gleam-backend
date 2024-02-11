package userHandler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/modules/user"
	"github.com/Win-TS/gleam-backend.git/modules/user/userUsecase"
	userdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/userdb/sqlc"
	"github.com/Win-TS/gleam-backend.git/pkg/request"
	"github.com/Win-TS/gleam-backend.git/pkg/response"
	"github.com/Win-TS/gleam-backend.git/pkg/utils"
	"github.com/labstack/echo/v4"
)

type (
	UserHttpHandlerService interface {
		GetUserProfile(c echo.Context) error
		RegisterNewUser(c echo.Context) error
		GetUserInfo(c echo.Context) error
		UploadProfilePhoto(c echo.Context) error
		EditUsername(c echo.Context) error
		EditPhoneNumber(c echo.Context) error
		DeleteUser(c echo.Context) error
	}

	userHttpHandler struct {
		cfg         *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserHttpHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserHttpHandlerService {
	return &userHttpHandler{cfg, userUsecase}
}

// GetUserProfile returns a response payload showing data for user profile, using request parameter "user_id".
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

// RegisterNewUser saves user data from CreateUserReq payload and returns a response payload of data for the user, "photo" parameter can be added.
func (h *userHttpHandler) RegisterNewUser(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(user.NewUserReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	file, _ := c.FormFile("photo")
	var url string

	fileid, err := h.userUsecase.GetLatestId(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}
	fileidStr := strconv.Itoa(fileid)

	if file != nil {
		src, err := file.Open()
		if err != nil {
			return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		bucketName := h.cfg.Firebase.StorageBucket
		objectPath := "userprofile"

		url, err = h.userUsecase.SaveToFirebaseStorage(c.Request().Context(), bucketName, objectPath, (fileidStr + utils.GetFileExtension(file.Filename)), src)
		if err != nil {
			return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	res, err := h.userUsecase.RegisterNewUser(ctx, req, url)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

// GetUserInfo returns a full information payload of user from user_id query param.
func (h *userHttpHandler) GetUserInfo(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.GetUserInfo(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

// UploadProfilePhoto uploads a file to userprofile bucket storage, and returns the link of the photo.
func (h *userHttpHandler) UploadProfilePhoto(c echo.Context) error {
	file, err := c.FormFile("photo")
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	bucketName := h.cfg.Firebase.StorageBucket
	objectPath := "userprofile"

	url, err := h.userUsecase.SaveToFirebaseStorage(c.Request().Context(), bucketName, objectPath, file.Filename, src)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, fmt.Sprintf("File %s uploaded successfully to Firebase Storage: %s", file.Filename, url))
}

// EditUsername updates the database with new username from query parameter of user_id and new_username.
func (h *userHttpHandler) EditUsername(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	args := userdb.ChangeUsernameParams{
		ID:       int32(userId),
		Username: c.QueryParam("new_username"),
	}

	res, err := h.userUsecase.EditUsername(ctx, args)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

// EditPhoneNumber updates the database with new phone number from query parameter of user_id and new_phone_no.
func (h *userHttpHandler) EditPhoneNumber(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	args := userdb.ChangePhoneNoParams{
		ID:      int32(userId),
		PhoneNo: c.QueryParam("new_phone_no"),
	}

	res, err := h.userUsecase.EditPhoneNumber(ctx, args)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

// DeleteUser deletes a user from the database by the user_id query parameter.
func (h *userHttpHandler) DeleteUser(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.userUsecase.DeleteUser(ctx, userId); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, fmt.Sprintf("Successfully deleted user id: %v", userId))
}
