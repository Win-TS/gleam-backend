package groupdb

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateNewGroup(ctx context.Context, args CreateGroupParams) (Group, error)
	AcceptGroupRequest(ctx context.Context, args AcceptGroupRequestParams) (GroupMember, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

type AcceptGroupRequestParams struct {
	GroupID    int `json:"group_id" form:"group_id" validate:"required"`
	MemberID   int `json:"member_id" form:"member_id" validate:"required"`
	AcceptorId int `json:"acceptor_id" form:"acceptor_id" validate:"required"`
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

// CreateNewGroup creates a new group with the given name and creator id and add the creator to the group_members table
func (store *SQLStore) CreateNewGroup(ctx context.Context, args CreateGroupParams) (Group, error) {
	var newGroup Group

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		newGroup, err = q.CreateGroup(ctx, args)
		if err != nil {
			return err
		}

		arg := AddGroupMemberParams{
			GroupID:  newGroup.GroupID,
			MemberID: newGroup.GroupCreatorID,
			Role:     "creator",
		}

		_, err = q.AddGroupMember(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return Group{}, err
	}

	return newGroup, nil
}

// AcceptGroupRequest accepts a group request by adding the member to the group and deleting the request
func (store *SQLStore) AcceptGroupRequest(ctx context.Context, args AcceptGroupRequestParams) (GroupMember, error) {
	var newMember GroupMember

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		groupReq, err := q.GetGroupRequest(ctx, GetGroupRequestParams{
			GroupID:  int32(args.GroupID),
			MemberID: int32(args.MemberID),
		})
		if err != nil {
			return err
		}

		newMember, err = q.AddGroupMember(ctx, AddGroupMemberParams{
			GroupID:  groupReq.GroupID,
			MemberID: groupReq.MemberID,
			Role:     "member",
		})
		if err != nil {
			return err
		}

		err = q.DeleteRequestToJoinGroup(ctx, DeleteRequestToJoinGroupParams{
			GroupID:  groupReq.GroupID,
			MemberID: groupReq.MemberID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return GroupMember{}, err
	}

	return newMember, nil
}
