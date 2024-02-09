package user

type (
	UserProfile struct {
		Username     string `json:"username"`
		Firstname    string `json:"firstname"`
		Lastname     string `json:"lastname"`
		FriendsCount int    `json:"friends_count"`
		PhotoUrl     string `json:"photo_url"`
	}

	NewUser struct {
		Username    string `json:"username" form:"username" validate:"required,max=30"`
		Firstname   string `json:"firstname" form:"firstname" validate:"required,max=255"`
		Lastname    string `json:"lastname" form:"lastname" validate:"required,max=255"`
		PhoneNo     string `json:"phone_no" form:"phone_no" validate:"required,max=20"`
		Email       string `json:"email" form:"email" validate:"required,email,max=255"`
		Nationality string `json:"nationality" form:"nationality" validate:"required,max=255"`
		Age         int    `json:"age" form:"age" validate:"required"`
		Birthday    string `json:"birthday" form:"birthday" validate:"required,max=255"`
		Gender      string `json:"gender" form:"gender" validate:"required,max=255"`
	}
)
