-- name: CreateStreakSet :one
INSERT INTO streak_set (
    group_id,
    member_id,
    end_date
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetStreakSetByGroupId :many
SELECT * FROM streak_set
WHERE group_id = $1;

-- name: GetStreakSetByGroupIdandUserId :many
SELECT * FROM streak_set
WHERE group_id = $1 AND member_id = $2;

-- name: GetStreakSetByStreakSetId :one
SELECT * FROM streak_set
WHERE streak_set_id = $1;

-- name: GetStreakSetByEndDate :many
SELECT * FROM streak_set
WHERE end_date = $1;

-- name: EditStreakSetEndDate :exec
UPDATE streak_set
SET end_date = $3
WHERE group_id = $1 AND member_id = $2
RETURNING *;

-- name: DeleteStreakSet :exec
DELETE FROM streak_set WHERE group_id = $1 AND member_id = $2;

-- name: GetStreakByStreakSetId :one
SELECT * FROM streaks
WHERE streak_set_id = $1;

-- name: GetStreakByMemberIDandGroupID :one
SELECT s.*, ss.group_id, ss.member_id, ss.end_date
FROM streaks s
JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
WHERE ss.member_id = $1
AND ss.group_id = $2;

-- name: GetIncompletedStreakByUserID :many
SELECT s.*, ss.group_id, ss.member_id, ss.end_date
FROM streaks s
JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
WHERE ss.member_id = $1
AND s.completed = false;

-- name: GetStreakByMemberId :many
SELECT s.*, ss.group_id, ss.member_id, ss.end_date
FROM streaks s 
JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
WHERE ss.member_id = $1;

-- name: CreateStreak :one
INSERT INTO streaks (
    streak_set_id
) VALUES (
    $1
) RETURNING *;

-- name: IncreaseStreak :one
UPDATE streaks
SET total_streak_count = total_streak_count + 1, 
    weekly_streak_count = weekly_streak_count + 1,
    max_streak_count = max_streak_count +1,
    recent_date_added = CURRENT_TIMESTAMP
WHERE streak_set_id IN (
    SELECT s.streak_set_id
    FROM streaks s
    JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
    WHERE ss.member_id = $1
    AND ss.group_id = $2
) RETURNING *;

-- name: ResetStreak :exec
UPDATE streaks
SET total_streak_count = 0, 
    weekly_streak_count = 0,
    completed = false
WHERE streak_set_id IN (
    SELECT s.streak_set_id
    FROM streaks s
    JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
    WHERE ss.member_id = $1
    AND ss.group_id = $2
) RETURNING *;

-- -- name: ResetWeeklyStreak :exec
-- UPDATE streaks
-- SET weekly_streak_count = 0,
-- completed = false
-- WHERE streak_set_id IN (
--     SELECT s.streak_set_id
--     FROM streaks s
--     JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
--     WHERE ss.member_id = $1
--     AND ss.group_id = $2
-- ) RETURNING *;

-- -- name: ResetTotalStreak :exec
-- UPDATE streaks
-- SET total_streak_count = 0
-- WHERE streak_set_id IN (
--     SELECT s.streak_set_id
--     FROM streaks s
--     JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
--     WHERE ss.member_id = $1
--     AND ss.group_id = $2
-- ) RETURNING *;

-- name: EditCompleteStatus :exec
UPDATE streaks
SET completed = $3
WHERE streak_set_id IN (
    SELECT s.streak_set_id
    FROM streaks s
    JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
    WHERE ss.member_id = $1
    AND ss.group_id = $2
);

-- name: ResetTotalStreak :exec
UPDATE streaks
SET total_streak_count = 0,
    weekly_streak_count = 0
WHERE streak_set_id IN (
    SELECT s.streak_set_id
    FROM streaks s
    JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
    WHERE s.completed = false 
    AND ss.end_date::date = CURRENT_DATE
);

-- name: ResetWeeklyStreak :exec
UPDATE streaks
SET weekly_streak_count = 0,
completed = false
WHERE streak_set_id IN (
    SELECT s.streak_set_id
    FROM streaks s
    JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
    WHERE s.completed = true
    AND ss.end_date::date = CURRENT_DATE
) RETURNING *;

-- name: GetMaxStreakUser :one
SELECT max_streak_count
FROM streaks
WHERE max_streak_count = (
    SELECT MAX(max_streak_count)
    FROM streaks
    WHERE streak_set_id IN (
        SELECT s.streak_set_id
        FROM streaks s
        JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
        WHERE ss.member_id = $1
    )
);
