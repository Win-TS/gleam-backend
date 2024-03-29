// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: streak.sql

package groupdb

import (
	"context"
	"database/sql"
	"time"
)

const createStreak = `-- name: CreateStreak :one
INSERT INTO streaks (
    streak_set_id,
    post_id,
    streak_count
) VALUES (
    $1, $2, $3
) RETURNING streak_id, streak_set_id, post_id, streak_count, created_at
`

type CreateStreakParams struct {
	StreakSetID int32         `json:"streak_set_id"`
	PostID      int32         `json:"post_id"`
	StreakCount sql.NullInt32 `json:"streak_count"`
}

func (q *Queries) CreateStreak(ctx context.Context, arg CreateStreakParams) (Streak, error) {
	row := q.db.QueryRowContext(ctx, createStreak, arg.StreakSetID, arg.PostID, arg.StreakCount)
	var i Streak
	err := row.Scan(
		&i.StreakID,
		&i.StreakSetID,
		&i.PostID,
		&i.StreakCount,
		&i.CreatedAt,
	)
	return i, err
}

const createStreakSet = `-- name: CreateStreakSet :one
INSERT INTO streak_set (
    group_id,
    user_id,
    streak_count
) VALUES (
    $1, $2, $3
) RETURNING streak_set_id, group_id, user_id, streak_count, ended, created_at
`

type CreateStreakSetParams struct {
	GroupID     int32         `json:"group_id"`
	UserID      int32         `json:"user_id"`
	StreakCount sql.NullInt32 `json:"streak_count"`
}

func (q *Queries) CreateStreakSet(ctx context.Context, arg CreateStreakSetParams) (StreakSet, error) {
	row := q.db.QueryRowContext(ctx, createStreakSet, arg.GroupID, arg.UserID, arg.StreakCount)
	var i StreakSet
	err := row.Scan(
		&i.StreakSetID,
		&i.GroupID,
		&i.UserID,
		&i.StreakCount,
		&i.Ended,
		&i.CreatedAt,
	)
	return i, err
}

const endStreakSet = `-- name: EndStreakSet :exec
UPDATE streak_set SET ended = true
WHERE streak_set_id = $1
`

func (q *Queries) EndStreakSet(ctx context.Context, streakSetID int32) error {
	_, err := q.db.ExecContext(ctx, endStreakSet, streakSetID)
	return err
}

const getLatestStreakSetByGroupIDAndUserID = `-- name: GetLatestStreakSetByGroupIDAndUserID :one
SELECT streak_set_id, group_id, user_id, streak_count, ended, created_at FROM streak_set
WHERE group_id = $1 AND user_id = $2
ORDER BY created_at DESC 
LIMIT 1
`

type GetLatestStreakSetByGroupIDAndUserIDParams struct {
	GroupID int32 `json:"group_id"`
	UserID  int32 `json:"user_id"`
}

func (q *Queries) GetLatestStreakSetByGroupIDAndUserID(ctx context.Context, arg GetLatestStreakSetByGroupIDAndUserIDParams) (StreakSet, error) {
	row := q.db.QueryRowContext(ctx, getLatestStreakSetByGroupIDAndUserID, arg.GroupID, arg.UserID)
	var i StreakSet
	err := row.Scan(
		&i.StreakSetID,
		&i.GroupID,
		&i.UserID,
		&i.StreakCount,
		&i.Ended,
		&i.CreatedAt,
	)
	return i, err
}

const getStreakByPostID = `-- name: GetStreakByPostID :one
SELECT streak_id, streak_set_id, post_id, streak_count, created_at FROM streaks
WHERE post_id = $1
`

func (q *Queries) GetStreakByPostID(ctx context.Context, postID int32) (Streak, error) {
	row := q.db.QueryRowContext(ctx, getStreakByPostID, postID)
	var i Streak
	err := row.Scan(
		&i.StreakID,
		&i.StreakSetID,
		&i.PostID,
		&i.StreakCount,
		&i.CreatedAt,
	)
	return i, err
}

const getStreakSetByUserID = `-- name: GetStreakSetByUserID :many
SELECT streak_set_id, group_id, user_id, streak_count, ended, created_at FROM streak_set
WHERE user_id = $1
`

