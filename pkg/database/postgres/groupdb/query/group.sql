-- name: CreateGroup :one
INSERT INTO groups (
        group_name,
        group_creator_id,
        photo_url,
        frequency,
        max_members,
        group_type,
        description,
        visibility,
        tag_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: AddGroupMember :one
INSERT INTO group_members (group_id, member_id, role)
VALUES ($1, $2, $3)
RETURNING *;

-- name: SendRequestToJoinGroup :one
INSERT INTO group_requests (group_id, member_id, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteRequestToJoinGroup :exec
DELETE FROM group_requests
WHERE group_id = $1
    AND member_id = $2;

-- name: GetGroupRequest :one
SELECT * FROM group_requests
WHERE group_id = $1
    AND member_id = $2;

-- name: GetGroupRequests :many
SELECT * FROM group_requests
WHERE group_id = $1;

-- name: GetMemberPendingGroupRequests :many
SELECT * FROM group_requests
WHERE member_id = $1;

-- name: GetGroupByID :one
SELECT groups.group_id,
    groups.group_name,
    groups.photo_url,
    groups.group_creator_id,
    groups.frequency,
    groups.max_members,
    groups.group_type,
    groups.visibility,
    groups.description,
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

-- name: EditGroupVisibility :exec
UPDATE groups
SET visibility = $2
WHERE group_id = $1;

-- name: EditGroupDescription :exec
UPDATE groups
SET description = $2
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

-- name: InitializeCategory :exec
INSERT INTO tag_category (category_name)
VALUES ('Sports and Fitness'), 
        ('Learning and development'),
		('Health and Wellness'), 
        ('Entertainment and Media'), 
        ('Hobbies and Leisure'),
        ('Others');

-- name: CreateNewTag :one
INSERT INTO tags (tag_name, icon_url, category_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAvailableCategory :many
SELECT *
FROM tag_category
ORDER BY category_id;

-- name: GetTagByCategory :many
SELECT *
FROM tags
WHERE category_id = $1;

-- name: GetTagByGroupId :one
SELECT groups.group_id, groups.group_name, tags.tag_id, tags.tag_name, tags.icon_url, tags.category_id, tag_category.category_name
FROM groups
JOIN tags ON groups.tag_id = tags.tag_id
JOIN tag_category ON tags.category_id = tag_category.category_id
WHERE groups.group_id = $1;


-- name: GetGroupsByCategoryID :many
SELECT groups.*, tags.category_id, tag_category.category_name
FROM groups
JOIN tags ON groups.tag_id = tags.tag_id
JOIN tag_category ON tags.category_id = tag_category.category_id
WHERE tags.category_id = $1;

-- name: EditTagName :exec
UPDATE tags
SET tag_name = $2
WHERE tag_id = $1;

-- name: EditTagCategory :exec
UPDATE tags
SET category_id = $2
WHERE tag_id = $1;

-- name: EditTagIcon :exec
UPDATE tags
SET icon_url = $2
WHERE tag_id = $1;

-- name: GetTagByTagID :one
SELECT *
FROM tags
WHERE tag_id = $1;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE tag_id = $1;

-- name: EditGroupTag :one
UPDATE groups 
SET tag_id = $2
WHERE group_id = $1
RETURNING *;