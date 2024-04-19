package userUsecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"

	"math/rand"
	"time"

	"strconv"

	"firebase.google.com/go/storage"
	"github.com/Win-TS/gleam-backend.git/modules/user"

	authPb "github.com/Win-TS/gleam-backend.git/modules/auth/authPb"
	groupPb "github.com/Win-TS/gleam-backend.git/modules/group/groupPb"
	userdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/userdb/sqlc"
	"github.com/Win-TS/gleam-backend.git/pkg/grpcconn"

	"github.com/Win-TS/gleam-backend.git/pkg/utils"
	"github.com/jaswdr/faker"
)

type UserUsecaseService interface {
	GetUserProfile(pctx context.Context, id int, grpcUrl string) (user.UserProfile, error)
	GetBatchUserProfiles(pctx context.Context, ids []int32) ([]userdb.GetBatchUserProfilesRow, error)
	GetUserProfileByEmail(pctx context.Context, email string) (user.UserProfile, error)
	GetUserProfileByUsername(pctx context.Context, username string) (user.UserProfile, error)
	GetLatestId(pctx context.Context) (int, error)
	RegisterNewUser(pctx context.Context, payload *user.NewUserReq, grpcUrl, photoUrl string) (*user.NewUserRes, error)
	SaveToFirebaseStorage(pctx context.Context, bucketName, objectPath, filename string, file io.Reader) (string, error)
	GetUserInfo(pctx context.Context, id int) (userdb.User, error)
	GetUserInfoByEmail(pctx context.Context, email string) (userdb.User, error)
	GetUserInfoByUsername(pctx context.Context, username string) (userdb.User, error)
	EditUsername(pctx context.Context, args userdb.ChangeUsernameParams, grpcUrl string) (user.UserProfile, error)
	EditPhoneNumber(pctx context.Context, args userdb.ChangePhoneNoParams, grpcUrl string) (user.UserProfile, error)
	EditName(ctx context.Context, userID int32, firstName, lastName, grpcUrl string) (user.UserProfile, error)
	DeleteUser(pctx context.Context, id int, authGrpcUrl, groupGrpcUrl string) error
	FriendInfo(ctx context.Context, args userdb.GetFriendParams) ([]userdb.Friend, error)
	FriendListById(pctx context.Context, args userdb.ListFriendsByUserIdParams) ([]userdb.ListFriendsByUserIdRow, error)
	FriendsCount(pctx context.Context, userId1 sql.NullInt32) (int64, error)
	FriendsRequestedList(pctx context.Context, args userdb.GetFriendsRequestedListParams) ([]userdb.User, error)
	FriendsPendingList(pctx context.Context, args userdb.GetFriendsPendingListParams) ([]userdb.User, error)
	AddFriend(pctx context.Context, args user.CreateFriendReq) (userdb.Friend, error)
	FriendAccept(pctx context.Context, args user.EditFriendStatusAcceptedReq) error
	FriendDecline(pctx context.Context, args user.EditFriendStatusDeclinedReq) error
	UserMockData(ctx context.Context, count int16) error
	EditUserPhoto(pctx context.Context, args userdb.EditUserProfilePictureParams, grpcUrl string) (user.UserProfile, error)
	SearchUsersByUsername(ctx context.Context, args userdb.SearchUsersByUsernameParams) ([]userdb.SearchUsersByUsernameRow, error)
	EditPrivateAccount(ctx context.Context, args userdb.EditPrivateAccountParams) (userdb.User, error)
	FriendListByIdNoPaginate(ctx context.Context, userId int) ([]userdb.ListFriendsByUserIdNoPaginateRow, error)
	GetFriendRequestCount(ctx context.Context, userId int) (int, error)
	MockupUser(ctx context.Context) error
	MockupFriend(ctx context.Context) error
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

func (u *userUsecase) GetUserProfile(pctx context.Context, id int, grpcUrl string) (user.UserProfile, error) {
	userData, err := u.store.GetUser(pctx, int32(id))
	if err != nil {
		return user.UserProfile{}, err
	}

	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("error - gRPC connection failed: %s", err.Error())
	}

	highestStreak, err := conn.Group().UserHighestStreak(pctx, &groupPb.UserHighestStreakReq{UserId: int32(id)})
	if err != nil {
		log.Printf("error - UserHighestStreak failed: %s", err.Error())
		return user.UserProfile{}, errors.New("error: UserHighestStreak failed")
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
		MaxStreak:    int(highestStreak.HighestStreak),
	}, nil
}

