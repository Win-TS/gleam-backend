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

-- name: GetFriendsCountByID :one
SELECT COUNT(*) FROM friends
WHERE (user_id1 = $1 OR user_id2 = $1) AND status = 'Accepted';

-- name: GetFriendsRequestedList :many
SELECT users.* FROM friends JOIN users ON friends.user_id2 = users.id
WHERE user_id1 = $1 AND status = 'Pending'
ORDER BY friends.created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetFriendsPendingList :many
SELECT users.* FROM friends JOIN users ON friends.user_id1 = users.id
WHERE user_id2 = $1 AND status = 'Pending'
ORDER BY friends.created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListFriendsByUserId :many
SELECT 
    CASE
        WHEN user_id1 = $1 THEN user_id2
        ELSE user_id1
    END AS friend_id,
    users.*
FROM 
    friends 
JOIN 
    users ON (CASE
                WHEN user_id1 = $1 THEN user_id2
                ELSE user_id1
            END) = users.id
WHERE 
    (user_id1 = $1 OR user_id2 = $1)
    AND status = 'Accepted'
ORDER BY 
    friend_id
LIMIT $2 OFFSET $3;

-- name: ListFriendsByUserIdNoPaginate :many
SELECT 
    CASE
        WHEN user_id1 = $1 THEN user_id2
        ELSE user_id1
    END AS friend_id,
    users.*
FROM 
    friends 
JOIN 
    users ON (CASE
                WHEN user_id1 = $1 THEN user_id2
                ELSE user_id1
            END) = users.id
WHERE 
    (user_id1 = $1 OR user_id2 = $1)
    AND status = 'Accepted'
ORDER BY 
    friend_id;

-- name: EditFriendStatusAccepted :exec
UPDATE friends SET status = 'Accepted'
WHERE user_id1 = $1 AND user_id2 = $2;


-- name: EditFriendStatusDeclined :exec
DELETE FROM friends
WHERE user_id1 = $1 AND user_id2 = $2;