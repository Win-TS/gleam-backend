// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package groupdb

import (
	"context"
)

type Querier interface {
	AddGroupMember(ctx context.Context, arg AddGroupMemberParams) (GroupMember, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (PostComment, error)
	CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error)
	CreateNewTag(ctx context.Context, arg CreateNewTagParams) (Tag, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreateReaction(ctx context.Context, arg CreateReactionParams) (PostReaction, error)
	DeleteComment(ctx context.Context, commentID int32) error
	DeleteGroup(ctx context.Context, groupID int32) error
	DeleteMember(ctx context.Context, arg DeleteMemberParams) error
	DeletePost(ctx context.Context, postID int32) error
	DeleteReaction(ctx context.Context, reactionID int32) error
	EditComment(ctx context.Context, arg EditCommentParams) error
	EditGroupName(ctx context.Context, arg EditGroupNameParams) error
	EditGroupPhoto(ctx context.Context, arg EditGroupPhotoParams) error
	EditMemberRole(ctx context.Context, arg EditMemberRoleParams) error
	EditPost(ctx context.Context, arg EditPostParams) error
	EditReaction(ctx context.Context, arg EditReactionParams) error
	GetAvailableTags(ctx context.Context) ([]Tag, error)
	GetCommentsByPostID(ctx context.Context, postID int32) ([]PostComment, error)
	GetCommentsCountByPostID(ctx context.Context, postID int32) (int64, error)
	GetGroupByID(ctx context.Context, groupID int32) (GetGroupByIDRow, error)
	GetGroupLatestId(ctx context.Context) (int32, error)
	GetGroupsByTagID(ctx context.Context, tagID int32) ([]Group, error)
	GetMembersByGroupID(ctx context.Context, groupID int32) ([]GroupMember, error)
	GetPostByPostID(ctx context.Context, postID int32) (Post, error)
	GetPostLatestId(ctx context.Context) (int32, error)
	GetPostsByGroupAndMemberID(ctx context.Context, arg GetPostsByGroupAndMemberIDParams) ([]Post, error)
	GetPostsByGroupID(ctx context.Context, groupID int32) ([]Post, error)
	GetPostsByMemberID(ctx context.Context, memberID int32) ([]Post, error)
	GetPostsForFeedByMemberID(ctx context.Context, memberID int32) ([]Post, error)
	GetReactionsByPostID(ctx context.Context, postID int32) ([]PostReaction, error)
	GetReactionsCountByPostID(ctx context.Context, postID int32) (int64, error)
	ListGroups(ctx context.Context, arg ListGroupsParams) ([]Group, error)
}

var _ Querier = (*Queries)(nil)
