package group

import (
	"database/sql"
	"time"

	groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
)

type (
	NewGroupReq struct {
		GroupName      string `json:"group_name" form:"group_name" validate:"required,max=255"`
		GroupCreatorId int    `json:"group_creator_id" form:"group_creator_id" validate:"required"`
	}

	NewPostReq struct {
		MemberID    int    `json:"member_id" form:"member_id" validate:"required"`
		GroupID     int    `json:"group_id" form:"group_id" validate:"required"`
		Description string `json:"description" form:"description" validate:"required"`
	}

	EditPostReq struct {
		PostID      int    `json:"post_id" form:"post_id" validate:"required"`
		Description string `json:"description" form:"description" validate:"required"`
	}

	NewTagReq struct {
		TagName string `json:"tag_name" form:"tag_name" validate:"required,max=255"`
		IconUrl string `json:"icon_url" form:"icon_url"`
	}

	AddMultipleTagsReq struct {
		GroupID int32   `json:"group_id" form:"group_id" validate:"required"`
		TagIDs  []int32 `json:"tag_ids" form:"tag_ids" validate:"required"`
	}

	GroupWithTagsRes struct {
		GroupID        int            `json:"group_id"`
		GroupName      string         `json:"group_name"`
		GroupCreatorID int            `json:"group_creator_id"`
		PhotoUrl       sql.NullString `json:"photo_url"`
		Tags           []groupdb.Tag  `json:"tags"`
		CreatedAt      time.Time   `json:"created_at"`
	}
)
