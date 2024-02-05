-- name: CreateFriend :one
INSERT INTO friends (
  user_id1,
  user_id2
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetFriend :one
SELECT * FROM friends
WHERE user_id1 = $1 AND user_id2= $2
LIMIT 1;

-- name: GetFriendForUpdate :one
SELECT * FROM friends
WHERE user_id1 = $1 AND user_id2 = $2 LIMIT 1 
FOR NO KEY UPDATE;

-- name: ListFriendsByUserId :many
SELECT * FROM friends
WHERE user_id1 = $1
ORDER BY id
LIMIT $2
OFFSET $3;