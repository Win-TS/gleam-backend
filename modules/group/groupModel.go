package group

import (
	"database/sql"
	"time"
	//groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
)

type (
	NewGroupReq struct {
		GroupName      string `json:"group_name" form:"group_name" validate:"required,max=255"`
		GroupCreatorId int    `json:"group_creator_id" form:"group_creator_id" validate:"required"`
		TagID          int    `json:"tag_id" form:"tag_id" validate:"required"`
		Description    string `json:"description" form:"description"`
		Frequency      int    `json:"frequency" form:"frequency" validate:"required"`
		MaxMembers     int    `json:"max_members" form:"max_members"`
		GroupType      string `json:"group_type" form:"group_type"`
		Visibility     bool   `json:"visibility" form:"visibility"`
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

	GroupWithTagsRes struct {
		GroupID        int            `json:"group_id"`
		GroupName      string         `json:"group_name"`
		GroupCreatorID int            `json:"group_creator_id"`
		PhotoUrl       sql.NullString `json:"photo_url"`
		Tags           string         `json:"tags"`
		CreatedAt      time.Time      `json:"created_at"`
	}

	RequestToJoinGroupReq struct {
		GroupID     int    `json:"group_id" form:"group_id" validate:"required"`
		MemberID    int    `json:"member_id" form:"member_id" validate:"required"`
		Description string `json:"description" form:"description"`
	}
)
