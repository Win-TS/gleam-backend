package auth

type (
	RequestPayload struct {
		Email       string `json:"email" form:"email" validate:"required,email,max=225"`
		PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required,max=32"` // PhoneNumber should be form of E.164 or +66...
		Password    string `json:"password" form:"password" validate:"required,max=32"`
		Username    string `json:"username" form:"username" validate:"required,max=255"`
		UserId      int    `json:"user_id" form:"user_id" validate:"required"`
	}

	UpdateUserReq struct {
		UID         string `json:"uid" form:"uid" validate:"required,max=255"`
		Email       string `json:"email" form:"email" validate:"required,email,max=225"`
		PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required,max=32"` // PhoneNumber should be form of E.164 or +66...
		Password    string `json:"password" form:"password" validate:"required,max=32"`
	}

	UpdatePasswordReq struct {
		UID      string `json:"uid" form:"uid" validate:"required,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
	}
)
