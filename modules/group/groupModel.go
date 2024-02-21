package group

type (
	NewGroupReq struct {
		GroupName      string `json:"group_name" form:"group_name" validate:"required,max=255"`
		GroupCreatorId int    `json:"group_creator_id" form:"group_creator_id" validate:"required"`
	}

	NewPostReq struct {
		MemberID    int  `json:"member_id" form:"member_id" validate:"required"`
		GroupID     int  `json:"group_id" form:"group_id" validate:"required"`
		Description string `json:"description" form:"description" validate:"required"`
	}
)
