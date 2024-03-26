package userUsecase

import (
	"context"
	"database/sql"
	//"errors"
	"fmt"
	"io"
	//"log"
	"math/rand"
	"time"

	"strconv"

	"firebase.google.com/go/storage"
	"github.com/Win-TS/gleam-backend.git/modules/user"
	userdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/userdb/sqlc"
	//userPb "github.com/Win-TS/gleam-backend.git/modules/user/userPb"
	//"github.com/Win-TS/gleam-backend.git/pkg/grpcconn"
	"github.com/Win-TS/gleam-backend.git/pkg/utils"
	"github.com/jaswdr/faker"
)

type UserUsecaseService interface {
	GetUserProfile(pctx context.Context, id int) (user.UserProfile, error)
	GetBatchUserProfiles(pctx context.Context, ids []int32) ([]userdb.GetBatchUserProfilesRow, error)
	GetUserProfileByEmail(pctx context.Context, email string) (user.UserProfile, error)
	GetUserProfileByUsername(pctx context.Context, username string) (user.UserProfile, error)
	GetLatestId(pctx context.Context) (int, error)
	RegisterNewUser(pctx context.Context, payload *user.NewUserReq, photoUrl string) (userdb.User, error)
	SaveToFirebaseStorage(pctx context.Context, bucketName, objectPath, filename string, file io.Reader) (string, error)
	GetUserInfo(pctx context.Context, id int) (userdb.User, error)
	GetUserInfoByEmail(pctx context.Context, email string) (userdb.User, error)
	GetUserInfoByUsername(pctx context.Context, username string) (userdb.User, error)
	EditUsername(pctx context.Context, args userdb.ChangeUsernameParams) (user.UserProfile, error)
	EditPhoneNumber(pctx context.Context, args userdb.ChangePhoneNoParams) (userdb.User, error)
	EditName(ctx context.Context, userID int32, firstName, lastName string) (user.UserProfile, error)
	DeleteUser(pctx context.Context, id int) error
	FriendInfo(ctx context.Context, args userdb.GetFriendParams) ([]userdb.Friend, error)
	FriendListById(pctx context.Context, id int) ([]userdb.ListFriendsByUserIdRow, error)
	FriendsCount(pctx context.Context, userId1 sql.NullInt32) (int64, error)
	FriendsRequestedList(pctx context.Context, userId1 sql.NullInt32) ([]userdb.User, error)
	FriendsPendingList(pctx context.Context, userId2 sql.NullInt32) ([]userdb.User, error)
	AddFriend(pctx context.Context, args user.CreateFriendReq) (userdb.Friend, error)
	FriendAccept(pctx context.Context, args user.EditFriendStatusAcceptedReq) error
	UserMockData(ctx context.Context, count int16) error
}

