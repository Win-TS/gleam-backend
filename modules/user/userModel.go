package user

import (
	"database/sql"
	"time"
)

type (
	UserProfile struct {
		Username     string `json:"username"`
		Firstname    string `json:"firstname"`
		Lastname     string `json:"lastname"`
		FriendsCount int    `json:"friends_count"`
		PhotoUrl     string `json:"photo_url"`
	}

	NewUserReq struct {
		Username    string `json:"username" form:"username" validate:"required,max=30"`
		Firstname   string `json:"firstname" form:"firstname" validate:"required,max=255"`
		Lastname    string `json:"lastname" form:"lastname" validate:"required,max=255"`
		PhoneNo     string `json:"phone_no" form:"phone_no" validate:"required,max=20"`
		Email       string `json:"email" form:"email" validate:"required,email,max=255"`
		Nationality string `json:"nationality" form:"nationality" validate:"required,max=255"`
		Birthday    string `json:"birthday" form:"birthday" validate:"required,max=255"`
		Gender      string `json:"gender" form:"gender" validate:"required,max=255"`
		PhotoUrl    string `json:"photo_url" form:"photo_url" validate:"max=255"`
		Password    string `json:"password" form:"password" validate:"required,max=255"`
	}

	NewUserRes struct {
		ID             int32          `json:"id"`
		FirebaseUID    string         `json:"firebase_uid"`
		Username       string         `json:"username"`
		Email          string         `json:"email"`
		Firstname      string         `json:"firstname"`
		Lastname       string         `json:"lastname"`
		PhoneNo        string         `json:"phone_no"`
		PrivateAccount bool           `json:"private_account"`
		Nationality    string         `json:"nationality"`
		Birthday       time.Time      `json:"birthday"`
		Gender         string         `json:"gender"`
		Photourl       sql.NullString `json:"photourl"`
		CreatedAt      time.Time      `json:"created_at"`
	}

	CreateFriendReq struct {
		User_id1 int `json:"user_id1"`
		User_id2 int `json:"user_id2"`
	}

	EditFriendStatusAcceptedReq struct {
		User_id1 int `json:"user_id1"`
		User_id2 int `json:"user_id2"`
	}

	EditUserNameReq struct {
		UserID    int32  `json:"user_id"`
		FirstName string `json:"firstname,omitempty"`
		LastName  string `json:"lastname,omitempty"`
	}
)
