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
	CreateStreak(ctx context.Context, arg CreateStreakParams) (Streak, error)
	CreateStreakSet(ctx context.Context, arg CreateStreakSetParams) (StreakSet, error)
	DeleteComment(ctx context.Context, commentID int32) error
	DeleteGroup(ctx context.Context, groupID int32) error
	DeleteMember(ctx context.Context, arg DeleteMemberParams) error
	DeletePost(ctx context.Context, postID int32) error
	DeleteReaction(ctx context.Context, reactionID int32) error
	DeleteRequestToJoinGroup(ctx context.Context, arg DeleteRequestToJoinGroupParams) error
	EditComment(ctx context.Context, arg EditCommentParams) error
	EditGroupDescription(ctx context.Context, arg EditGroupDescriptionParams) error
	EditGroupName(ctx context.Context, arg EditGroupNameParams) error
	EditGroupPhoto(ctx context.Context, arg EditGroupPhotoParams) error
	EditGroupVisibility(ctx context.Context, arg EditGroupVisibilityParams) error
	EditMemberRole(ctx context.Context, arg EditMemberRoleParams) error
	EditPost(ctx context.Context, arg EditPostParams) error
	EditReaction(ctx context.Context, arg EditReactionParams) error
	EndStreakSet(ctx context.Context, streakSetID int32) error
	GetAvailableTags(ctx context.Context) ([]Tag, error)
	GetCommentByCommentId(ctx context.Context, commentID int32) (PostComment, error)
	GetCommentsByPostID(ctx context.Context, postID int32) ([]PostComment, error)
	GetCommentsCountByPostID(ctx context.Context, postID int32) (int64, error)
	GetGroupByID(ctx context.Context, groupID int32) (GetGroupByIDRow, error)
	GetGroupLatestId(ctx context.Context) (int32, error)
	GetGroupRequest(ctx context.Context, arg GetGroupRequestParams) (GroupRequest, error)
	GetGroupRequests(ctx context.Context, groupID int32) ([]GroupRequest, error)
	GetGroupsByTagID(ctx context.Context, tagID int32) ([]Group, error)
	GetLatestStreakSetByGroupIDAndUserID(ctx context.Context, arg GetLatestStreakSetByGroupIDAndUserIDParams) (StreakSet, error)
	GetMemberInfo(ctx context.Context, arg GetMemberInfoParams) (GetMemberInfoRow, error)
	GetMemberPendingGroupRequests(ctx context.Context, memberID int32) ([]GroupRequest, error)
	GetMembersByGroupID(ctx context.Context, groupID int32) ([]GroupMember, error)
	GetPostByPostID(ctx context.Context, postID int32) (Post, error)
	GetPostLatestId(ctx context.Context) (int32, error)
	GetPostsByGroupAndMemberID(ctx context.Context, arg GetPostsByGroupAndMemberIDParams) ([]Post, error)
	GetPostsByGroupID(ctx context.Context, groupID int32) ([]Post, error)
	GetPostsByMemberID(ctx context.Context, memberID int32) ([]Post, error)
	GetPostsForOngoingFeedByMemberID(ctx context.Context, memberID int32) ([]GetPostsForOngoingFeedByMemberIDRow, error)
	GetReactionById(ctx context.Context, reactionID int32) (PostReaction, error)
	GetReactionsByPostID(ctx context.Context, postID int32) ([]PostReaction, error)
	GetReactionsCountByPostID(ctx context.Context, postID int32) (int64, error)
	GetStreakByPostID(ctx context.Context, postID int32) (Streak, error)
	GetStreakSetByUserID(ctx context.Context, userID int32) ([]StreakSet, error)
	GetStreaksByGroupIDAndUserID(ctx context.Context, arg GetStreaksByGroupIDAndUserIDParams) ([]GetStreaksByGroupIDAndUserIDRow, error)
	GetStreaksByStreakSetID(ctx context.Context, streakSetID int32) ([]Streak, error)
	GetUnendedStreakSetByUserID(ctx context.Context, userID int32) ([]StreakSet, error)
	ListGroups(ctx context.Context, arg ListGroupsParams) ([]Group, error)
	SendRequestToJoinGroup(ctx context.Context, arg SendRequestToJoinGroupParams) (GroupRequest, error)
	UpdateStreakSetCount(ctx context.Context, arg UpdateStreakSetCountParams) error
}

var _ Querier = (*Queries)(nil)
