package groupdb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Win-TS/gleam-backend.git/pkg/utils"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateGroupWithTags(ctx context.Context, args CreateGroupWithTagsParams) (CreateGroupWithTagsRes, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

type CreateGroupWithTagsParams struct {
	GroupName      string `json:"group_name" form:"group_name" validate:"required,max=255"`
	GroupCreatorId int    `json:"group_creator_id" form:"group_creator_id" validate:"required"`
	PhotoUrl       string `json:"photo_url" form:"photo_url"`
	TagIDs         []int  `json:"tag_ids" form:"tag_ids"`
}

type CreateGroupWithTagsRes struct {
	GroupID        int       `json:"group_id"`
	GroupName      string    `json:"group_name"`
	GroupCreatorId int       `json:"group_creator_id"`
	PhotoUrl       string    `json:"photo_url"`
	TagIDs         []int     `json:"tag_ids"`
	CreatedAt      time.Time `json:"created_at"`
}

// Create new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// CreateGroupWithTags creates a new group with tags
func (store *SQLStore) CreateGroupWithTags(ctx context.Context, args CreateGroupWithTagsParams) (CreateGroupWithTagsRes, error) {
	var groupInfo Group

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		groupInfo, err = q.CreateGroup(ctx, CreateGroupParams{
			GroupName:      args.GroupName,
			GroupCreatorID: int32(args.GroupCreatorId),
			PhotoUrl:       utils.ConvertStringToSqlNullString(args.PhotoUrl),
		})
		if err != nil {
			return err
		}

		arg := AddGroupMemberParams{
			GroupID:  int32(groupInfo.GroupID),
			MemberID: int32(groupInfo.GroupCreatorID),
			Role:     "creator",
		}
		_, err = store.AddGroupMember(ctx, arg)
		if err != nil {
			return err
		}

		for _, tagID := range args.TagIDs {
			_, err := q.AddGroupTag(ctx, AddGroupTagParams{
				GroupID: groupInfo.GroupID,
				TagID:   int32(tagID),
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	return CreateGroupWithTagsRes{
		GroupID:        int(groupInfo.GroupID),
		GroupName:      groupInfo.GroupName,
		GroupCreatorId: int(groupInfo.GroupCreatorID),
		PhotoUrl:       groupInfo.PhotoUrl.String,
		TagIDs:         args.TagIDs,
		CreatedAt:      groupInfo.CreatedAt,
	}, err
}