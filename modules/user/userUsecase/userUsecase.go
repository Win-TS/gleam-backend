package userUsecase

import (
	"context"
	"database/sql"
	"io"
	"time"

	"firebase.google.com/go/storage"
	"github.com/Win-TS/gleam-backend.git/modules/user"
	userdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/userdb/sqlc"
	"github.com/Win-TS/gleam-backend.git/pkg/utils"
)

type UserUsecaseService interface {
	GetUserProfile(pctx context.Context, id int) (user.UserProfile, error)
	GetLatestId(pctx context.Context) (int, error)
	RegisterNewUser(pctx context.Context, payload *user.NewUserReq, photoUrl string) (userdb.User, error)
	SaveToFirebaseStorage(pctx context.Context, bucketName, objectPath, filename string, file io.Reader) (string, error)
	GetUserInfo(pctx context.Context, id int) (userdb.User, error)
	EditUsername(pctx context.Context, args userdb.ChangeUsernameParams) (user.UserProfile, error)
	EditPhoneNumber(pctx context.Context, args userdb.ChangePhoneNoParams) (userdb.User, error)
	DeleteUser(pctx context.Context, id int) error
}

type userUsecase struct {
	store userdb.Store
	storageClient *storage.Client
}

func NewUserUsecase(store userdb.Store, storageClient *storage.Client) UserUsecaseService {
	return &userUsecase{store, storageClient}
}

func (u *userUsecase) GetUserInfo(pctx context.Context, id int) (userdb.User, error) {
	userData, err := u.store.GetUser(pctx, int32(id))
	if err != nil {
		return userdb.User{}, err
	}

	return userData, nil
}

func (u *userUsecase) GetUserProfile(pctx context.Context, id int) (user.UserProfile, error) {
	userData, err := u.store.GetUser(pctx, int32(id))
	if err != nil {
		return user.UserProfile{}, err
	}

	userID := sql.NullInt32{Int32: int32(id), Valid: true}

	userFriendCount, err := u.store.GetFriendsCountByID(pctx, userID)
	if err != nil {
		return user.UserProfile{}, err
	}

	return user.UserProfile{
		Username:     userData.Username,
		Firstname:    userData.Firstname,
		Lastname:     userData.Lastname,
		FriendsCount: int(userFriendCount),
		PhotoUrl:     userData.Photourl.String,
	}, nil
}

func (u *userUsecase) GetLatestId(pctx context.Context) (int, error) {
	latestId, err := u.store.GetLatestId(pctx)
	if err != nil {
		return 0, err
	}

	return int(latestId + 1), nil
}

func (u *userUsecase) RegisterNewUser(pctx context.Context, payload *user.NewUserReq, photoUrl string) (userdb.User, error) {

	birthdayTime, err := time.Parse("2006-01-02", payload.Birthday)
	if err != nil {
		return userdb.User{}, err
	}

	sqlPhotoUrl := utils.ConvertStringToSqlNullString(photoUrl)

	return u.store.CreateUser(pctx, userdb.CreateUserParams{
		Username:    payload.Username,
		Firstname:   payload.Firstname,
		Lastname:    payload.Lastname,
		PhoneNo:     payload.PhoneNo,
		Email:       payload.Email,
		Nationality: payload.Nationality,
		Age:         int32(payload.Age),
		Birthday:    birthdayTime,
		Gender:      payload.Gender,
		Photourl:    sqlPhotoUrl,
	})
}

func (u *userUsecase) SaveToFirebaseStorage(pctx context.Context, bucketName, objectPath, filename string, file io.Reader) (string, error) {
	bucket, _ := u.storageClient.Bucket(bucketName)
	obj := bucket.Object(objectPath + "/" + filename)

	wc := obj.NewWriter(pctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	url := "https://firebasestorage.googleapis.com/v0/b/" + bucketName + "/o/" + objectPath + "%" + "2F" + filename + "?alt=media"

	return url, nil
}

func (u *userUsecase) EditUsername(pctx context.Context, args userdb.ChangeUsernameParams) (user.UserProfile, error) {
	if err := u.store.ChangeUsername(pctx, args); err != nil {
		return user.UserProfile{}, err
	}
	return u.GetUserProfile(pctx, int(args.ID))
}

func (u *userUsecase) EditPhoneNumber(pctx context.Context, args userdb.ChangePhoneNoParams) (userdb.User, error) {
	if err := u.store.ChangePhoneNo(pctx, args); err != nil {
		return userdb.User{}, err
	}
	return u.GetUserInfo(pctx, int(args.ID))
}

func (u *userUsecase) DeleteUser(pctx context.Context, id int) error {
	if err := u.store.DeleteUser(pctx, int32(id)); err != nil {
		return err
	}
	return nil
}