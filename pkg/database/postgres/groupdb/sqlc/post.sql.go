// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: post.sql

package groupdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createComment = `-- name: CreateComment :one
INSERT INTO post_comments (
    post_id,
    member_id,
    comment
) VALUES (
    $1, $2, $3
) RETURNING comment_id, post_id, member_id, comment, created_at
`

type CreateCommentParams struct {
	PostID   int32  `json:"post_id"`
	MemberID int32  `json:"member_id"`
	Comment  string `json:"comment"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (PostComment, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.PostID, arg.MemberID, arg.Comment)
	var i PostComment
	err := row.Scan(
		&i.CommentID,
		&i.PostID,
		&i.MemberID,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
    member_id,
    group_id,
    photo_url,
    description
) VALUES (
    $1, $2, $3, $4
) RETURNING post_id, member_id, group_id, photo_url, description, created_at
`

type CreatePostParams struct {
	MemberID    int32          `json:"member_id"`
	GroupID     int32          `json:"group_id"`
	PhotoUrl    sql.NullString `json:"photo_url"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.MemberID,
		arg.GroupID,
		arg.PhotoUrl,
		arg.Description,
	)
	var i Post
	err := row.Scan(
		&i.PostID,
		&i.MemberID,
		&i.GroupID,
		&i.PhotoUrl,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const createReaction = `-- name: CreateReaction :one
INSERT INTO post_reactions (
    post_id,
    member_id,
    reaction
) VALUES (
    $1, $2, $3
) RETURNING reaction_id, post_id, member_id, reaction, created_at
`

type CreateReactionParams struct {
	PostID   int32  `json:"post_id"`
	MemberID int32  `json:"member_id"`
	Reaction string `json:"reaction"`
}

func (q *Queries) CreateReaction(ctx context.Context, arg CreateReactionParams) (PostReaction, error) {
	row := q.db.QueryRowContext(ctx, createReaction, arg.PostID, arg.MemberID, arg.Reaction)
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

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM post_comments
WHERE comment_id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, commentID int32) error {
	_, err := q.db.ExecContext(ctx, deleteComment, commentID)
	return err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE post_id = $1
`

func (q *Queries) DeletePost(ctx context.Context, postID int32) error {
	_, err := q.db.ExecContext(ctx, deletePost, postID)
	return err
}

const deleteReaction = `-- name: DeleteReaction :exec
DELETE FROM post_reactions
WHERE post_id = $1 AND member_id = $2 AND reaction = $3
`

type DeleteReactionParams struct {
	PostID   int32  `json:"post_id"`
	MemberID int32  `json:"member_id"`
	Reaction string `json:"reaction"`
}

func (q *Queries) DeleteReaction(ctx context.Context, arg DeleteReactionParams) error {
	_, err := q.db.ExecContext(ctx, deleteReaction, arg.PostID, arg.MemberID, arg.Reaction)
	return err
}

const editComment = `-- name: EditComment :exec
UPDATE post_comments SET comment = $2
WHERE comment_id = $1
`

type EditCommentParams struct {
	CommentID int32  `json:"comment_id"`
	Comment   string `json:"comment"`
}

func (q *Queries) EditComment(ctx context.Context, arg EditCommentParams) error {
	_, err := q.db.ExecContext(ctx, editComment, arg.CommentID, arg.Comment)
	return err
}

const editPost = `-- name: EditPost :exec
UPDATE posts SET description = $2
WHERE post_id = $1
`

type EditPostParams struct {
	PostID      int32          `json:"post_id"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) EditPost(ctx context.Context, arg EditPostParams) error {
	_, err := q.db.ExecContext(ctx, editPost, arg.PostID, arg.Description)
	return err
}

const editReaction = `-- name: EditReaction :exec
UPDATE post_reactions SET reaction = $2
WHERE reaction_id = $1
`

type EditReactionParams struct {
	ReactionID int32  `json:"reaction_id"`
	Reaction   string `json:"reaction"`
}

func (q *Queries) EditReaction(ctx context.Context, arg EditReactionParams) error {
	_, err := q.db.ExecContext(ctx, editReaction, arg.ReactionID, arg.Reaction)
	return err
}

const getCommentByCommentId = `-- name: GetCommentByCommentId :one
SELECT comment_id, post_id, member_id, comment, created_at
FROM post_comments
WHERE comment_id = $1
`

func (q *Queries) GetCommentByCommentId(ctx context.Context, commentID int32) (PostComment, error) {
	row := q.db.QueryRowContext(ctx, getCommentByCommentId, commentID)
	var i PostComment
	err := row.Scan(
		&i.CommentID,
		&i.PostID,
		&i.MemberID,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const getCommentsByPostID = `-- name: GetCommentsByPostID :many
SELECT comment_id, post_id, member_id, comment, created_at FROM post_comments
WHERE post_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetCommentsByPostIDParams struct {
	PostID int32 `json:"post_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetCommentsByPostID(ctx context.Context, arg GetCommentsByPostIDParams) ([]PostComment, error) {
	rows, err := q.db.QueryContext(ctx, getCommentsByPostID, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PostComment{}
	for rows.Next() {
		var i PostComment
		if err := rows.Scan(
			&i.CommentID,
			&i.PostID,
			&i.MemberID,
			&i.Comment,
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

const getCommentsCountByPostID = `-- name: GetCommentsCountByPostID :one
SELECT COUNT(*) FROM post_comments
WHERE post_id = $1
`

func (q *Queries) GetCommentsCountByPostID(ctx context.Context, postID int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCommentsCountByPostID, postID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getPostByPostID = `-- name: GetPostByPostID :one
SELECT post_id, member_id, group_id, photo_url, description, created_at FROM posts
WHERE post_id = $1
`

func (q *Queries) GetPostByPostID(ctx context.Context, postID int32) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPostByPostID, postID)
	var i Post
	err := row.Scan(
		&i.PostID,
		&i.MemberID,
		&i.GroupID,
		&i.PhotoUrl,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getPostLatestId = `-- name: GetPostLatestId :one
SELECT COALESCE(MAX(post_id), 0)::integer FROM posts
`

func (q *Queries) GetPostLatestId(ctx context.Context) (int32, error) {
	row := q.db.QueryRowContext(ctx, getPostLatestId)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const getPostsByGroupAndMemberID = `-- name: GetPostsByGroupAndMemberID :many
SELECT post_id, member_id, group_id, photo_url, description, created_at FROM posts
WHERE group_id = $1 AND member_id = $2
ORDER BY created_at DESC
LIMIT $3 OFFSET $4
`

type GetPostsByGroupAndMemberIDParams struct {
	GroupID  int32 `json:"group_id"`
	MemberID int32 `json:"member_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) GetPostsByGroupAndMemberID(ctx context.Context, arg GetPostsByGroupAndMemberIDParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByGroupAndMemberID,
		arg.GroupID,
		arg.MemberID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.PostID,
			&i.MemberID,
			&i.GroupID,
			&i.PhotoUrl,
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

const getPostsByGroupID = `-- name: GetPostsByGroupID :many
SELECT post_id, member_id, group_id, photo_url, description, created_at FROM posts
WHERE group_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetPostsByGroupIDParams struct {
	GroupID int32 `json:"group_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) GetPostsByGroupID(ctx context.Context, arg GetPostsByGroupIDParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByGroupID, arg.GroupID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.PostID,
			&i.MemberID,
			&i.GroupID,
			&i.PhotoUrl,
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

const getPostsByMemberID = `-- name: GetPostsByMemberID :many
SELECT post_id, member_id, group_id, photo_url, description, created_at FROM posts
WHERE member_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetPostsByMemberIDParams struct {
	MemberID int32 `json:"member_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) GetPostsByMemberID(ctx context.Context, arg GetPostsByMemberIDParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByMemberID, arg.MemberID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.PostID,
			&i.MemberID,
			&i.GroupID,
			&i.PhotoUrl,
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

const getPostsForFollowingFeedByMemberId = `-- name: GetPostsForFollowingFeedByMemberId :many
SELECT posts.post_id, posts.member_id, posts.group_id, posts.photo_url, posts.description, posts.created_at, groups.group_name, groups.photo_url AS group_photo_url FROM posts
JOIN groups ON posts.group_id = groups.group_id
WHERE posts.member_id = ANY($1::int[]) AND visibility = true
ORDER BY posts.created_at DESC
LIMIT $2 OFFSET $3
`

type GetPostsForFollowingFeedByMemberIdParams struct {
	Column1 []int32 `json:"column_1"`
	Limit   int32   `json:"limit"`
	Offset  int32   `json:"offset"`
}

type GetPostsForFollowingFeedByMemberIdRow struct {
	PostID        int32          `json:"post_id"`
	MemberID      int32          `json:"member_id"`
	GroupID       int32          `json:"group_id"`
	PhotoUrl      sql.NullString `json:"photo_url"`
	Description   sql.NullString `json:"description"`
	CreatedAt     time.Time      `json:"created_at"`
	GroupName     string         `json:"group_name"`
	GroupPhotoUrl sql.NullString `json:"group_photo_url"`
}

func (q *Queries) GetPostsForFollowingFeedByMemberId(ctx context.Context, arg GetPostsForFollowingFeedByMemberIdParams) ([]GetPostsForFollowingFeedByMemberIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsForFollowingFeedByMemberId, pq.Array(arg.Column1), arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsForFollowingFeedByMemberIdRow{}
	for rows.Next() {
		var i GetPostsForFollowingFeedByMemberIdRow
		if err := rows.Scan(
			&i.PostID,
			&i.MemberID,
			&i.GroupID,
			&i.PhotoUrl,
			&i.Description,
			&i.CreatedAt,
			&i.GroupName,
			&i.GroupPhotoUrl,
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

const getPostsForOngoingFeedByMemberID = `-- name: GetPostsForOngoingFeedByMemberID :many
SELECT DISTINCT posts.post_id, posts.member_id, posts.group_id, posts.photo_url, posts.description, posts.created_at, groups.group_name, groups.photo_url AS group_photo_url FROM posts
JOIN group_members ON posts.group_id = group_members.group_id JOIN groups ON posts.group_id = groups.group_id
WHERE group_members.member_id = $1
ORDER BY posts.created_at DESC
LIMIT $2 OFFSET $3
`

type GetPostsForOngoingFeedByMemberIDParams struct {
	MemberID int32 `json:"member_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

type GetPostsForOngoingFeedByMemberIDRow struct {
	PostID        int32          `json:"post_id"`
	MemberID      int32          `json:"member_id"`
	GroupID       int32          `json:"group_id"`
	PhotoUrl      sql.NullString `json:"photo_url"`
	Description   sql.NullString `json:"description"`
	CreatedAt     time.Time      `json:"created_at"`
	GroupName     string         `json:"group_name"`
	GroupPhotoUrl sql.NullString `json:"group_photo_url"`
}

func (q *Queries) GetPostsForOngoingFeedByMemberID(ctx context.Context, arg GetPostsForOngoingFeedByMemberIDParams) ([]GetPostsForOngoingFeedByMemberIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsForOngoingFeedByMemberID, arg.MemberID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsForOngoingFeedByMemberIDRow{}
	for rows.Next() {
		var i GetPostsForOngoingFeedByMemberIDRow
		if err := rows.Scan(
			&i.PostID,
			&i.MemberID,
			&i.GroupID,
			&i.PhotoUrl,
			&i.Description,
			&i.CreatedAt,
			&i.GroupName,
			&i.GroupPhotoUrl,
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

const getReactionByPostIDAndUserID = `-- name: GetReactionByPostIDAndUserID :one
SELECT reaction_id, post_id, member_id, reaction, created_at FROM post_reactions
WHERE post_id = $1 AND member_id = $2
`

type GetReactionByPostIDAndUserIDParams struct {
	PostID   int32 `json:"post_id"`
	MemberID int32 `json:"member_id"`
}

func (q *Queries) GetReactionByPostIDAndUserID(ctx context.Context, arg GetReactionByPostIDAndUserIDParams) (PostReaction, error) {
	row := q.db.QueryRowContext(ctx, getReactionByPostIDAndUserID, arg.PostID, arg.MemberID)
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

const getReactionsByPostID = `-- name: GetReactionsByPostID :many
SELECT reaction_id, post_id, member_id, reaction, created_at FROM post_reactions
WHERE post_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetReactionsByPostIDParams struct {
	PostID int32 `json:"post_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetReactionsByPostID(ctx context.Context, arg GetReactionsByPostIDParams) ([]PostReaction, error) {
	rows, err := q.db.QueryContext(ctx, getReactionsByPostID, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PostReaction{}
	for rows.Next() {
		var i PostReaction
		if err := rows.Scan(
			&i.ReactionID,
			&i.PostID,
			&i.MemberID,
			&i.Reaction,
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

const getReactionsCountByPostID = `-- name: GetReactionsCountByPostID :one
SELECT COUNT(*) FROM post_reactions
WHERE post_id = $1
`

func (q *Queries) GetReactionsCountByPostID(ctx context.Context, postID int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, getReactionsCountByPostID, postID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getReactionsWithTypeByPostID = `-- name: GetReactionsWithTypeByPostID :many
SELECT reaction FROM post_reactions
WHERE post_id = $1
`

func (q *Queries) GetReactionsWithTypeByPostID(ctx context.Context, postID int32) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getReactionsWithTypeByPostID, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var reaction string
		if err := rows.Scan(&reaction); err != nil {
			return nil, err
		}
		items = append(items, reaction)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
