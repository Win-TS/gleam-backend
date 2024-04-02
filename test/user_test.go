package test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/Win-TS/gleam-backend.git/modules/user"
	mocks "github.com/Win-TS/gleam-backend.git/modules/user/userMock"
	userdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/userdb/sqlc"
	"github.com/Win-TS/gleam-backend.git/pkg/utils"
	"github.com/golang/mock/gomock"
)

func TestGetUserProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockObj := mocks.NewMockUserUsecaseService(ctrl)

	userID := 123
	expectedUserProfile := user.UserProfile{
		Username:     "testuser",
		Firstname:    "Test",
		Lastname:     "User",
		FriendsCount: 0,
		PhotoUrl:     "test.jpg",
	}
	mockObj.EXPECT().GetUserProfile(gomock.Any(), userID).Return(expectedUserProfile, nil)

	actualUserProfile, err := mockObj.GetUserProfile(mockObj, userID)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if actualUserProfile != expectedUserProfile {
		t.Errorf("Unexpected user profile returned. Expected: %v, but got: %v", expectedUserProfile, actualUserProfile)
	}
}

func TestGetUserProfile_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockObj := mocks.NewMockUserUsecaseService(ctrl)

	userID := 123
	expectedError := errors.New("database error")
	mockObj.EXPECT().GetUserProfile(gomock.Any(), userID).Return(user.UserProfile{}, expectedError)

	_, err := mockObj.GetUserProfile(mockObj, userID)

	if err == nil {
		t.Errorf("Expected error, but got none")
	}

	if err != expectedError {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetUserProfile_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockObj := mocks.NewMockUserUsecaseService(ctrl)

	userID := 123
	expectedError := errors.New("user not found")
	mockObj.EXPECT().GetUserProfile(gomock.Any(), userID).Return(user.UserProfile{}, expectedError)

	_, err := mockObj.GetUserProfile(mockObj, userID)

	if err == nil {
		t.Errorf("Expected error, but got none")
	}

	if err != expectedError {
		t.Errorf("Unexpected error: %v", err)
	}
}

// func TestGetUserProfile_InvalidUserID_VerifyNotCalled(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	// **No need to create a mock object here**

// 	invalidID := "4" // or any invalid user ID format

// 	_, err := GetUserProfile(invalidID) // Pass the invalidID as needed

// 	if err == nil {
// 		t.Errorf("Expected error for invalid user ID, but got none")
// 	}

// 	// No verification needed since we're not using a mock object
// }

func TestGetLatestIdSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserUsecaseService(ctrl)

	expectedID := 123
	mockService.EXPECT().GetLatestId(gomock.Any()).Return(expectedID, nil)

	id, err := mockService.GetLatestId(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if id != expectedID {
		t.Errorf("unexpected ID returned. Expected: %d, got: %d", expectedID, id)
	}
}

func TestGetLatestId_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserUsecaseService(ctrl)

	expectedError := errors.New("database error")
	mockService.EXPECT().GetLatestId(gomock.Any()).Return(0, expectedError)

	id, err := mockService.GetLatestId(context.Background())

	if err == nil {
		t.Errorf("Expected error, but got none")
		return
	}

	if err != expectedError {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if id != 0 {
		t.Errorf("Unexpected ID returned. Expected: 0, got: %d", id)
	}
}

func TestRegisterNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := mocks.NewMockUserUsecaseService(ctrl)
	birthdayStr := "2005-05-16"
	birthday, err := time.Parse("2006-01-02", birthdayStr)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	payload := &user.NewUserReq{
		Username:    "john_doe",
		Firstname:   "John",
		Lastname:    "Doe",
		PhoneNo:     "123456789",
		Email:       "john@example.com",
		Nationality: "US",
		Birthday:    birthdayStr,
		Gender:      "male",
	}

	photoURL := "test.jpg"
	expectedUser := userdb.User{
		Username:    "john_doe",
		Firstname:   "John",
		Lastname:    "Doe",
		PhoneNo:     "123456789",
		Email:       "john@example.com",
		Nationality: "US",
		Birthday:    birthday,
		Gender:      "male",
	}

	mockService.EXPECT().RegisterNewUser(gomock.Any(), payload, photoURL).Return(expectedUser, nil)

	newUser, err := mockService.RegisterNewUser(context.Background(), payload, photoURL)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if newUser != expectedUser {
		t.Errorf("unexpected user returned. Expected: %v, got: %v", expectedUser, newUser)
	}
}

