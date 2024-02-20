-- name: CreatePost :one
INSERT INTO posts (
    member_id,
    group_id,
    photo_url,
    description
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: CreateReaction :one
INSERT INTO post_reactions (
    post_id,
    reaction
) VALUES (
    $1, $2
) RETURNING *;

-- name: CreateComment :one
INSERT INTO post_comments (
    post_id,
    comment
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetPostByPostID :one
SELECT * FROM posts
WHERE post_id = $1;

-- name: GetPostsByGroupID :many
SELECT * FROM posts
WHERE group_id = $1
ORDER BY created_at DESC;

-- name: GetPostsByGroupAndMemberID :many
SELECT * FROM posts
WHERE group_id = $1 AND member_id = $2
ORDER BY created_at DESC;

-- name: GetPostsForFeedByMemberID :many
SELECT posts.* FROM posts
JOIN group_members ON posts.group_id = group_members.group_id
WHERE group_members.member_id = $1
ORDER BY posts.created_at DESC;

-- name: GetPostsByMemberID :many
SELECT * FROM posts
WHERE member_id = $1
ORDER BY created_at DESC;

-- name: GetReactionsByPostID :many
SELECT * FROM post_reactions
WHERE post_id = $1
ORDER BY created_at DESC;

-- name: GetReactionsCountByPostID :one
SELECT COUNT(*) FROM post_reactions
WHERE post_id = $1;

-- name: GetCommentsByPostID :many
SELECT * FROM post_comments
WHERE post_id = $1
ORDER BY created_at DESC;

-- name: GetCommentsCountByPostID :one
SELECT COUNT(*) FROM post_comments
WHERE post_id = $1;

-- name: EditPost :exec
UPDATE posts SET description = $2 AND photo_url = $3
WHERE post_id = $1;

-- name: EditReaction :exec
UPDATE post_reactions SET reaction = $2
WHERE reaction_id = $1;

-- name: EditComment :exec
UPDATE post_comments SET comment = $2
WHERE comment_id = $1;

-- name: DeletePost :exec
DELETE FROM posts
WHERE post_id = $1;

-- name: DeleteReaction :exec
DELETE FROM post_reactions
WHERE reaction_id = $1;

-- name: DeleteComment :exec
DELETE FROM post_comments
WHERE comment_id = $1;