func (q *Queries) GetStreakSetByUserID(ctx context.Context, userID int32) ([]StreakSet, error) {
	rows, err := q.db.QueryContext(ctx, getStreakSetByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []StreakSet{}
	for rows.Next() {
		var i StreakSet
		if err := rows.Scan(
			&i.StreakSetID,
			&i.GroupID,
			&i.UserID,
			&i.StreakCount,
			&i.Ended,
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

const getStreaksByGroupIDAndUserID = `-- name: GetStreaksByGroupIDAndUserID :many
SELECT streak_set.streak_set_id, group_id, user_id, streak_set.streak_count, ended, streak_set.created_at, streak_id, streaks.streak_set_id, post_id, streaks.streak_count, streaks.created_at FROM streak_set
JOIN streaks ON streak_set.streak_set_id = streaks.streak_set_id
WHERE group_id = $1 AND user_id = $2
ORDER BY streak_set.created_at DESC
`

type GetStreaksByGroupIDAndUserIDParams struct {
	GroupID int32 `json:"group_id"`
	UserID  int32 `json:"user_id"`
}

type GetStreaksByGroupIDAndUserIDRow struct {
	StreakSetID   int32         `json:"streak_set_id"`
	GroupID       int32         `json:"group_id"`
	UserID        int32         `json:"user_id"`
	StreakCount   sql.NullInt32 `json:"streak_count"`
	Ended         bool          `json:"ended"`
	CreatedAt     time.Time     `json:"created_at"`
	StreakID      int32         `json:"streak_id"`
	StreakSetID_2 int32         `json:"streak_set_id_2"`
	PostID        int32         `json:"post_id"`
	StreakCount_2 sql.NullInt32 `json:"streak_count_2"`
	CreatedAt_2   time.Time     `json:"created_at_2"`
}

func (q *Queries) GetStreaksByGroupIDAndUserID(ctx context.Context, arg GetStreaksByGroupIDAndUserIDParams) ([]GetStreaksByGroupIDAndUserIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getStreaksByGroupIDAndUserID, arg.GroupID, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetStreaksByGroupIDAndUserIDRow{}
	for rows.Next() {
		var i GetStreaksByGroupIDAndUserIDRow
		if err := rows.Scan(
			&i.StreakSetID,
			&i.GroupID,
			&i.UserID,
			&i.StreakCount,
			&i.Ended,
			&i.CreatedAt,
			&i.StreakID,
			&i.StreakSetID_2,
			&i.PostID,
			&i.StreakCount_2,
			&i.CreatedAt_2,
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

const getStreaksByStreakSetID = `-- name: GetStreaksByStreakSetID :many
SELECT streak_id, streak_set_id, post_id, streak_count, created_at FROM streaks
WHERE streak_set_id = $1
`

func (q *Queries) GetStreaksByStreakSetID(ctx context.Context, streakSetID int32) ([]Streak, error) {
	rows, err := q.db.QueryContext(ctx, getStreaksByStreakSetID, streakSetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Streak{}
	for rows.Next() {
		var i Streak
		if err := rows.Scan(
			&i.StreakID,
			&i.StreakSetID,
			&i.PostID,
			&i.StreakCount,
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

const getUnendedStreakSetByUserID = `-- name: GetUnendedStreakSetByUserID :many
SELECT streak_set_id, group_id, user_id, streak_count, ended, created_at FROM streak_set
WHERE user_id = $1 AND ended = false
`

func (q *Queries) GetUnendedStreakSetByUserID(ctx context.Context, userID int32) ([]StreakSet, error) {
	rows, err := q.db.QueryContext(ctx, getUnendedStreakSetByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []StreakSet{}
	for rows.Next() {
		var i StreakSet
		if err := rows.Scan(
			&i.StreakSetID,
			&i.GroupID,
			&i.UserID,
			&i.StreakCount,
			&i.Ended,
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

const updateStreakSetCount = `-- name: UpdateStreakSetCount :exec
UPDATE streak_set SET streak_count = $2
WHERE streak_set_id = $1
`

type UpdateStreakSetCountParams struct {
	StreakSetID int32         `json:"streak_set_id"`
	StreakCount sql.NullInt32 `json:"streak_count"`
}

func (q *Queries) UpdateStreakSetCount(ctx context.Context, arg UpdateStreakSetCountParams) error {
	_, err := q.db.ExecContext(ctx, updateStreakSetCount, arg.StreakSetID, arg.StreakCount)
	return err
}
