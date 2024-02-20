package group

type (
	NewGroupReq struct {
		GroupName      string `json:"group_name"`
		GroupCreatorId int    `json:"group_creator_id"`
	}
)