func (u *userUsecase) GetBatchUserProfiles(pctx context.Context, ids []int32) ([]userdb.GetBatchUserProfilesRow, error) {
	userData := make([]userdb.GetBatchUserProfilesRow, 0, len(ids))
	for _, id := range ids {
		user, err := u.GetUserInfo(pctx, int(id))
		if err != nil {
			return []userdb.GetBatchUserProfilesRow{}, err
		}
		userData = append(userData, userdb.GetBatchUserProfilesRow{
			ID:        int32(id),
			Username:  user.Username,
			Email:     user.Email,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Photourl:  utils.ConvertStringToSqlNullString(user.Photourl.String),
		})
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

func (u *userUsecase) RegisterNewUser(pctx context.Context, payload *user.NewUserReq, grpcUrl, photoUrl string) (*user.NewUserRes, error) {

	birthdayTime, err := time.Parse("2006-01-02", payload.Birthday)
	if err != nil {
		return &user.NewUserRes{}, err
	}

	sqlPhotoUrl := utils.ConvertStringToSqlNullString(photoUrl)

	newUser, err := u.store.CreateUser(pctx, userdb.CreateUserParams{
		Username:    payload.Username,
		Firstname:   payload.Firstname,
		Lastname:    payload.Lastname,
		PhoneNo:     payload.PhoneNo,
		Email:       payload.Email,
		Nationality: payload.Nationality,
		Birthday:    birthdayTime,
		Gender:      payload.Gender,
		Photourl:    sqlPhotoUrl,
	})
	if err != nil {
		return &user.NewUserRes{}, err
	}

	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		u.DeleteUser(pctx, int(newUser.ID), grpcUrl, "")
		log.Printf("error - gRPC connection failed: %s", err.Error())
		return &user.NewUserRes{}, errors.New("error: gRPC connection failed")
	}

	result, err := conn.Auth().RegisterNewUser(pctx, &authPb.RegisterNewUserReq{
		Email:    payload.Email,
		PhoneNo:  "+" + (payload.PhoneNo),
		Password: payload.Password,
		Username: payload.Username,
		UserId:   int32(newUser.ID),
	})
	if err != nil {
		u.DeleteUser(pctx, int(newUser.ID), grpcUrl, "")
		log.Printf("error - RegisterNewUser failed: %s", err.Error())
		return &user.NewUserRes{}, errors.New("error: RegisterNewUser failed")
	}

	if !result.Success {
		u.DeleteUser(pctx, int(newUser.ID), grpcUrl, "")
		return &user.NewUserRes{}, errors.New("error: RegisterNewUser failed")
	}

	return &user.NewUserRes{
		ID:             newUser.ID,
		FirebaseUID:    result.Uid,
		Username:       newUser.Username,
		Email:          newUser.Email,
		Firstname:      newUser.Firstname,
		Lastname:       newUser.Lastname,
		PhoneNo:        newUser.PhoneNo,
		PrivateAccount: newUser.PrivateAccount,
		Nationality:    newUser.Nationality,
		Birthday:       newUser.Birthday,
		Gender:         newUser.Gender,
		Photourl:       newUser.Photourl,
		CreatedAt:      newUser.CreatedAt,
	}, nil
}

func (u *userUsecase) EditUserPhoto(pctx context.Context, args userdb.EditUserProfilePictureParams, grpcUrl string) (user.UserProfile, error) {
	if err := u.store.EditUserProfilePicture(pctx, args); err != nil {
		return user.UserProfile{}, err
	}
	return u.GetUserProfile(pctx, int(args.ID), grpcUrl)
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

func (u *userUsecase) EditUsername(pctx context.Context, args userdb.ChangeUsernameParams, grpcUrl string) (user.UserProfile, error) {
	if err := u.store.ChangeUsername(pctx, args); err != nil {
		return user.UserProfile{}, err
	}
	return u.GetUserProfile(pctx, int(args.ID), grpcUrl)
}

func (u *userUsecase) EditPhoneNumber(pctx context.Context, args userdb.ChangePhoneNoParams, grpcUrl string) (user.UserProfile, error) {
	if err := u.store.ChangePhoneNo(pctx, args); err != nil {
		return user.UserProfile{}, err
	}
	return u.GetUserProfile(pctx, int(args.ID), grpcUrl)
}

func (u *userUsecase) DeleteUser(pctx context.Context, id int, authGrpcUrl, groupGrpcUrl string) error {
	if err := u.store.DeleteUser(pctx, int32(id)); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	// user, err := u.GetUserInfo(pctx, id)
	// if err != nil {
	// 	return err
	// }

	groupConn, err := grpcconn.NewGrpcClient(groupGrpcUrl)
	if err != nil {
		log.Printf("error - gRPC connection failed: %s", err.Error())
		return errors.New("error: gRPC connection failed")
	}

	_, err = groupConn.Group().DeleteUserData(ctx, &groupPb.DeleteUserDataReq{UserId: int32(id)})
	if err != nil {
		log.Printf("error - DeleteUserData failed: %s", err.Error())
		return errors.New("error: DeleteUserData failed")
	}

	// authConn, err := grpcconn.NewGrpcClient(authGrpcUrl)
	// if err != nil {
	// 	log.Printf("error - gRPC connection failed: %s", err.Error())
	// 	return errors.New("error: gRPC connection failed")
	// }

	// uidRes, err := authConn.Auth().GetUidFromEmail(pctx, &authPb.GetUidFromEmailReq{Email: user.Email})
	// if err != nil {
	// 	log.Printf("error - DeleteUser (GetUidFromEmail) failed: %s", err.Error())
	// 	return errors.New("error: DeleteUser failed")
	// }

	// result, err := authConn.Auth().DeleteUser(pctx, &authPb.DeleteUserReq{Uid: uidRes.Uid})
	// if err != nil {
	// 	log.Printf("error - DeleteUser (DeleteUser) failed: %s", err.Error())
	// 	return errors.New("error: DeleteUser failed")
	// }

	// if !result.Success {
	// 	return errors.New("error: DeleteUser failed")
	// }

	return nil
}

func (u *userUsecase) FriendInfo(pctx context.Context, args userdb.GetFriendParams) ([]userdb.Friend, error) {
	friend, err := u.store.GetFriend(pctx, args)
	if err != nil {
		return nil, err
	}
	return []userdb.Friend{friend}, nil
}

func (u *userUsecase) FriendListById(pctx context.Context, args userdb.ListFriendsByUserIdParams) ([]userdb.ListFriendsByUserIdRow, error) {
	friends, err := u.store.ListFriendsByUserId(pctx, args)
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

func (u *userUsecase) FriendsRequestedList(pctx context.Context, args userdb.GetFriendsRequestedListParams) ([]userdb.User, error) {
	friends, err := u.store.GetFriendsRequestedList(pctx, args)
	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (u *userUsecase) FriendsPendingList(pctx context.Context, args userdb.GetFriendsPendingListParams) ([]userdb.User, error) {
	friends, err := u.store.GetFriendsPendingList(pctx, args)
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

func (u *userUsecase) FriendDecline(pctx context.Context, args user.EditFriendStatusDeclinedReq) error {
	userID1 := utils.ConvertIntToSqlNullInt32(args.User_id1)
	userID2 := utils.ConvertIntToSqlNullInt32(args.User_id2)
	arg := userdb.EditFriendStatusDeclinedParams{
		UserId1: userID1,
		UserId2: userID2,
	}
	err := u.store.EditFriendStatusDeclined(pctx, arg)
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

	filename := "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2F18.webp?alt=media"
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

func (u *userUsecase) EditName(ctx context.Context, userID int32, firstName, lastName, grpcUrl string) (user.UserProfile, error) {
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

	updatedProfile, err := u.GetUserProfile(ctx, int(userID), grpcUrl)
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

func (u *userUsecase) SearchUsersByUsername(ctx context.Context, args userdb.SearchUsersByUsernameParams) ([]userdb.SearchUsersByUsernameRow, error) {
	return u.store.SearchUsersByUsername(ctx, args)
}

func (u *userUsecase) EditPrivateAccount(ctx context.Context, args userdb.EditPrivateAccountParams) (userdb.User, error) {
	if err := u.store.EditPrivateAccount(ctx, args); err != nil {
		return userdb.User{}, err
	}

	return u.GetUserInfo(ctx, int(args.ID))
}

func (u *userUsecase) FriendListByIdNoPaginate(ctx context.Context, userId int) ([]userdb.ListFriendsByUserIdNoPaginateRow, error) {
	return u.store.ListFriendsByUserIdNoPaginate(ctx, utils.ConvertIntToSqlNullInt32(userId))
}

func (u *userUsecase) GetFriendRequestCount(ctx context.Context, userId int) (int, error) {
	count, err := u.store.GetFriendRequestCount(ctx, utils.ConvertIntToSqlNullInt32(userId))
	if err != nil {
		return -1, err
	}
	return int(count), nil
}

func (u *userUsecase) MockupUser(ctx context.Context) error {

	userDetails := []map[string]interface{}{
		{
			"username":    "Bankie888",
			"firstname":   "Sethanan",
			"lastname":    "BankBank",
			"phone_no":    "+123456123",
			"email":       "Bank@gmail.com",
			"nationality": "TH",
			"birthday":    time.Date(2003, time.January, 2, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fbankie.jpeg?alt=media&token=857d7f0b-2858-4666-9388-177337a81502",
		},
		{
			"username":    "Betty552",
			"firstname":   "Elizabeth",
			"lastname":    "Bethbeth",
			"phone_no":    "+9876541231",
			"email":       "Bethh@gmail.com",
			"nationality": "UK",
			"birthday":    time.Date(2001, time.May, 15, 0, 0, 0, 0, time.UTC),
			"gender":      "female",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fbet%26lyly.jpeg?alt=media&token=cc8e2634-159c-4a38-8912-aa8d128c41ab",
		},
		{
			"username":    "dunepw",
			"firstname":   "Pitiphon",
			"lastname":    "Chaicharoen",
			"phone_no":    "+98765123121",
			"email":       "Dune@gmail.com",
			"nationality": "TH",
			"birthday":    time.Date(2005, time.May, 15, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fdunepw.jpeg?alt=media&token=cc34ea0f-7a3d-4c93-ade4-031c5e54beb6",
		},
		{
			"username":    "jajajedi",
			"firstname":   "Theerothai",
			"lastname":    "Sithlord",
			"phone_no":    "+987123121",
			"email":       "Jedi@gmail.com",
			"nationality": "JP",
			"birthday":    time.Date(2001, time.October, 15, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fjajajedi.jpeg?alt=media&token=80889185-4327-4a0b-b0e7-7d064c5e911e",
		},
		{
			"username":    "kaoskywalker",
			"firstname":   "Thanthai",
			"lastname":    "Kruthong",
			"phone_no":    "+12312321",
			"email":       "999999999@gmail.com",
			"nationality": "US",
			"birthday":    time.Date(2007, time.February, 15, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fkaoskywalkerz.jpeg?alt=media&token=938f051a-71ea-460c-887b-550e22d74842",
		},
		{
			"username":    "Kri7x",
			"firstname":   "Krit",
			"lastname":    "KritKirtKirt",
			"phone_no":    "+1231231241",
			"email":       "krit@gmail.com",
			"nationality": "TH",
			"birthday":    time.Date(2002, time.March, 15, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fkri7x.jpeg?alt=media&token=346e539d-76f0-42ce-ae5f-5f68b7980e72",
		},
		{
			"username":    "GuYzaza888",
			"firstname":   "Krittin",
			"lastname":    "Guyguy",
			"phone_no":    "+423423421",
			"email":       "Guy@gmail.com",
			"nationality": "JP",
			"birthday":    time.Date(2003, time.May, 15, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fkrittineke.jpeg?alt=media&token=a8ab5e13-49fd-4dd1-af5a-e6ed028870f2",
		},
		{
			"username":    "mearzwong999",
			"firstname":   "Wongsapat",
			"lastname":    "Wong",
			"phone_no":    "+9821231",
			"email":       "mearzwong@gmail.com",
			"nationality": "TH",
			"birthday":    time.Date(2001, time.May, 15, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fmearzwong.jpeg?alt=media&token=5d832972-66a5-4765-8f73-8287e7c09309",
		},
		{
			"username":    "Minniecyp888",
			"firstname":   "Chayapa",
			"lastname":    "Minniemouse",
			"phone_no":    "+23421341",
			"email":       "mnmn@gmail.com",
			"nationality": "FN",
			"birthday":    time.Date(2005, time.May, 15, 0, 0, 0, 0, time.UTC),
			"gender":      "female",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fminnie.jpeg?alt=media&token=7a405ab1-a9c4-4007-b90b-d3c4954ab201",
		},
		{
			"username":    "oatptchy",
			"firstname":   "Nadech",
			"lastname":    "Koogimiya",
			"phone_no":    "+2342351",
			"email":       "oat555@gmail.com",
			"nationality": "TH",
			"birthday":    time.Date(2004, time.October, 15, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Foatptchy.jpeg?alt=media&token=4caccd1e-e964-44b7-8834-aedbdfe5130b",
		},
		{
			"username":    "pungdevil66",
			"firstname":   "Tassanai",
			"lastname":    "Peesarj",
			"phone_no":    "+987654234234221",
			"email":       "kanompang@gmail.com",
			"nationality": "TH",
			"birthday":    time.Date(2001, time.May, 15, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fpungmonster.jpeg?alt=media&token=7e1ca9e8-8b14-421c-92cd-d33faa13363c",
		},
		{
			"username":    "rushaoosh",
			"firstname":   "Napat",
			"lastname":    "Laokai",
			"phone_no":    "+23423123151",
			"email":       "menjoo@gmail.com",
			"nationality": "TH",
			"birthday":    time.Date(2003, time.May, 15, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://example.com/janesmith.jpg",
		},
		{
			"username":    "teenoisukiThailand",
			"firstname":   "Tee",
			"lastname":    "Teenoi",
			"phone_no":    "+9002304321",
			"email":       "tee@gmail.com",
			"nationality": "TH",
			"birthday":    time.Date(2005, time.May, 11, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fteenoisuki.jpeg?alt=media&token=e55abfff-74f9-484c-9a1c-b281bc4d2b89",
		},
		{
			"username":    "wints",
			"firstname":   "Win",
			"lastname":    "Joetoo",
			"phone_no":    "+9894321",
			"email":       "yaitoe@gmail.com",
			"nationality": "TH",
			"birthday":    time.Date(2004, time.May, 11, 0, 0, 0, 0, time.UTC),
			"gender":      "male",
			"photourl":    "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fwin_ts.jpeg?alt=media&token=035ef9f4-9b8d-4cd9-b05f-cff1f895445b",
		},
	}

	for _, userDetails := range userDetails {
		_, err := u.store.CreateUser(ctx, userdb.CreateUserParams{
			Username:    userDetails["username"].(string),
			Firstname:   userDetails["firstname"].(string),
			Lastname:    userDetails["lastname"].(string),
			PhoneNo:     userDetails["phone_no"].(string),
			Email:       userDetails["email"].(string),
			Nationality: userDetails["nationality"].(string),
			Birthday:    userDetails["birthday"].(time.Time),
			Gender:      userDetails["gender"].(string),
			Photourl:    sql.NullString{String: userDetails["photourl"].(string), Valid: true},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *userUsecase) MockupFriend(ctx context.Context) error {
	userIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	numFriendships := 20

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numFriendships; i++ {
		userID1 := userIDs[rand.Intn(len(userIDs))]
		userID2 := userIDs[rand.Intn(len(userIDs))]

		for userID1 == userID2 {
			userID2 = userIDs[rand.Intn(len(userIDs))]
		}

		// Check if the friendship already exists
		_, err := u.store.GetFriend(ctx, userdb.GetFriendParams{
			UserId1: utils.ConvertIntToSqlNullInt32(userID1),
			UserId2: utils.ConvertIntToSqlNullInt32(userID2),
		})
		if err == nil {
			continue
		}

		// Create the friendship
		_, err = u.store.CreateFriend(ctx, userdb.CreateFriendParams{
			UserId1: utils.ConvertIntToSqlNullInt32(userID1),
			UserId2: utils.ConvertIntToSqlNullInt32(userID2),
		})
		if err != nil {
			return err
		}

		if rand.Intn(2) == 0 {
			err = u.store.EditFriendStatusAccepted(ctx, userdb.EditFriendStatusAcceptedParams{
				UserId1: utils.ConvertIntToSqlNullInt32(userID1),
				UserId2: utils.ConvertIntToSqlNullInt32(userID2),
			})
			if err != nil {
				return err
			}
		}

		time.Sleep(time.Millisecond * 100)
	}

	return nil
}
