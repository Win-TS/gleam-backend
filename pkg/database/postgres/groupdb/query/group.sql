-- name: CreateGroup :one
INSERT INTO groups (
    group_name,
    group_creator_id,
    photo_url
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: AddGroupMember :one
INSERT INTO group_members (
    group_id,
    member_id,
    role
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetGroupAndMemberByID :many
SELECT groups.group_id,
    groups.group_name,
    groups.photo_url,
    groups.created_at,
    group_members.member_id,
    group_members.role,
    group_members.created_at
FROM groups JOIN group_members ON groups.group_id = group_members.group_id
WHERE groups.group_id = $1;

-- name: GetGroupByID :one
SELECT * FROM groups
WHERE group_id = $1;

-- name: GetMembersByGroupID :many
SELECT * FROM group_members
WHERE group_id = $1;

-- name: ListGroups :many
SELECT * FROM groups
ORDER BY group_id
LIMIT $1
OFFSET $2;

-- name: EditGroupName :exec
UPDATE groups SET group_name = $2
WHERE group_id = $1;

-- name: EditGroupPhoto :exec
UPDATE groups SET photo_url = $2
WHERE group_id = $1;

-- name: EditMemberRole :exec
UPDATE group_members SET role = $3
WHERE group_id = $1 AND member_id = $2;

-- name: DeleteGroup :exec
DELETE FROM groups
WHERE group_id = $1;

-- name: DeleteMember :exec
DELETE FROM group_members 
WHERE member_id = $1 AND group_id = $2;

-- name: GetGroupLatestId :one
SELECT COALESCE(MAX(group_id), 0)::integer FROM groups;

-- name: CreateNewTag :one
INSERT INTO tags (
    tag_name,
    icon_url
) VALUES (
    $1, $2
) RETURNING *;

-- name: AddGroupTag :one
INSERT INTO group_tags (
    group_id,
    tag_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: AddMultipleTagsToGroup :many
INSERT INTO group_tags (
    group_id,
    tag_id
) VALUES (
    $1, unnest($2::integer[])
) RETURNING *;

-- name: GetAvailableTags :many
SELECT * FROM tags
ORDER BY tag_id;

-- name: GetGroupsByTagName :many
SELECT groups.* FROM groups
JOIN group_tags ON groups.group_id = group_tags.group_id
JOIN tags ON group_tags.tag_id = tags.tag_id
WHERE tags.tag_name = $1;

-- name: GetTagsByGroupID :many
SELECT tags.* FROM tags
JOIN group_tags ON tags.tag_id = group_tags.tag_id
WHERE group_tags.group_id = $1;