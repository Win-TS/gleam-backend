package group

import (
	"database/sql"
	"time"

	groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
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
		TagName    string `json:"tag_name" form:"tag_name" validate:"required,max=255"`
		IconUrl    string `json:"icon_url" form:"icon_url"`
		CategoryId int    `json:"category_id" form:"category_id" validate:"required"`
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

	GroupRequestRes struct {
		GroupID      int32          `json:"group_id"`
		MemberID     int32          `json:"member_id"`
		Description  sql.NullString `json:"description"`
		CreatedAt    time.Time      `json:"created_at"`
		UserID       int32          `json:"user_id"`
		Username     string         `json:"username"`
		UserPhotourl string         `json:"user_photourl"`
	}

	GroupMemberRes struct {
		GroupID      int32     `json:"group_id"`
		MemberID     int32     `json:"member_id"`
		Role         string    `json:"role"`
		CreatedAt    time.Time `json:"created_at"`
		UserID       int32     `json:"user_id"`
		Username     string    `json:"username"`
		UserPhotourl string    `json:"user_photourl"`
	}
	PostByGroupRes struct {
		PostID       int32          `json:"post_id"`
		MemberID     int32          `json:"member_id"`
		GroupID      int32          `json:"group_id"`
		PhotoUrl     sql.NullString `json:"photo_url"`
		Description  sql.NullString `json:"description"`
		CreatedAt    time.Time      `json:"created_at"`
		UserID       int32          `json:"user_id"`
		Username     string         `json:"username"`
		UserPhotourl string         `json:"user_photourl"`
	}
	ReactionPostRes struct {
		ReactionID   int32     `json:"reaction_id"`
		PostID       int32     `json:"post_id"`
		MemberID     int32     `json:"member_id"`
		Reaction     string    `json:"reaction"`
		CreatedAt    time.Time `json:"created_at"`
		UserID       int32     `json:"user_id"`
		Username     string    `json:"username"`
		UserPhotourl string    `json:"user_photourl"`
	}

	CommentRes struct {
		CommentID    int32     `json:"comment_id"`
		PostID       int32     `json:"post_id"`
		MemberID     int32     `json:"member_id"`
		Comment      string    `json:"comment"`
		CreatedAt    time.Time `json:"created_at"`
		UserID       int32     `json:"user_id"`
		Username     string    `json:"username"`
		UserPhotourl string    `json:"user_photourl"`
	}

	PostsForFeedRes struct {
		PostID         int32          `json:"post_id"`
		MemberID       int32          `json:"member_id"`
		GroupID        int32          `json:"group_id"`
		PhotoUrl       sql.NullString `json:"photo_url"`
		Description    sql.NullString `json:"description"`
		CreatedAt      time.Time      `json:"created_at"`
		GroupName      string         `json:"group_name"`
		GroupPhotoUrl  sql.NullString `json:"group_photo_url"`
		PosterUsername string         `json:"poster_username"`
		PosterPhotoUrl string         `json:"poster_photo_url"`
	}

	GroupInfoWithMemberRes struct {
		GroupID        int32          `json:"group_id"`
		GroupName      string         `json:"group_name"`
		PhotoUrl       sql.NullString `json:"photo_url"`
		GroupCreatorID int32          `json:"group_creator_id"`
		Frequency      sql.NullInt32  `json:"frequency"`
		MaxMembers     int32          `json:"max_members"`
		GroupType      string         `json:"group_type"`
		Visibility     bool           `json:"visibility"`
		Description    sql.NullString `json:"description"`
		CreatedAt      time.Time      `json:"created_at"`
		TagName        string         `json:"tag_name"`
		TotalMember    int32          `json:"total_member"`
	}

	GetGroupByIdRes struct {
		GroupInfo groupdb.GetGroupByIDRow `json:"group_info"`
		UserId    int32                   `json:"user_id"`
		Status    string                  `json:"status"`
	}

	GetUserGroupRes struct {
		SocialGroups   []groupdb.GetUserGroupsRow `json:"social_groups"`
		PersonalGroups []groupdb.GetUserGroupsRow `json:"personal_groups"`
	}
)
