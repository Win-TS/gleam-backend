-- name: CreateGroup :one
INSERT INTO groups (
        group_name,
        group_creator_id,
        photo_url,
        frequency,
        tag_id
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: AddGroupMember :one
INSERT INTO group_members (group_id, member_id, role)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetGroupByID :one
SELECT groups.group_id,
    groups.group_name,
    groups.photo_url,
    groups.group_creator_id,
    groups.frequency,
    groups.created_at,
    tags.tag_name
FROM groups
    JOIN tags ON groups.tag_id = tags.tag_id
WHERE group_id = $1;

-- name: GetMembersByGroupID :many
SELECT *
FROM group_members
WHERE group_id = $1;

-- name: ListGroups :many
SELECT *
FROM groups
ORDER BY group_id
LIMIT $1 OFFSET $2;

-- name: EditGroupName :exec
UPDATE groups
SET group_name = $2
WHERE group_id = $1;

-- name: EditGroupPhoto :exec
UPDATE groups
SET photo_url = $2
WHERE group_id = $1;

-- name: EditMemberRole :exec
UPDATE group_members
SET role = $3
WHERE group_id = $1
    AND member_id = $2;

-- name: DeleteGroup :exec
DELETE FROM groups
WHERE group_id = $1;

-- name: DeleteMember :exec
DELETE FROM group_members
WHERE member_id = $1
    AND group_id = $2;

-- name: GetGroupLatestId :one
SELECT COALESCE(MAX(group_id), 0)::integer
FROM groups;

-- name: CreateNewTag :one
INSERT INTO tags (tag_name, icon_url)
VALUES ($1, $2)
RETURNING *;

-- name: GetAvailableTags :many
SELECT *
FROM tags
ORDER BY tag_id;

-- name: GetGroupsByTagID :many
SELECT *
from groups
WHERE tag_id = $1;

-- name: GetMemberInfo :one
SELECT gm.*, g.group_name
FROM group_members gm
JOIN groups g ON gm.group_id = g.group_id
WHERE gm.member_id = $1
AND gm.group_id = $2;

-- name: GetReactionById :one
SELECT *
FROM post_reactions
WHERE reaction_id = $1;