// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: group.sql

package groupdb

import (
	"context"
	"database/sql"
	"time"
)

const addGroupMember = `-- name: AddGroupMember :one
INSERT INTO group_members (group_id, member_id, role)
VALUES ($1, $2, $3)
RETURNING group_id, member_id, role, created_at
`

type AddGroupMemberParams struct {
	GroupID  int32  `json:"group_id"`
	MemberID int32  `json:"member_id"`
	Role     string `json:"role"`
}

func (q *Queries) AddGroupMember(ctx context.Context, arg AddGroupMemberParams) (GroupMember, error) {
	row := q.db.QueryRowContext(ctx, addGroupMember, arg.GroupID, arg.MemberID, arg.Role)
	var i GroupMember
	err := row.Scan(
		&i.GroupID,
		&i.MemberID,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const checkMemberInGroup = `-- name: CheckMemberInGroup :one
SELECT group_id, member_id, role, created_at FROM group_members
WHERE group_id = $1 AND member_id = $2
`

type CheckMemberInGroupParams struct {
	GroupID  int32 `json:"group_id"`
	MemberID int32 `json:"member_id"`
}

func (q *Queries) CheckMemberInGroup(ctx context.Context, arg CheckMemberInGroupParams) (GroupMember, error) {
	row := q.db.QueryRowContext(ctx, checkMemberInGroup, arg.GroupID, arg.MemberID)
	var i GroupMember
	err := row.Scan(
		&i.GroupID,
		&i.MemberID,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const createGroup = `-- name: CreateGroup :one
INSERT INTO groups (
        group_name,
        group_creator_id,
        photo_url,
        frequency,
        max_members,
        group_type,
        description,
        visibility,
        tag_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING group_id, group_name, group_creator_id, description, photo_url, tag_id, frequency, max_members, group_type, visibility, created_at
`

type CreateGroupParams struct {
	GroupName      string         `json:"group_name"`
	GroupCreatorID int32          `json:"group_creator_id"`
	PhotoUrl       sql.NullString `json:"photo_url"`
	Frequency      sql.NullInt32  `json:"frequency"`
	MaxMembers     int32          `json:"max_members"`
	GroupType      string         `json:"group_type"`
	Description    sql.NullString `json:"description"`
	Visibility     bool           `json:"visibility"`
	TagID          int32          `json:"tag_id"`
}

func (q *Queries) CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error) {
	row := q.db.QueryRowContext(ctx, createGroup,
		arg.GroupName,
		arg.GroupCreatorID,
		arg.PhotoUrl,
		arg.Frequency,
		arg.MaxMembers,
		arg.GroupType,
		arg.Description,
		arg.Visibility,
		arg.TagID,
	)
	var i Group
	err := row.Scan(
		&i.GroupID,
		&i.GroupName,
		&i.GroupCreatorID,
		&i.Description,
		&i.PhotoUrl,
		&i.TagID,
		&i.Frequency,
		&i.MaxMembers,
		&i.GroupType,
		&i.Visibility,
		&i.CreatedAt,
	)
	return i, err
}

const createNewTag = `-- name: CreateNewTag :one
INSERT INTO tags (tag_name, icon_url, category_id)
VALUES ($1, $2, $3)
RETURNING tag_id, tag_name, icon_url, category_id
`

type CreateNewTagParams struct {
	TagName    string         `json:"tag_name"`
	IconUrl    sql.NullString `json:"icon_url"`
	CategoryID sql.NullInt32  `json:"category_id"`
}

func (q *Queries) CreateNewTag(ctx context.Context, arg CreateNewTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, createNewTag, arg.TagName, arg.IconUrl, arg.CategoryID)
	var i Tag
	err := row.Scan(
		&i.TagID,
		&i.TagName,
		&i.IconUrl,
		&i.CategoryID,
	)
	return i, err
}

const deleteGroup = `-- name: DeleteGroup :exec
DELETE FROM groups
WHERE group_id = $1
`

func (q *Queries) DeleteGroup(ctx context.Context, groupID int32) error {
	_, err := q.db.ExecContext(ctx, deleteGroup, groupID)
	return err
}

const deleteGroupMembers = `-- name: DeleteGroupMembers :exec
DELETE FROM group_members WHERE member_id = $1
`

func (q *Queries) DeleteGroupMembers(ctx context.Context, memberID int32) error {
	_, err := q.db.ExecContext(ctx, deleteGroupMembers, memberID)
	return err
}

const deleteGroupRequests = `-- name: DeleteGroupRequests :exec
DELETE FROM group_requests WHERE member_id = $1
`

func (q *Queries) DeleteGroupRequests(ctx context.Context, memberID int32) error {
	_, err := q.db.ExecContext(ctx, deleteGroupRequests, memberID)
	return err
}

const deleteMember = `-- name: DeleteMember :exec
DELETE FROM group_members
WHERE member_id = $1
    AND group_id = $2
`

type DeleteMemberParams struct {
	MemberID int32 `json:"member_id"`
	GroupID  int32 `json:"group_id"`
}

func (q *Queries) DeleteMember(ctx context.Context, arg DeleteMemberParams) error {
	_, err := q.db.ExecContext(ctx, deleteMember, arg.MemberID, arg.GroupID)
	return err
}

const deletePostComments = `-- name: DeletePostComments :exec
DELETE FROM post_comments WHERE member_id = $1
`

func (q *Queries) DeletePostComments(ctx context.Context, memberID int32) error {
	_, err := q.db.ExecContext(ctx, deletePostComments, memberID)
	return err
}

const deletePostReactions = `-- name: DeletePostReactions :exec
DELETE FROM post_reactions WHERE member_id = $1
`

func (q *Queries) DeletePostReactions(ctx context.Context, memberID int32) error {
	_, err := q.db.ExecContext(ctx, deletePostReactions, memberID)
	return err
}

const deletePosts = `-- name: DeletePosts :exec
DELETE FROM posts WHERE member_id = $1
`

func (q *Queries) DeletePosts(ctx context.Context, memberID int32) error {
	_, err := q.db.ExecContext(ctx, deletePosts, memberID)
	return err
}

const deleteRequestToJoinGroup = `-- name: DeleteRequestToJoinGroup :exec
DELETE FROM group_requests
WHERE group_id = $1
    AND member_id = $2
`

type DeleteRequestToJoinGroupParams struct {
	GroupID  int32 `json:"group_id"`
	MemberID int32 `json:"member_id"`
}

func (q *Queries) DeleteRequestToJoinGroup(ctx context.Context, arg DeleteRequestToJoinGroupParams) error {
	_, err := q.db.ExecContext(ctx, deleteRequestToJoinGroup, arg.GroupID, arg.MemberID)
	return err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE FROM tags
WHERE tag_id = $1
`

func (q *Queries) DeleteTag(ctx context.Context, tagID int32) error {
	_, err := q.db.ExecContext(ctx, deleteTag, tagID)
	return err
}

const editGroupDescription = `-- name: EditGroupDescription :exec
UPDATE groups
SET description = $2
WHERE group_id = $1
`

type EditGroupDescriptionParams struct {
	GroupID     int32          `json:"group_id"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) EditGroupDescription(ctx context.Context, arg EditGroupDescriptionParams) error {
	_, err := q.db.ExecContext(ctx, editGroupDescription, arg.GroupID, arg.Description)
	return err
}

const editGroupName = `-- name: EditGroupName :exec
UPDATE groups
SET group_name = $2
WHERE group_id = $1
`

type EditGroupNameParams struct {
	GroupID   int32  `json:"group_id"`
	GroupName string `json:"group_name"`
}

func (q *Queries) EditGroupName(ctx context.Context, arg EditGroupNameParams) error {
	_, err := q.db.ExecContext(ctx, editGroupName, arg.GroupID, arg.GroupName)
	return err
}

const editGroupPhoto = `-- name: EditGroupPhoto :exec
UPDATE groups
SET photo_url = $2
WHERE group_id = $1
`

type EditGroupPhotoParams struct {
	GroupID  int32          `json:"group_id"`
	PhotoUrl sql.NullString `json:"photo_url"`
}

func (q *Queries) EditGroupPhoto(ctx context.Context, arg EditGroupPhotoParams) error {
	_, err := q.db.ExecContext(ctx, editGroupPhoto, arg.GroupID, arg.PhotoUrl)
	return err
}

const editGroupTag = `-- name: EditGroupTag :one
UPDATE groups 
SET tag_id = $2
WHERE group_id = $1
RETURNING group_id, group_name, group_creator_id, description, photo_url, tag_id, frequency, max_members, group_type, visibility, created_at
`

type EditGroupTagParams struct {
	GroupID int32 `json:"group_id"`
	TagID   int32 `json:"tag_id"`
}

func (q *Queries) EditGroupTag(ctx context.Context, arg EditGroupTagParams) (Group, error) {
	row := q.db.QueryRowContext(ctx, editGroupTag, arg.GroupID, arg.TagID)
	var i Group
	err := row.Scan(
		&i.GroupID,
		&i.GroupName,
		&i.GroupCreatorID,
		&i.Description,
		&i.PhotoUrl,
		&i.TagID,
		&i.Frequency,
		&i.MaxMembers,
		&i.GroupType,
		&i.Visibility,
		&i.CreatedAt,
	)
	return i, err
}

const editGroupVisibility = `-- name: EditGroupVisibility :exec
UPDATE groups
SET visibility = $2
WHERE group_id = $1
`

type EditGroupVisibilityParams struct {
	GroupID    int32 `json:"group_id"`
	Visibility bool  `json:"visibility"`
}

func (q *Queries) EditGroupVisibility(ctx context.Context, arg EditGroupVisibilityParams) error {
	_, err := q.db.ExecContext(ctx, editGroupVisibility, arg.GroupID, arg.Visibility)
	return err
}

const editMemberRole = `-- name: EditMemberRole :exec
UPDATE group_members
SET role = $3
WHERE group_id = $1
    AND member_id = $2
`

type EditMemberRoleParams struct {
	GroupID  int32  `json:"group_id"`
	MemberID int32  `json:"member_id"`
	Role     string `json:"role"`
}

func (q *Queries) EditMemberRole(ctx context.Context, arg EditMemberRoleParams) error {
	_, err := q.db.ExecContext(ctx, editMemberRole, arg.GroupID, arg.MemberID, arg.Role)
	return err
}

const editTagCategory = `-- name: EditTagCategory :exec
UPDATE tags
SET category_id = $2
WHERE tag_id = $1
`

type EditTagCategoryParams struct {
	TagID      int32         `json:"tag_id"`
	CategoryID sql.NullInt32 `json:"category_id"`
}

func (q *Queries) EditTagCategory(ctx context.Context, arg EditTagCategoryParams) error {
	_, err := q.db.ExecContext(ctx, editTagCategory, arg.TagID, arg.CategoryID)
	return err
}

const editTagIcon = `-- name: EditTagIcon :exec
UPDATE tags
SET icon_url = $2
WHERE tag_id = $1
`

type EditTagIconParams struct {
	TagID   int32          `json:"tag_id"`
	IconUrl sql.NullString `json:"icon_url"`
}

func (q *Queries) EditTagIcon(ctx context.Context, arg EditTagIconParams) error {
	_, err := q.db.ExecContext(ctx, editTagIcon, arg.TagID, arg.IconUrl)
	return err
}

const editTagName = `-- name: EditTagName :exec
UPDATE tags
SET tag_name = $2
WHERE tag_id = $1
`

type EditTagNameParams struct {
	TagID   int32  `json:"tag_id"`
	TagName string `json:"tag_name"`
}

func (q *Queries) EditTagName(ctx context.Context, arg EditTagNameParams) error {
	_, err := q.db.ExecContext(ctx, editTagName, arg.TagID, arg.TagName)
	return err
}

const getAcceptorGroupRequests = `-- name: GetAcceptorGroupRequests :many
SELECT g.group_id, g.group_name, g.photo_url,
       COUNT(*) AS request_count
FROM group_requests gr 
JOIN group_members gm ON gr.group_id = gm.group_id
JOIN groups g ON gr.group_id = g.group_id
WHERE gm.member_id = $1 AND gm.role != 'member'
GROUP BY g.group_id, g.group_name, g.photo_url
`

type GetAcceptorGroupRequestsRow struct {
	GroupID      int32          `json:"group_id"`
	GroupName    string         `json:"group_name"`
	PhotoUrl     sql.NullString `json:"photo_url"`
	RequestCount int64          `json:"request_count"`
}

func (q *Queries) GetAcceptorGroupRequests(ctx context.Context, memberID int32) ([]GetAcceptorGroupRequestsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAcceptorGroupRequests, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAcceptorGroupRequestsRow{}
	for rows.Next() {
		var i GetAcceptorGroupRequestsRow
		if err := rows.Scan(
			&i.GroupID,
			&i.GroupName,
			&i.PhotoUrl,
			&i.RequestCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAcceptorGroupRequestsCount = `-- name: GetAcceptorGroupRequestsCount :one
SELECT gm.member_id,
       COUNT(*) AS request_count
FROM group_requests gr 
JOIN group_members gm ON gr.group_id = gm.group_id
WHERE gm.member_id = $1 AND gm.role != 'member'
GROUP BY gm.member_id
`

type GetAcceptorGroupRequestsCountRow struct {
	MemberID     int32 `json:"member_id"`
	RequestCount int64 `json:"request_count"`
}

func (q *Queries) GetAcceptorGroupRequestsCount(ctx context.Context, memberID int32) (GetAcceptorGroupRequestsCountRow, error) {
	row := q.db.QueryRowContext(ctx, getAcceptorGroupRequestsCount, memberID)
	var i GetAcceptorGroupRequestsCountRow
	err := row.Scan(&i.MemberID, &i.RequestCount)
	return i, err
}

const getAvailableCategory = `-- name: GetAvailableCategory :many
SELECT category_id, category_name
FROM tag_category
ORDER BY category_id
`

func (q *Queries) GetAvailableCategory(ctx context.Context) ([]TagCategory, error) {
	rows, err := q.db.QueryContext(ctx, getAvailableCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TagCategory{}
	for rows.Next() {
		var i TagCategory
		if err := rows.Scan(&i.CategoryID, &i.CategoryName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAvailableTags = `-- name: GetAvailableTags :many
SELECT tag_id, tag_name, icon_url, category_id
FROM tags
ORDER BY tag_id
`

func (q *Queries) GetAvailableTags(ctx context.Context) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, getAvailableTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.TagID,
			&i.TagName,
			&i.IconUrl,
			&i.CategoryID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroupByID = `-- name: GetGroupByID :one
SELECT groups.group_id,
    groups.group_name,
    groups.photo_url,
    groups.group_creator_id,
    groups.frequency,
    groups.max_members,
    groups.group_type,
    groups.visibility,
    groups.description,
    groups.created_at,
    tags.tag_name,
    (SELECT COUNT(*)
    FROM group_members A
    WHERE A.group_id = $1) AS total_member
FROM groups
JOIN tags ON groups.tag_id = tags.tag_id
WHERE group_id = $1
`

type GetGroupByIDRow struct {
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
	TotalMember    int64          `json:"total_member"`
}

func (q *Queries) GetGroupByID(ctx context.Context, groupID int32) (GetGroupByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getGroupByID, groupID)
	var i GetGroupByIDRow
	err := row.Scan(
		&i.GroupID,
		&i.GroupName,
		&i.PhotoUrl,
		&i.GroupCreatorID,
		&i.Frequency,
		&i.MaxMembers,
		&i.GroupType,
		&i.Visibility,
		&i.Description,
		&i.CreatedAt,
		&i.TagName,
		&i.TotalMember,
	)
	return i, err
}

const getGroupLatestId = `-- name: GetGroupLatestId :one
SELECT COALESCE(MAX(group_id), 0)::integer
FROM groups
`

func (q *Queries) GetGroupLatestId(ctx context.Context) (int32, error) {
	row := q.db.QueryRowContext(ctx, getGroupLatestId)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const getGroupRequest = `-- name: GetGroupRequest :one
SELECT group_id, member_id, description, created_at FROM group_requests
WHERE group_id = $1
    AND member_id = $2
`

type GetGroupRequestParams struct {
	GroupID  int32 `json:"group_id"`
	MemberID int32 `json:"member_id"`
}

func (q *Queries) GetGroupRequest(ctx context.Context, arg GetGroupRequestParams) (GroupRequest, error) {
	row := q.db.QueryRowContext(ctx, getGroupRequest, arg.GroupID, arg.MemberID)
	var i GroupRequest
	err := row.Scan(
		&i.GroupID,
		&i.MemberID,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getGroupRequests = `-- name: GetGroupRequests :many
SELECT group_id, member_id, description, created_at FROM group_requests
WHERE group_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetGroupRequestsParams struct {
	GroupID int32 `json:"group_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) GetGroupRequests(ctx context.Context, arg GetGroupRequestsParams) ([]GroupRequest, error) {
	rows, err := q.db.QueryContext(ctx, getGroupRequests, arg.GroupID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GroupRequest{}
	for rows.Next() {
		var i GroupRequest
		if err := rows.Scan(
			&i.GroupID,
			&i.MemberID,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroupsByCategoryID = `-- name: GetGroupsByCategoryID :many
SELECT groups.group_id, groups.group_name, groups.group_creator_id, groups.description, groups.photo_url, groups.tag_id, groups.frequency, groups.max_members, groups.group_type, groups.visibility, groups.created_at, tags.category_id, tag_category.category_name
FROM groups
JOIN tags ON groups.tag_id = tags.tag_id
JOIN tag_category ON tags.category_id = tag_category.category_id
WHERE tags.category_id = $1
`

type GetGroupsByCategoryIDRow struct {
	GroupID        int32          `json:"group_id"`
	GroupName      string         `json:"group_name"`
	GroupCreatorID int32          `json:"group_creator_id"`
	Description    sql.NullString `json:"description"`
	PhotoUrl       sql.NullString `json:"photo_url"`
	TagID          int32          `json:"tag_id"`
	Frequency      sql.NullInt32  `json:"frequency"`
	MaxMembers     int32          `json:"max_members"`
	GroupType      string         `json:"group_type"`
	Visibility     bool           `json:"visibility"`
	CreatedAt      time.Time      `json:"created_at"`
	CategoryID     sql.NullInt32  `json:"category_id"`
	CategoryName   string         `json:"category_name"`
}

func (q *Queries) GetGroupsByCategoryID(ctx context.Context, categoryID sql.NullInt32) ([]GetGroupsByCategoryIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getGroupsByCategoryID, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetGroupsByCategoryIDRow{}
	for rows.Next() {
		var i GetGroupsByCategoryIDRow
		if err := rows.Scan(
			&i.GroupID,
			&i.GroupName,
			&i.GroupCreatorID,
			&i.Description,
			&i.PhotoUrl,
			&i.TagID,
			&i.Frequency,
			&i.MaxMembers,
			&i.GroupType,
			&i.Visibility,
			&i.CreatedAt,
			&i.CategoryID,
			&i.CategoryName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroupsByTagID = `-- name: GetGroupsByTagID :many
SELECT group_id, group_name, group_creator_id, description, photo_url, tag_id, frequency, max_members, group_type, visibility, created_at
from groups
WHERE tag_id = $1
`

func (q *Queries) GetGroupsByTagID(ctx context.Context, tagID int32) ([]Group, error) {
	rows, err := q.db.QueryContext(ctx, getGroupsByTagID, tagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Group{}
	for rows.Next() {
		var i Group
		if err := rows.Scan(
			&i.GroupID,
			&i.GroupName,
			&i.GroupCreatorID,
			&i.Description,
			&i.PhotoUrl,
			&i.TagID,
			&i.Frequency,
			&i.MaxMembers,
			&i.GroupType,
			&i.Visibility,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMemberInfo = `-- name: GetMemberInfo :one
SELECT gm.group_id, gm.member_id, gm.role, gm.created_at, g.group_name
FROM group_members gm
JOIN groups g ON gm.group_id = g.group_id
WHERE gm.member_id = $1
AND gm.group_id = $2
`

type GetMemberInfoParams struct {
	MemberID int32 `json:"member_id"`
	GroupID  int32 `json:"group_id"`
}

type GetMemberInfoRow struct {
	GroupID   int32     `json:"group_id"`
	MemberID  int32     `json:"member_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	GroupName string    `json:"group_name"`
}

func (q *Queries) GetMemberInfo(ctx context.Context, arg GetMemberInfoParams) (GetMemberInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getMemberInfo, arg.MemberID, arg.GroupID)
	var i GetMemberInfoRow
	err := row.Scan(
		&i.GroupID,
		&i.MemberID,
		&i.Role,
		&i.CreatedAt,
		&i.GroupName,
	)
	return i, err
}

const getMemberPendingGroupRequests = `-- name: GetMemberPendingGroupRequests :many
SELECT group_id, member_id, description, created_at FROM group_requests
WHERE member_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetMemberPendingGroupRequestsParams struct {
	MemberID int32 `json:"member_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) GetMemberPendingGroupRequests(ctx context.Context, arg GetMemberPendingGroupRequestsParams) ([]GroupRequest, error) {
	rows, err := q.db.QueryContext(ctx, getMemberPendingGroupRequests, arg.MemberID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GroupRequest{}
	for rows.Next() {
		var i GroupRequest
		if err := rows.Scan(
			&i.GroupID,
			&i.MemberID,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMembersByGroupID = `-- name: GetMembersByGroupID :many
SELECT group_id, member_id, role, created_at
FROM group_members
WHERE group_id = $1
LIMIT $2 OFFSET $3
`

type GetMembersByGroupIDParams struct {
	GroupID int32 `json:"group_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) GetMembersByGroupID(ctx context.Context, arg GetMembersByGroupIDParams) ([]GroupMember, error) {
	rows, err := q.db.QueryContext(ctx, getMembersByGroupID, arg.GroupID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GroupMember{}
	for rows.Next() {
		var i GroupMember
		if err := rows.Scan(
			&i.GroupID,
			&i.MemberID,
			&i.Role,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReactionById = `-- name: GetReactionById :one
SELECT reaction_id, post_id, member_id, reaction, created_at
FROM post_reactions
WHERE reaction_id = $1
`

func (q *Queries) GetReactionById(ctx context.Context, reactionID int32) (PostReaction, error) {
	row := q.db.QueryRowContext(ctx, getReactionById, reactionID)
	var i PostReaction
	err := row.Scan(
		&i.ReactionID,
		&i.PostID,
		&i.MemberID,
		&i.Reaction,
		&i.CreatedAt,
	)
	return i, err
}

const getTagByCategory = `-- name: GetTagByCategory :many
SELECT tag_id, tag_name, icon_url, category_id
FROM tags
WHERE category_id = $1
`

func (q *Queries) GetTagByCategory(ctx context.Context, categoryID sql.NullInt32) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, getTagByCategory, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.TagID,
			&i.TagName,
			&i.IconUrl,
			&i.CategoryID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTagByGroupId = `-- name: GetTagByGroupId :one
SELECT groups.group_id, groups.group_name, tags.tag_id, tags.tag_name, tags.icon_url, tags.category_id, tag_category.category_name
FROM groups
JOIN tags ON groups.tag_id = tags.tag_id
JOIN tag_category ON tags.category_id = tag_category.category_id
WHERE groups.group_id = $1
`

type GetTagByGroupIdRow struct {
	GroupID      int32          `json:"group_id"`
	GroupName    string         `json:"group_name"`
	TagID        int32          `json:"tag_id"`
	TagName      string         `json:"tag_name"`
	IconUrl      sql.NullString `json:"icon_url"`
	CategoryID   sql.NullInt32  `json:"category_id"`
	CategoryName string         `json:"category_name"`
}

func (q *Queries) GetTagByGroupId(ctx context.Context, groupID int32) (GetTagByGroupIdRow, error) {
	row := q.db.QueryRowContext(ctx, getTagByGroupId, groupID)
	var i GetTagByGroupIdRow
	err := row.Scan(
		&i.GroupID,
		&i.GroupName,
		&i.TagID,
		&i.TagName,
		&i.IconUrl,
		&i.CategoryID,
		&i.CategoryName,
	)
	return i, err
}

const getTagByTagID = `-- name: GetTagByTagID :one
SELECT tag_id, tag_name, icon_url, category_id
FROM tags
WHERE tag_id = $1
`

func (q *Queries) GetTagByTagID(ctx context.Context, tagID int32) (Tag, error) {
	row := q.db.QueryRowContext(ctx, getTagByTagID, tagID)
	var i Tag
	err := row.Scan(
		&i.TagID,
		&i.TagName,
		&i.IconUrl,
		&i.CategoryID,
	)
	return i, err
}

const getUserGroups = `-- name: GetUserGroups :many
SELECT g.group_id, g.group_name, g.photo_url, g.group_type
FROM group_members gm
JOIN groups g ON g.group_id = gm.group_id
WHERE gm.member_id = $1
`

type GetUserGroupsRow struct {
	GroupID   int32          `json:"group_id"`
	GroupName string         `json:"group_name"`
	PhotoUrl  sql.NullString `json:"photo_url"`
	GroupType string         `json:"group_type"`
}

func (q *Queries) GetUserGroups(ctx context.Context, memberID int32) ([]GetUserGroupsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserGroups, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserGroupsRow{}
	for rows.Next() {
		var i GetUserGroupsRow
		if err := rows.Scan(
			&i.GroupID,
			&i.GroupName,
			&i.PhotoUrl,
			&i.GroupType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const initializeCategory = `-- name: InitializeCategory :exec
INSERT INTO tag_category (category_name)
VALUES ('Sports and Fitness'), 
        ('Learning and Development'),
		('Health and Wellness'), 
        ('Entertainment and Media'), 
        ('Hobbies and Leisure'),
        ('Others')
`

func (q *Queries) InitializeCategory(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, initializeCategory)
	return err
}

const listGroups = `-- name: ListGroups :many
SELECT groups.group_id,
    groups.group_name,
    groups.photo_url,
    groups.group_creator_id,
    groups.frequency,
    groups.max_members,
    groups.group_type,
    groups.visibility,
    groups.description,
    groups.created_at,
    tags.tag_name,
    (SELECT COUNT(*)
    FROM group_members S
    WHERE S.group_id = groups.group_id) AS total_member
FROM groups
JOIN tags ON groups.tag_id = tags.tag_id
ORDER BY group_id
LIMIT $1 OFFSET $2
`

type ListGroupsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListGroupsRow struct {
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
	TotalMember    int64          `json:"total_member"`
}

func (q *Queries) ListGroups(ctx context.Context, arg ListGroupsParams) ([]ListGroupsRow, error) {
	rows, err := q.db.QueryContext(ctx, listGroups, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListGroupsRow{}
	for rows.Next() {
		var i ListGroupsRow
		if err := rows.Scan(
			&i.GroupID,
			&i.GroupName,
			&i.PhotoUrl,
			&i.GroupCreatorID,
			&i.Frequency,
			&i.MaxMembers,
			&i.GroupType,
			&i.Visibility,
			&i.Description,
			&i.CreatedAt,
			&i.TagName,
			&i.TotalMember,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const numberMemberInGroup = `-- name: NumberMemberInGroup :one
SELECT COUNT(*)
FROM group_members
WHERE group_id = $1
`

func (q *Queries) NumberMemberInGroup(ctx context.Context, groupID int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, numberMemberInGroup, groupID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const searchGroupByGroupName = `-- name: SearchGroupByGroupName :many
SELECT groups.group_id,
    groups.group_name,
    groups.photo_url,
    groups.group_creator_id,
    groups.frequency,
    groups.max_members,
    groups.group_type,
    groups.visibility,
    groups.description,
    groups.created_at,
    tags.tag_name,
    (SELECT COUNT(*)
    FROM group_members A
    WHERE A.group_id = groups.group_id) AS total_member
FROM groups
JOIN tags ON groups.tag_id = tags.tag_id
WHERE groups.group_name ILIKE '%' || $1 || '%'
LIMIT $2 OFFSET $3
`

type SearchGroupByGroupNameParams struct {
	Column1 sql.NullString `json:"column_1"`
	Limit   int32          `json:"limit"`
	Offset  int32          `json:"offset"`
}

type SearchGroupByGroupNameRow struct {
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
	TotalMember    int64          `json:"total_member"`
}

func (q *Queries) SearchGroupByGroupName(ctx context.Context, arg SearchGroupByGroupNameParams) ([]SearchGroupByGroupNameRow, error) {
	rows, err := q.db.QueryContext(ctx, searchGroupByGroupName, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchGroupByGroupNameRow{}
	for rows.Next() {
		var i SearchGroupByGroupNameRow
		if err := rows.Scan(
			&i.GroupID,
			&i.GroupName,
			&i.PhotoUrl,
			&i.GroupCreatorID,
			&i.Frequency,
			&i.MaxMembers,
			&i.GroupType,
			&i.Visibility,
			&i.Description,
			&i.CreatedAt,
			&i.TagName,
			&i.TotalMember,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const sendRequestToJoinGroup = `-- name: SendRequestToJoinGroup :one
INSERT INTO group_requests (group_id, member_id, description)
VALUES ($1, $2, $3)
RETURNING group_id, member_id, description, created_at
`

type SendRequestToJoinGroupParams struct {
	GroupID     int32          `json:"group_id"`
	MemberID    int32          `json:"member_id"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) SendRequestToJoinGroup(ctx context.Context, arg SendRequestToJoinGroupParams) (GroupRequest, error) {
	row := q.db.QueryRowContext(ctx, sendRequestToJoinGroup, arg.GroupID, arg.MemberID, arg.Description)
	var i GroupRequest
	err := row.Scan(
		&i.GroupID,
		&i.MemberID,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}
