// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package groupdb

import (
	"context"
	"database/sql"
	"time"
)

type Querier interface {
	AddGroupMember(ctx context.Context, arg AddGroupMemberParams) (GroupMember, error)
	CheckMemberInGroup(ctx context.Context, arg CheckMemberInGroupParams) (GroupMember, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (PostComment, error)
	CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error)
	CreateNewTag(ctx context.Context, arg CreateNewTagParams) (Tag, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreateReaction(ctx context.Context, arg CreateReactionParams) (PostReaction, error)
	CreateStreak(ctx context.Context, streakSetID int32) (Streak, error)
	CreateStreakSet(ctx context.Context, arg CreateStreakSetParams) (StreakSet, error)
	DeleteComment(ctx context.Context, commentID int32) error
	DeleteGroup(ctx context.Context, groupID int32) error
	DeleteGroupMembers(ctx context.Context, memberID int32) error
	DeleteGroupRequests(ctx context.Context, memberID int32) error
	DeleteMember(ctx context.Context, arg DeleteMemberParams) error
	DeletePost(ctx context.Context, postID int32) error
	DeletePostComments(ctx context.Context, memberID int32) error
	DeletePostReactions(ctx context.Context, memberID int32) error
	DeletePosts(ctx context.Context, memberID int32) error
	DeleteReaction(ctx context.Context, arg DeleteReactionParams) error
	DeleteRequestToJoinGroup(ctx context.Context, arg DeleteRequestToJoinGroupParams) error
	DeleteStreakSet(ctx context.Context, arg DeleteStreakSetParams) error
	DeleteTag(ctx context.Context, tagID int32) error
	EditComment(ctx context.Context, arg EditCommentParams) error
	// -- name: ResetWeeklyStreak :exec
	// UPDATE streaks
	// SET weekly_streak_count = 0,
	// completed = false
	// WHERE streak_set_id IN (
	//     SELECT s.streak_set_id
	//     FROM streaks s
	//     JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
	//     WHERE ss.member_id = $1
	//     AND ss.group_id = $2
	// ) RETURNING *;
	// -- name: ResetTotalStreak :exec
	// UPDATE streaks
	// SET total_streak_count = 0
	// WHERE streak_set_id IN (
	//     SELECT s.streak_set_id
	//     FROM streaks s
	//     JOIN streak_set ss ON s.streak_set_id = ss.streak_set_id
	//     WHERE ss.member_id = $1
	//     AND ss.group_id = $2
	// ) RETURNING *;
	EditCompleteStatus(ctx context.Context, arg EditCompleteStatusParams) error
	EditGroupDescription(ctx context.Context, arg EditGroupDescriptionParams) error
	EditGroupName(ctx context.Context, arg EditGroupNameParams) error
	EditGroupPhoto(ctx context.Context, arg EditGroupPhotoParams) error
	EditGroupTag(ctx context.Context, arg EditGroupTagParams) (Group, error)
	EditGroupVisibility(ctx context.Context, arg EditGroupVisibilityParams) error
	EditMemberRole(ctx context.Context, arg EditMemberRoleParams) error
	EditPost(ctx context.Context, arg EditPostParams) error
	EditReaction(ctx context.Context, arg EditReactionParams) error
	EditStreakSetEndDate(ctx context.Context, arg EditStreakSetEndDateParams) error
	EditTagCategory(ctx context.Context, arg EditTagCategoryParams) error
	EditTagIcon(ctx context.Context, arg EditTagIconParams) error
	EditTagName(ctx context.Context, arg EditTagNameParams) error
	GetAcceptorGroupRequests(ctx context.Context, memberID int32) ([]GetAcceptorGroupRequestsRow, error)
	GetAcceptorGroupRequestsCount(ctx context.Context, memberID int32) (GetAcceptorGroupRequestsCountRow, error)
	GetAvailableCategory(ctx context.Context) ([]TagCategory, error)
	GetAvailableTags(ctx context.Context) ([]Tag, error)
	GetCommentByCommentId(ctx context.Context, commentID int32) (PostComment, error)
	GetCommentsByPostID(ctx context.Context, arg GetCommentsByPostIDParams) ([]PostComment, error)
	GetCommentsCountByPostID(ctx context.Context, postID int32) (int64, error)
	GetGroupByID(ctx context.Context, groupID int32) (GetGroupByIDRow, error)
	GetGroupLatestId(ctx context.Context) (int32, error)
	GetGroupRequest(ctx context.Context, arg GetGroupRequestParams) (GroupRequest, error)
	GetGroupRequestCount(ctx context.Context, groupID int32) (int64, error)
	GetGroupRequests(ctx context.Context, arg GetGroupRequestsParams) ([]GroupRequest, error)
	GetGroupsByCategoryID(ctx context.Context, categoryID sql.NullInt32) ([]GetGroupsByCategoryIDRow, error)
	GetGroupsByTagID(ctx context.Context, tagID int32) ([]Group, error)
	GetIncompletedStreakByUserID(ctx context.Context, memberID int32) ([]GetIncompletedStreakByUserIDRow, error)
	GetMaxStreakUser(ctx context.Context, memberID int32) (int32, error)
	GetMemberInfo(ctx context.Context, arg GetMemberInfoParams) (GetMemberInfoRow, error)
	GetMemberPendingGroupRequests(ctx context.Context, arg GetMemberPendingGroupRequestsParams) ([]GroupRequest, error)
	GetMembersByGroupID(ctx context.Context, arg GetMembersByGroupIDParams) ([]GroupMember, error)
	GetPostByPostID(ctx context.Context, postID int32) (Post, error)
	GetPostLatestId(ctx context.Context) (int32, error)
	GetPostsByGroupAndMemberID(ctx context.Context, arg GetPostsByGroupAndMemberIDParams) ([]Post, error)
	GetPostsByGroupID(ctx context.Context, arg GetPostsByGroupIDParams) ([]Post, error)
	GetPostsByMemberID(ctx context.Context, arg GetPostsByMemberIDParams) ([]Post, error)
	GetPostsForFollowingFeedByMemberId(ctx context.Context, arg GetPostsForFollowingFeedByMemberIdParams) ([]GetPostsForFollowingFeedByMemberIdRow, error)
	GetPostsForOngoingFeedByMemberID(ctx context.Context, arg GetPostsForOngoingFeedByMemberIDParams) ([]GetPostsForOngoingFeedByMemberIDRow, error)
	GetReactionById(ctx context.Context, reactionID int32) (PostReaction, error)
	GetReactionByPostIDAndUserID(ctx context.Context, arg GetReactionByPostIDAndUserIDParams) (PostReaction, error)
	GetReactionsByPostID(ctx context.Context, arg GetReactionsByPostIDParams) ([]PostReaction, error)
	GetReactionsCountByPostID(ctx context.Context, postID int32) (int64, error)
	GetReactionsWithTypeByPostID(ctx context.Context, postID int32) ([]string, error)
	GetRequestFromGroup(ctx context.Context, arg GetRequestFromGroupParams) ([]GroupRequest, error)
	GetStreakByMemberIDandGroupID(ctx context.Context, arg GetStreakByMemberIDandGroupIDParams) (GetStreakByMemberIDandGroupIDRow, error)
	GetStreakByMemberId(ctx context.Context, memberID int32) ([]GetStreakByMemberIdRow, error)
	GetStreakByStreakSetId(ctx context.Context, streakSetID int32) (Streak, error)
	GetStreakSetByEndDate(ctx context.Context, endDate time.Time) ([]StreakSet, error)
	GetStreakSetByGroupId(ctx context.Context, groupID int32) ([]StreakSet, error)
	GetStreakSetByGroupIdandUserId(ctx context.Context, arg GetStreakSetByGroupIdandUserIdParams) ([]StreakSet, error)
	GetStreakSetByStreakSetId(ctx context.Context, streakSetID int32) (StreakSet, error)
	GetTagByCategory(ctx context.Context, categoryID sql.NullInt32) ([]Tag, error)
	GetTagByGroupId(ctx context.Context, groupID int32) (GetTagByGroupIdRow, error)
	GetTagByTagID(ctx context.Context, tagID int32) (Tag, error)
	GetUserGroups(ctx context.Context, memberID int32) ([]GetUserGroupsRow, error)
	IncreaseStreak(ctx context.Context, arg IncreaseStreakParams) (Streak, error)
	InitializeCategory(ctx context.Context) error
	ListGroups(ctx context.Context, arg ListGroupsParams) ([]ListGroupsRow, error)
	NumberMemberInGroup(ctx context.Context, groupID int32) (int64, error)
	ResetStreak(ctx context.Context, arg ResetStreakParams) error
	ResetTotalStreak(ctx context.Context) error
	ResetWeeklyStreak(ctx context.Context) error
	SearchGroupByGroupName(ctx context.Context, arg SearchGroupByGroupNameParams) ([]SearchGroupByGroupNameRow, error)
	SendRequestToJoinGroup(ctx context.Context, arg SendRequestToJoinGroupParams) (GroupRequest, error)
}

var _ Querier = (*Queries)(nil)
