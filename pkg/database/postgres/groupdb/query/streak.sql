--name: CreateStreakSet :one
INSERT INTO streak_set (
    group_id,
    user_id,
    streak_count
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetStreakSetByUserID :many
SELECT * FROM streak_set
WHERE user_id = $1;

-- name: GetLatestStreakSetByGroupIDAndUserID :one
SELECT * FROM streak_set
WHERE group_id = $1 AND user_id = $2
ORDER BY created_at DESC 
LIMIT 1;

-- name: GetUnendedStreakSetByUserID :many
SELECT * FROM streak_set
WHERE user_id = $1 AND ended = false;

-- name: GetStreaksByStreakSetID :many
SELECT * FROM streaks
WHERE streak_set_id = $1;

-- name: GetStreaksByGroupIDAndUserID :many
SELECT * FROM streak_set
JOIN streaks ON streak_set.streak_set_id = streaks.streak_set_id
WHERE group_id = $1 AND user_id = $2
ORDER BY streak_set.created_at DESC;

-- name: UpdateStreakSetCount :exec
UPDATE streak_set SET streak_count = $2
WHERE streak_set_id = $1;

-- name: EndStreakSet :exec
UPDATE streak_set SET ended = true
WHERE streak_set_id = $1;

-- name: CreateStreak :one
INSERT INTO streaks (
    streak_set_id,
    post_id,
    streak_count
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetStreakByPostID :one
SELECT * FROM streaks
WHERE post_id = $1;