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

-- name: SearchGroupByGroupName :many
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
    tags.tag_name,
    (SELECT COUNT(*)
    FROM group_members A
    WHERE A.group_id = groups.group_id) AS total_member
FROM groups
JOIN tags ON groups.tag_id = tags.tag_id
WHERE groups.group_name ILIKE '%' || $1 || '%'
LIMIT $2 OFFSET $3;

-- name: AddGroupMember :one
INSERT INTO group_members (group_id, member_id, role)
VALUES ($1, $2, $3)
RETURNING *;

-- name: NumberMemberInGroup :one
SELECT COUNT(*)
FROM group_members
WHERE group_id = $1;

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
WHERE group_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetGroupRequestCount :one
SELECT COUNT(*) FROM group_requests
WHERE group_id = $1;

-- name: GetMemberPendingGroupRequests :many
SELECT * FROM group_requests
WHERE member_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

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
    tags.tag_name,
    (SELECT COUNT(*)
    FROM group_members A
    WHERE A.group_id = $1) AS total_member
FROM groups
JOIN tags ON groups.tag_id = tags.tag_id
WHERE group_id = $1;

-- name: GetMembersByGroupID :many
SELECT *
FROM group_members
WHERE group_id = $1
LIMIT $2 OFFSET $3;

-- name: ListGroups :many
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
    tags.tag_name,
    (SELECT COUNT(*)
    FROM group_members S
    WHERE S.group_id = groups.group_id) AS total_member
FROM groups
JOIN tags ON groups.tag_id = tags.tag_id
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
        ('Learning and Development'),
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


-- name: DeleteGroupMembers :exec
DELETE FROM group_members WHERE member_id = $1;

-- name: DeleteGroupRequests :exec
DELETE FROM group_requests WHERE member_id = $1;

-- name: DeletePosts :exec
DELETE FROM posts WHERE member_id = $1;

-- name: DeletePostReactions :exec
DELETE FROM post_reactions WHERE member_id = $1;

-- name: DeletePostComments :exec
DELETE FROM post_comments WHERE member_id = $1;

-- name: GetAcceptorGroupRequests :many
SELECT g.group_id, g.group_name, g.photo_url,
       COUNT(*) AS request_count
FROM group_requests gr 
JOIN group_members gm ON gr.group_id = gm.group_id
JOIN groups g ON gr.group_id = g.group_id
WHERE gm.member_id = $1 AND gm.role != 'member'
GROUP BY g.group_id, g.group_name, g.photo_url;

-- name: GetAcceptorGroupRequestsCount :one
SELECT gm.member_id,
       COUNT(*) AS request_count
FROM group_requests gr 
JOIN group_members gm ON gr.group_id = gm.group_id
WHERE gm.member_id = $1 AND gm.role != 'member'
GROUP BY gm.member_id;

-- name: GetUserGroups :many
SELECT g.group_id, g.group_name, g.photo_url, g.group_type
FROM group_members gm
JOIN groups g ON g.group_id = gm.group_id
WHERE gm.member_id = $1;

-- name: CheckMemberInGroup :one
SELECT * FROM group_members
WHERE group_id = $1 AND member_id = $2;

-- name: GetRequestFromGroup :many
SELECT * FROM group_requests
WHERE group_id = $1 AND member_id = $2;

-- name: GetCategoryIDByName :one
SELECT category_id
FROM tag_category
WHERE category_name = $1;