type userUsecase struct {
	store         userdb.Store
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

func (u *userUsecase) GetUserInfoByEmail(pctx context.Context, email string) (userdb.User, error) {
	userData, err := u.store.GetUserByEmail(pctx, email)
	if err != nil {
		return userdb.User{}, err
	}

	return userData, nil
}

func (u *userUsecase) GetUserInfoByUsername(pctx context.Context, username string) (userdb.User, error) {
	userData, err := u.store.GetUserByUsername(pctx, username)
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

func (u *userUsecase) GetBatchUserProfiles(pctx context.Context, ids []int32) ([]userdb.GetBatchUserProfilesRow, error) {
	userData, err := u.store.GetBatchUserProfiles(pctx, ids)
	if err != nil {
		return []userdb.GetBatchUserProfilesRow{}, err
	}
	return userData, nil
}

func (u *userUsecase) GetUserProfileByEmail(pctx context.Context, email string) (user.UserProfile, error) {
	userData, err := u.store.GetUserByEmail(pctx, email)
	if err != nil {
		return user.UserProfile{}, err
	}

	userID := sql.NullInt32{Int32: userData.ID, Valid: true}
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

func (u *userUsecase) GetUserProfileByUsername(pctx context.Context, username string) (user.UserProfile, error) {
	userData, err := u.store.GetUserByUsername(pctx, username)
	if err != nil {
		return user.UserProfile{}, err
	}

	userID := sql.NullInt32{Int32: userData.ID, Valid: true}
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

func (u *userUsecase) FriendInfo(pctx context.Context, args userdb.GetFriendParams) ([]userdb.Friend, error) {
	friend, err := u.store.GetFriend(pctx, args)
	if err != nil {
		return nil, err
	}
	return []userdb.Friend{friend}, nil
}

func (u *userUsecase) FriendListById(pctx context.Context, id int) ([]userdb.ListFriendsByUserIdRow, error) {
	friends, err := u.store.ListFriendsByUserId(pctx, utils.ConvertIntToSqlNullInt32(id))
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (u *userUsecase) FriendsCount(pctx context.Context, userId1 sql.NullInt32) (int64, error) {
	count, err := u.store.GetFriendsCountByID(pctx, userId1)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *userUsecase) FriendsRequestedList(pctx context.Context, userId1 sql.NullInt32) ([]userdb.User, error) {
	friends, err := u.store.GetFriendsRequestedList(pctx, userId1)
	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (u *userUsecase) FriendsPendingList(pctx context.Context, userId2 sql.NullInt32) ([]userdb.User, error) {
	friends, err := u.store.GetFriendsPendingList(pctx, userId2)
	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (u *userUsecase) AddFriend(pctx context.Context, args user.CreateFriendReq) (userdb.Friend, error) {
	userID1 := utils.ConvertIntToSqlNullInt32(args.User_id1)
	userID2 := utils.ConvertIntToSqlNullInt32(args.User_id2)
	arg := userdb.CreateFriendParams{
		UserId1: userID1,
		UserId2: userID2,
	}

	newFriend, err := u.store.CreateFriend(pctx, arg)
	if err != nil {
		return userdb.Friend{}, err
	}

	return newFriend, nil
}

func (u *userUsecase) FriendAccept(pctx context.Context, args user.EditFriendStatusAcceptedReq) error {
	userID1 := utils.ConvertIntToSqlNullInt32(args.User_id1)
	userID2 := utils.ConvertIntToSqlNullInt32(args.User_id2)
	arg := userdb.EditFriendStatusAcceptedParams{
		UserId1: userID1,
		UserId2: userID2,
	}
	err := u.store.EditFriendStatusAccepted(pctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) UserMockData(ctx context.Context, count int16) error {
	createdUsers := make([]userdb.User, 0, count)
	for i := int16(0); i < count; i++ {
		userData, err := u.createUser(ctx, i)
		if err != nil {
			return err
		}
		createdUsers = append(createdUsers, userData)
	}
	for _, user := range createdUsers {
		if err := u.createFakeFriends(ctx, user.ID, int32(count)); err != nil {
			return err
		}
	}

	return nil
}

func (u *userUsecase) createUser(ctx context.Context, seed int16) (userdb.User, error) {
	var userData userdb.CreateUserParams

	fake := faker.NewWithSeed(rand.NewSource(time.Now().UnixNano() + int64(seed)))

	userData.Nationality = "Thai"
	userData.Age = int32(rand.Intn(40-10+1) + 10)
	userData.Gender = fake.Person().Gender()

	phoneNumber := fmt.Sprintf("%010d", rand.Intn(10000000000))
	userData.PhoneNo = phoneNumber
	userData.Email = fake.Internet().Email()

	firstName := fake.Person().FirstName()
	lastName := fake.Person().LastName()
	username := firstName + lastName + strconv.Itoa(int(seed))
	userData.Firstname = firstName
	userData.Lastname = lastName
	userData.Username = username

	fakeBirthdayString := fake.Time().UnixDate(time.Now())
	fakeBirthday, err := time.Parse(time.UnixDate, fakeBirthdayString)
	if err != nil {
		return userdb.User{}, err
	}
	userData.Birthday = fakeBirthday

	fakeImageFile := fake.Image().Image(200, 200)
	filename := fakeImageFile.Name()
	userData.Photourl = sql.NullString{String: filename, Valid: true}

	createdUser, err := u.store.CreateUser(ctx, userData)
	if err != nil {
		return userdb.User{}, err
	}

	return createdUser, nil
}

func (u *userUsecase) createFakeFriends(ctx context.Context, userID int32, totalUsers int32) error {
	existingUserIDs := make([]int32, 0, totalUsers)
	for i := int32(1); i <= totalUsers; i++ {
		if i != userID { // Exclude the current user
			existingUserIDs = append(existingUserIDs, i)
		}
	}
	createdFriendships := make(map[[2]int32]struct{})
	numFriends := rand.Intn(int(totalUsers))
	for i := 0; i < numFriends; i++ {
		friendID := existingUserIDs[rand.Intn(len(existingUserIDs))]

		user1, user2 := userID, friendID
		if userID > friendID {
			user1, user2 = friendID, userID
		}
		friendship := [2]int32{user1, user2}
		if _, exists := createdFriendships[friendship]; exists {
			continue
		}
		_, err := u.store.CreateFriend(ctx, userdb.CreateFriendParams{
			UserId1: utils.ConvertIntToSqlNullInt32(int(user1)),
			UserId2: utils.ConvertIntToSqlNullInt32(int(user2)),
		})
		if err != nil {
			return err
		}

		createdFriendships[friendship] = struct{}{}
	}

	return nil
}

func (u *userUsecase) EditName(ctx context.Context, userID int32, firstName, lastName string) (user.UserProfile, error) {
	var updatedProfile user.UserProfile

	if firstName != "" && lastName == "" {
		err := u.editFirstNameOnly(ctx, userID, firstName)
		if err != nil {
			return updatedProfile, err
		}
	}

	if firstName == "" && lastName != "" {
		err := u.editLastNameOnly(ctx, userID, lastName)
		if err != nil {
			return updatedProfile, err
		}
	}

	if firstName != "" && lastName != "" {
		err := u.editBothNames(ctx, userID, firstName, lastName)
		if err != nil {
			return updatedProfile, err
		}
	}

	updatedProfile, err := u.GetUserProfile(ctx, int(userID))
	if err != nil {
		return updatedProfile, err
	}

	return updatedProfile, nil
}

func (u *userUsecase) editFirstNameOnly(ctx context.Context, userID int32, firstName string) error {
	return u.store.EditFirstNameOnly(ctx, userdb.EditFirstNameOnlyParams{
		ID:        userID,
		Firstname: firstName,
	})
}

func (u *userUsecase) editLastNameOnly(ctx context.Context, userID int32, lastName string) error {
	return u.store.EditLastNameOnly(ctx, userdb.EditLastNameOnlyParams{
		ID:       userID,
		Lastname: lastName,
	})
}

func (u *userUsecase) editBothNames(ctx context.Context, userID int32, firstName, lastName string) error {
	return u.store.EditBothNames(ctx, userdb.EditBothNamesParams{
		ID:        userID,
		Firstname: firstName,
		Lastname:  lastName,
	})
}
