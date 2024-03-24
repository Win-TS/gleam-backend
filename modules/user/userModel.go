package user

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
		Age         int    `json:"age" form:"age" validate:"required"`
		Birthday    string `json:"birthday" form:"birthday" validate:"required,max=255"`
		Gender      string `json:"gender" form:"gender" validate:"required,max=255"`
		//PhotoUrl    string `json:"photo_url" form:"photo_url" validate:"max=255"`
	}

	NewUserRes struct {
		User_id     int    `json:"user_id"`
		Username    string `json:"username"`
		Firstname   string `json:"firstname"`
		Lastname    string `json:"lastname"`
		PhoneNo     string `json:"phone_no"`
		Email       string `json:"email"`
		Nationality string `json:"nationality"`
		Age         int    `json:"age"`
		Birthday    string `json:"birthday"`
		Gender      string `json:"gender"`
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
