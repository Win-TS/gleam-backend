package auth

type (
	RequestPayload struct {
		Email       string `json:"email" form:"email" validate:"required,email,max=225"`
		PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required,max=32"` // PhoneNumber should be form of E.164 or +66...
		Password    string `json:"password" form:"password" validate:"required,max=32"`
	}

	EmailCheck struct {
		Email string `json:"email" form:"email" validate:"required,email,max=225"`
	}

	PhoneCheck struct {
		PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required,max=32"` // PhoneNumber should be form of E.164 or +66...
	}

	UIDCheck struct {
		UID string `json:"uid" form:"uid" validate:"required,max=255"`
	}
)
