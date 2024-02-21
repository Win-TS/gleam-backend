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
	EditFriendStatusAccepted(ctx context.Context, arg EditFriendStatusAcceptedParams) error
	EditFriendStatusDeclined(ctx context.Context, arg EditFriendStatusDeclinedParams) error
	GetFriend(ctx context.Context, arg GetFriendParams) (Friend, error)
	GetFriendForUpdate(ctx context.Context, arg GetFriendForUpdateParams) (Friend, error)
	GetFriendsCountByID(ctx context.Context, userId1 sql.NullInt32) (int64, error)
	GetFriendsListByID(ctx context.Context, userId1 sql.NullInt32) ([]interface{}, error)
	GetFriendsPendingList(ctx context.Context, userId2 sql.NullInt32) ([]Friend, error)
	GetLatestId(ctx context.Context) (int32, error)
	GetUser(ctx context.Context, id int32) (User, error)
	GetUserForUpdate(ctx context.Context, id int32) (User, error)
	ListFriendsByUserId(ctx context.Context, arg ListFriendsByUserIdParams) ([]interface{}, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) error
}

var _ Querier = (*Queries)(nil)
