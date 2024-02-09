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

-- name: GetFriendsCountByID: one
SELECT COUNT(*) FROM friends
WHERE (user_id1 = $1 OR user_id2 = $1) AND status = 'Accepted';

-- name: GetFriendsListByID: many
SELECT
    CASE
        WHEN user_id1 = $1 THEN user_id2
        ELSE user_id1
    END AS friend_id
FROM friends
WHERE (user_id1 = $1 OR user_id2 = $1) AND status = 'Accepted';

-- name: GetFriendsPendingList: many
SELECT * FROM friends
WHERE user_id2 = $1 AND status = 'Pending';

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