// func TestRegisterNewUser_InvalidData(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockService := mocks.NewMockUserUsecaseService(ctrl)

// 	// Create a payload with missing required fields
// 	payload := &user.NewUserReq{
// 		// Missing Username, Firstname, Lastname, PhoneNo, Email fields
// 		Nationality: "US",
// 		Age:         30,
// 		Birthday:    "2005-05-16",
// 		Gender:      "male",
// 	}

// 	photoURL := "test.jpg"

// 	// Expect the RegisterNewUser function not to be called
// 	mockService.EXPECT().RegisterNewUser(gomock.Any(), payload, photoURL).Times(0)

// 	// Call RegisterNewUser with the incomplete payload
// 	_, err := mockService.RegisterNewUser(context.Background(), payload, photoURL)

// 	// Verify that an error is returned
// 	if err == nil {
// 		t.Error("Expected error but got nil")
// 	} else {
// 		t.Logf("Expected error: %v", err)
// 	}
// }

func TestSaveToFirebaseStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserUsecaseService(ctrl)

	ctx := context.Background()
	bucketName := "test-bucket"
	objectPath := "test-path"
	filename := "test-file.txt"
	fileContent := "Hello, World!"
	file := strings.NewReader(fileContent)

	mockService.EXPECT().SaveToFirebaseStorage(gomock.Any(), bucketName, objectPath, filename, gomock.Any()).Return("https://example.com", nil)

	url, err := mockService.SaveToFirebaseStorage(ctx, bucketName, objectPath, filename, file)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expectedURL := "https://example.com"
	if url != expectedURL {
		t.Errorf("Unexpected URL returned. Expected: %s, Got: %s", expectedURL, url)
	}
}

func TestEditUsernameSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserUsecaseService(ctrl)

	userID := 123
	args := userdb.ChangeUsernameParams{
		ID:       int32(userID),
		Username: "new_username",
	}

	expectedUserProfile := user.UserProfile{Username: "new_username", Firstname: "John", Lastname: "Doe", FriendsCount: 0, PhotoUrl: "test.jpg"}
	mockService.EXPECT().EditUsername(gomock.Any(), args).Return(expectedUserProfile, nil)

	userProfile, err := mockService.EditUsername(context.Background(), args)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if userProfile != expectedUserProfile {
		t.Errorf("Unexpected user profile returned. Expected: %v, got: %v", expectedUserProfile, userProfile)
	}
}

func TestEditUsername_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserUsecaseService(ctrl)

	// Define test data
	userID := 123
	newUsername := "new_username"
	args := userdb.ChangeUsernameParams{
		ID:       int32(userID),
		Username: newUsername,
	}

	expectedError := errors.New("failed to edit username")

	mockService.EXPECT().EditUsername(gomock.Any(), args).Return(user.UserProfile{}, expectedError)

	userProfile, err := mockService.EditUsername(context.Background(), args)

	if err == nil {
		t.Error("Expected an error, but got nil")
		return
	}

	if userProfile != (user.UserProfile{}) {
		t.Errorf("Unexpected user profile returned. Expected empty profile, but got: %v", userProfile)
	}

	expectedErrorMessage := "failed to edit username"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Unexpected error message. Expected: %s, got: %s", expectedErrorMessage, err.Error())
	}
}

// func TestEditUsername_InvalidUserID(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockService := mocks.NewMockUserUsecaseService(ctrl)

// 	invalidUserID := -1 // Invalid userID
// 	newUsername := "new_username"
// 	args := userdb.ChangeUsernameParams{
// 		ID:       int32(invalidUserID),
// 		Username: newUsername,
// 	}

// 	gomock.InOrder(
// 		mockService.EXPECT().EditUsername(gomock.Any(), gomock.Any()).Times(0),
// 	)

// 	// Call the function under test
// 	userProfile, err := mockService.EditUsername(context.Background(), args)

// 	// Verify the error
// 	if err == nil {
// 		t.Error("Expected an error, but got nil")
// 		return
// 	}


// 	// // We don't expect a user profile to be returned on error
// 	var expectedUserProfile user.UserProfile
// 	if userProfile != expectedUserProfile {
// 		t.Errorf("Unexpected user profile returned. Expected empty user profile on error, but got: %v", userProfile)
// 	}
// }

func TestEditPhoneNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserUsecaseService(ctrl)

	userID := 123
	args := userdb.ChangePhoneNoParams{
		ID:      int32(userID),
		PhoneNo: "123456789",
	}

	expectedUser := userdb.User{ID: 123, Username: "john_doe", Firstname: "John", Lastname: "Doe", PhoneNo: "123456789", Email: "john@example.com", Nationality: "US", Birthday: time.Now(), Gender: "male"}
	mockService.EXPECT().EditPhoneNumber(gomock.Any(), args).Return(expectedUser, nil)

	returnedUser, err := mockService.EditPhoneNumber(context.Background(), args)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if returnedUser != expectedUser {
		t.Errorf("Unexpected user returned. Expected: %v, got: %v", expectedUser, returnedUser)
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserUsecaseService(ctrl)

	userID := 123

	mockService.EXPECT().DeleteUser(gomock.Any(), userID).Return(nil)

	err := mockService.DeleteUser(context.Background(), userID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

}

func TestFriendInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserUsecaseService(ctrl)

	args := userdb.GetFriendParams{
		UserId1: utils.ConvertIntToSqlNullInt32(1),
		UserId2: utils.ConvertIntToSqlNullInt32(2),
	}

	expectedFriend := userdb.Friend{
		ID:        1,
		UserId1:   utils.ConvertIntToSqlNullInt32(1),
		UserId2:   utils.ConvertIntToSqlNullInt32(2),
		Status:    utils.ConvertStringToSqlNullString("accepted"),
		CreatedAt: time.Now(),
	}

	mockService.EXPECT().FriendInfo(gomock.Any(), args).Return([]userdb.Friend{expectedFriend}, nil)

	friends, err := mockService.FriendInfo(context.Background(), args)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if len(friends) != 1 || friends[0] != expectedFriend {
		t.Errorf("Unexpected friends returned. Expected: %v, got: %v", []userdb.Friend{expectedFriend}, friends)
	}
}

func TestFriendsRequestedList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserUsecaseService(ctrl)
	userID := utils.ConvertIntToSqlNullInt32(123)

	expectedFriends := []userdb.User{
		{ID: 1, Username: "user1", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
	}

	mockService.EXPECT().FriendsRequestedList(gomock.Any(), userID).Return(expectedFriends, nil)

	friends, err := mockService.FriendsRequestedList(context.Background(), userID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if len(friends) != len(expectedFriends) {
		t.Errorf("Unexpected number of friends returned. Expected: %d, got: %d", len(expectedFriends), len(friends))
	}
	for i, friend := range friends {
		if friend != expectedFriends[i] {
			t.Errorf("Unexpected friend returned at index %d. Expected: %v, got: %v", i, expectedFriends[i], friend)
		}
	}
}

func TestFriendsPendingList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserUsecaseService(ctrl)
	userID := utils.ConvertIntToSqlNullInt32(123)

	expectedFriends := []userdb.User{
		{ID: 1, Username: "user1", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
	}

	mockService.EXPECT().FriendsPendingList(gomock.Any(), userID).Return(expectedFriends, nil)

	friends, err := mockService.FriendsPendingList(context.Background(), userID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if len(friends) != len(expectedFriends) {
		t.Errorf("Unexpected number of friends returned. Expected: %d, got: %d", len(expectedFriends), len(friends))
	}
	for i, friend := range friends {
		if friend != expectedFriends[i] {
			t.Errorf("Unexpected friend returned at index %d. Expected: %v, got: %v", i, expectedFriends[i], friend)
		}
	}
}

func TestAddFriend(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserUsecaseService(ctrl)

	args := user.CreateFriendReq{
		User_id1: 1,
		User_id2: 2,
	}

	userID1 := utils.ConvertIntToSqlNullInt32(args.User_id1)
	userID2 := utils.ConvertIntToSqlNullInt32(args.User_id2)

	newFriend := userdb.Friend{ID: 7, UserId1: userID1, UserId2: userID2}
	mockService.EXPECT().AddFriend(gomock.Any(), user.CreateFriendReq{User_id1: int(userID1.Int32), User_id2: int(userID2.Int32)}).Return(newFriend, nil)

	friend, err := mockService.AddFriend(context.Background(), args)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expectedFriend := newFriend
	if friend != expectedFriend {
		t.Errorf("Unexpected friend returned. Expected: %v, Got: %v", expectedFriend, friend)
	}
}