// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package userdb

import (
	"context"
	"database/sql"
)

type Querier interface {
	ChangePhoneNo(ctx context.Context, arg ChangePhoneNoParams) error
	ChangeUsername(ctx context.Context, arg ChangeUsernameParams) error
	CreateFriend(ctx context.Context, arg CreateFriendParams) (Friend, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, id int32) error
	EditBothNames(ctx context.Context, arg EditBothNamesParams) error
	EditFirstNameOnly(ctx context.Context, arg EditFirstNameOnlyParams) error
	EditFriendStatusAccepted(ctx context.Context, arg EditFriendStatusAcceptedParams) error
	EditFriendStatusDeclined(ctx context.Context, arg EditFriendStatusDeclinedParams) error
	EditLastNameOnly(ctx context.Context, arg EditLastNameOnlyParams) error
	EditUserProfilePicture(ctx context.Context, arg EditUserProfilePictureParams) error
	GetBatchUserProfiles(ctx context.Context, dollar_1 []int32) ([]GetBatchUserProfilesRow, error)
	GetFriend(ctx context.Context, arg GetFriendParams) (Friend, error)
	GetFriendsCountByID(ctx context.Context, userId1 sql.NullInt32) (int64, error)
	GetFriendsPendingList(ctx context.Context, userId2 sql.NullInt32) ([]User, error)
	GetFriendsRequestedList(ctx context.Context, userId1 sql.NullInt32) ([]User, error)
	GetLatestId(ctx context.Context) (int32, error)
	GetUser(ctx context.Context, id int32) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserForUpdate(ctx context.Context, id int32) (User, error)
	ListFriendsByUserId(ctx context.Context, userId1 sql.NullInt32) ([]ListFriendsByUserIdRow, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) error
}

var _ Querier = (*Queries)(nil)
