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
    member_id,
    reaction
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: CreateComment :one
INSERT INTO post_comments (
    post_id,
    member_id,
    comment
) VALUES (
    $1, $2, $3
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

-- name: GetPostsForOngoingFeedByMemberID :many
SELECT DISTINCT posts.*, groups.group_name, groups.photo_url AS group_photo_url FROM posts
JOIN group_members ON posts.group_id = group_members.group_id JOIN groups ON posts.group_id = groups.group_id
WHERE group_members.member_id = $1
ORDER BY posts.created_at DESC;

-- name: GetPostsForFollowingFeedByMemberId :many
SELECT posts.*, groups.group_name, groups.photo_url AS group_photo_url FROM posts
JOIN groups ON posts.group_id = groups.group_id
WHERE posts.member_id = ANY($1::int[]) AND visibility = true
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
UPDATE posts SET description = $2
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
WHERE post_id = $1 AND member_id = $2 AND reaction = $3;

-- name: DeleteComment :exec
DELETE FROM post_comments
WHERE comment_id = $1;

-- name: GetPostLatestId :one
SELECT COALESCE(MAX(post_id), 0)::integer FROM posts;


-- name: GetCommentByCommentId :one
SELECT *
FROM post_comments
WHERE comment_id = $1;