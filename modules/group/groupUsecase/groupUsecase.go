package groupUsecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"

	"firebase.google.com/go/storage"
	//groupPb "github.com/Win-TS/gleam-backend.git/modules/group/groupPb"
	"github.com/Win-TS/gleam-backend.git/modules/group"
	userPb "github.com/Win-TS/gleam-backend.git/modules/user/userPb"
	groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
	"github.com/Win-TS/gleam-backend.git/pkg/grpcconn"
	"github.com/Win-TS/gleam-backend.git/pkg/utils"
	"github.com/jaswdr/faker"
)

type (
	GroupUsecaseService interface {
		SearchUser(pctx context.Context, grpcUrl string, req *userPb.SearchUserReq) (*userPb.SearchUserRes, error)
		GetRole(pctx context.Context, userID int32, groupID int32) (group.Role, error)
		CreateNewGroup(pctx context.Context, args groupdb.CreateGroupParams) (groupdb.Group, error)
		SendRequestToJoinGroup(pctx context.Context, args groupdb.SendRequestToJoinGroupParams) (groupdb.GroupRequest, error)
		AcceptGroupRequest(pctx context.Context, args groupdb.AcceptGroupRequestParams) (groupdb.GroupMember, error)
		DeclineGroupRequest(pctx context.Context, args groupdb.DeleteRequestToJoinGroupParams, declinerId int) error
		GetGroupJoinRequests(pctx context.Context, args groupdb.GetGroupRequestsParams, grpcUrl string) ([]group.GroupRequestRes, error)
		GetGroupJoinRequestCount(pctx context.Context, groupId int) (int, error)
		GetUserJoinRequests(pctx context.Context, args groupdb.GetMemberPendingGroupRequestsParams) ([]groupdb.GroupRequest, error)
		GetGroupById(pctx context.Context, groupId, userId int) (group.GetGroupByIdRes, error)
		GetGroupMembersByGroupId(pctx context.Context, args groupdb.GetMembersByGroupIDParams, grpcUrl string) ([]group.GroupMemberRes, error)
		ListGroups(pctx context.Context, args groupdb.ListGroupsParams) ([]groupdb.ListGroupsRow, error)
		EditGroupName(pctx context.Context, args groupdb.EditGroupNameParams, editorId int32) (groupdb.GetGroupByIDRow, error)
		EditGroupPhoto(pctx context.Context, args groupdb.EditGroupPhotoParams, editorId int32) (groupdb.GetGroupByIDRow, error)
		EditGroupVisibility(pctx context.Context, args groupdb.EditGroupVisibilityParams, editorId int32) (groupdb.GetGroupByIDRow, error)
		EditGroupDescription(pctx context.Context, args groupdb.EditGroupDescriptionParams, editorId int32) (groupdb.GetGroupByIDRow, error)
		EditMemberRole(pctx context.Context, args groupdb.EditMemberRoleParams, editorId int32) (groupdb.GetMemberInfoRow, error)
		DeleteGroup(pctx context.Context, groupId int, editorId int32) error
		DeleteGroupMember(pctx context.Context, args groupdb.DeleteMemberParams, editorId int32) error
		CreatePost(pctx context.Context, args groupdb.CreatePostParams) (groupdb.Post, error)
		GetPostByPostId(pctx context.Context, postId, userId int, grpcUrl string) (groupdb.Post, *userPb.GetUserProfileRes, *groupdb.PostReaction, error)
		GetPostsByGroupId(pctx context.Context, args groupdb.GetPostsByGroupIDParams, grpcUrl string) ([]group.PostByGroupRes, error)
		GetPostsByUserId(pctx context.Context, args groupdb.GetPostsByMemberIDParams) ([]groupdb.Post, error)
		GetPostsByGroupAndMemberId(pctx context.Context, args groupdb.GetPostsByGroupAndMemberIDParams) ([]groupdb.Post, error)
		EditPost(pctx context.Context, args groupdb.EditPostParams) (groupdb.Post, error)
		DeletePost(pctx context.Context, postId int) error
		GetPostsForOngoingFeedByMemberId(pctx context.Context, args groupdb.GetPostsForOngoingFeedByMemberIDParams, grpcUrl string) ([]group.PostsForFeedRes, error)
		CreateReaction(pctx context.Context, args groupdb.CreateReactionParams) (groupdb.PostReaction, error)
		GetReactionsByPostId(pctx context.Context, args groupdb.GetReactionsByPostIDParams, grpcUrl string) ([]group.ReactionPostRes, error)
		GetReactionsCountByPostId(pctx context.Context, postId int) (map[string]int, int, error)
		EditReaction(pctx context.Context, args groupdb.EditReactionParams) (groupdb.PostReaction, error)
		DeleteReaction(pctx context.Context, args groupdb.DeleteReactionParams) error
		CreateComment(pctx context.Context, args groupdb.CreateCommentParams) (groupdb.PostComment, error)
		GetCommentsByPostId(pctx context.Context, args groupdb.GetCommentsByPostIDParams, grpcUrl string) ([]group.CommentRes, error)
		GetCommentCountByPostId(pctx context.Context, postId int) (int, error)
		EditComment(pctx context.Context, args groupdb.EditCommentParams) (groupdb.PostComment, error)
		DeleteComment(pctx context.Context, commentId int) error
		SaveToFirebaseStorage(pctx context.Context, bucketName, objectPath, filename string, file io.Reader) (string, error)
		GetGroupLatestId(pctx context.Context) (int, error)
		GetPostLatestId(pctx context.Context) (int, error)
		CreateNewTag(pctx context.Context, args groupdb.CreateNewTagParams) (groupdb.Tag, error)
		GetAvailableTags(pctx context.Context) ([]groupdb.Tag, error)
		GetGroupsByTagID(pctx context.Context, tagId int) ([]groupdb.Group, error)
		GetTagByCategory(pctx context.Context, categoryID int32) ([]groupdb.Tag, error)
		GetTagByGroupId(pctx context.Context, groupId int32) (groupdb.GetTagByGroupIdRow, error)
		GetGroupsByCategoryID(pctx context.Context, categoryId int32) ([]groupdb.GetGroupsByCategoryIDRow, error)
		GetTagByTagID(pctx context.Context, tagId int32) (groupdb.Tag, error)
		EditTagName(pctx context.Context, args groupdb.EditTagNameParams) (groupdb.Tag, error)
		EditTagCategory(pctx context.Context, args groupdb.EditTagCategoryParams) (groupdb.Tag, error)
		EditTagIcon(pctx context.Context, args groupdb.EditTagIconParams) (groupdb.Tag, error)
		DeleteTag(pctx context.Context, tagId int) error
		EditGroupTag(pctx context.Context, args groupdb.EditGroupTagParams, memberId int32) (groupdb.Group, error)
		GroupMockData(pctx context.Context, count int) error
		PostMockData(ctx context.Context, count int) error
		GetBatchUserProfiles(pctx context.Context, grpcUrl string, ids []int32) (*userPb.GetBatchUserProfileRes, error)
		GetUserProfile(pctx context.Context, grpcUrl string, req *userPb.GetUserProfileReq) (*userPb.GetUserProfileRes, error)
		GetPostsForFollowingFeedByMemberId(pctx context.Context, userId, limit, offset int, grpcUrl string) ([]group.PostsForFeedRes, error)
		SearchGroupByGroupName(ctx context.Context, args groupdb.SearchGroupByGroupNameParams) ([]groupdb.SearchGroupByGroupNameRow, error)
		DeleteUserData(ctx context.Context, userID int32) error
		GetAcceptorGroupRequests(ctx context.Context, userId int32) ([]groupdb.GetAcceptorGroupRequestsRow, error)
		GetAcceptorGroupRequestsCount(ctx context.Context, userId int32) (groupdb.GetAcceptorGroupRequestsCountRow, error)
		GetUserGroups(ctx context.Context, userId int32) (group.GetUserGroupRes, error)
		GetStreakByMemberId(ctx context.Context, MemberId int32) ([]groupdb.GetStreakByMemberIdRow, error)
		GetStreakByMemberIDandGroupID(ctx context.Context, args groupdb.GetStreakByMemberIDandGroupIDParams) (groupdb.GetStreakByMemberIDandGroupIDRow, error)
		GetIncompletedStreakByUserID(ctx context.Context, memberId int32) ([]groupdb.GetIncompletedStreakByUserIDRow, error)
		GetMaxStreakByMemberId(ctx context.Context, MemberId int32) (int32, error)
		IncreaseStreak(ctx context.Context, args groupdb.GetStreakByMemberIDandGroupIDParams) (groupdb.Streak, error)
		MockupTag(ctx context.Context) error
		MockupGroup(ctx context.Context) error
		MockupMember(ctx context.Context) error
		MockupPost(ctx context.Context) error
		MockupComment(ctx context.Context) error
		MockupReactions(ctx context.Context) error
	}

	groupUsecase struct {
		store         groupdb.Store
		storageClient *storage.Client
	}
)

func NewGroupUsecase(store groupdb.Store, storageClient *storage.Client) GroupUsecaseService {
	return &groupUsecase{store, storageClient}
}

func (u *groupUsecase) SearchUser(pctx context.Context, grpcUrl string, req *userPb.SearchUserReq) (*userPb.SearchUserRes, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("error - gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.User().SearchUser(ctx, req)
	if err != nil {
		log.Printf("error - SearchUser failed: %s", err.Error())
		return nil, errors.New("error: userId not found")
	}

	return result, nil
}

func (u *groupUsecase) GetUserProfile(pctx context.Context, grpcUrl string, req *userPb.GetUserProfileReq) (*userPb.GetUserProfileRes, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("error - gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.User().GetUserProfile(ctx, req)
	if err != nil {
		log.Printf("error - GetUserProfile failed: %s", err.Error())
		return nil, errors.New("error: GetUserProfile failed")
	}

	return result, nil
}

func (u *groupUsecase) GetBatchUserProfiles(pctx context.Context, grpcUrl string, ids []int32) (*userPb.GetBatchUserProfileRes, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("error - gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.User().GetBatchUserProfiles(ctx, &userPb.GetBatchUserProfileReq{UserIds: ids})
	if err != nil {
		log.Printf("error - GetBatchUserProfiles failed: %s", err.Error())
		return nil, errors.New("error: Get'BatchUserProfiles failed")
	}

	return result, nil
}

func (u *groupUsecase) GetRole(pctx context.Context, userID int32, groupID int32) (group.Role, error) {
	arg := groupdb.GetMemberInfoParams{
		MemberID: userID,
		GroupID:  groupID,
	}
	memberInfo, err := u.store.GetMemberInfo(pctx, arg)
	if err != nil {
		return group.Role(""), err
	}
	return group.Role(memberInfo.Role), nil
}

func (u *groupUsecase) CreateNewGroup(pctx context.Context, args groupdb.CreateGroupParams) (groupdb.Group, error) {
	newGroup, err := u.store.CreateNewGroup(pctx, args)
	if err != nil {
		return groupdb.Group{}, err
	}
	_, err = u.CreateStreak(pctx, int32(newGroup.GroupID), int32(newGroup.GroupCreatorID))
	if err != nil {
		return groupdb.Group{}, err
	}
	return newGroup, nil
}

func (u *groupUsecase) SendRequestToJoinGroup(pctx context.Context, args groupdb.SendRequestToJoinGroupParams) (groupdb.GroupRequest, error) {
	_, err := u.store.GetMemberInfo(pctx, groupdb.GetMemberInfoParams{
		MemberID: args.MemberID,
		GroupID:  args.GroupID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			groupInfo, err := u.store.GetGroupByID(pctx, int32(args.GroupID))
			if err != nil {
				return groupdb.GroupRequest{}, err
			}

			numMember, err := u.store.NumberMemberInGroup(pctx, int32(args.GroupID))
			if err != nil || numMember > int64(groupInfo.MaxMembers) {
				return groupdb.GroupRequest{}, errors.New("group is full")
			}

			newRequest, err := u.store.SendRequestToJoinGroup(pctx, args)
			if err != nil {
				return groupdb.GroupRequest{}, err
			}
			return newRequest, nil
		}
		return groupdb.GroupRequest{}, err
	}
	return groupdb.GroupRequest{}, errors.New("member is already in the group")
}

func (u *groupUsecase) AcceptGroupRequest(pctx context.Context, args groupdb.AcceptGroupRequestParams) (groupdb.GroupMember, error) {
	role, err := u.GetRole(pctx, int32(args.AcceptorId), int32(args.GroupID))
	if err != nil {
		return groupdb.GroupMember{}, err
	}
	if role != group.Admin && role != group.Moderator {
		return groupdb.GroupMember{}, errors.New("no permission")
	}
	groupInfo, err := u.store.GetGroupByID(pctx, int32(args.GroupID))
	if err != nil {
		return groupdb.GroupMember{}, err
	}

	numMember, err := u.store.NumberMemberInGroup(pctx, int32(args.GroupID))
	if err != nil || numMember > int64(groupInfo.MaxMembers) {
		return groupdb.GroupMember{}, errors.New("group is full")

	}

	newMember, err := u.store.AcceptGroupRequest(pctx, args)
	if err != nil {
		return groupdb.GroupMember{}, err
	}
	u.CreateStreak(pctx, int32(args.GroupID), int32(args.MemberID))
	return newMember, nil
}

func (u *groupUsecase) DeclineGroupRequest(pctx context.Context, args groupdb.DeleteRequestToJoinGroupParams, declinerId int) error {
	role, err := u.GetRole(pctx, int32(declinerId), int32(args.GroupID))
	if err != nil {
		return err
	}
	if role != group.Admin && role != group.Moderator {
		return errors.New("no permission")
	}

	if err := u.store.DeleteRequestToJoinGroup(pctx, args); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) GetGroupJoinRequests(pctx context.Context, args groupdb.GetGroupRequestsParams, grpcUrl string) ([]group.GroupRequestRes, error) {
	requests, err := u.store.GetGroupRequests(pctx, args)
	if err != nil {
		return []group.GroupRequestRes{}, err
	}

	var memberIds []int32
	for _, request := range requests {
		memberIds = append(memberIds, request.MemberID)
	}

	userProfiles, err := u.GetBatchUserProfiles(pctx, grpcUrl, memberIds)
	if err != nil {
		return []group.GroupRequestRes{}, err
	}
	groupRequestRes := make([]group.GroupRequestRes, 0)
	if len(userProfiles.UserProfiles) != 0 {
		for i := range requests {
			groupRequestRes = append(groupRequestRes, group.GroupRequestRes{
				GroupID:      requests[i].GroupID,
				MemberID:     requests[i].MemberID,
				Description:  requests[i].Description,
				CreatedAt:    requests[i].CreatedAt,
				UserID:       userProfiles.UserProfiles[i].UserId,
				Username:     userProfiles.UserProfiles[i].Username,
				UserPhotourl: userProfiles.UserProfiles[i].Photourl,
			})
		}
	}

	return groupRequestRes, nil
}

func (u *groupUsecase) GetGroupJoinRequestCount(pctx context.Context, groupId int) (int, error) {
	count, err := u.store.GetGroupRequestCount(pctx, int32(groupId))
	if err != nil {
		return -1, err
	}
	return int(count), nil
}

func (u *groupUsecase) GetUserJoinRequests(pctx context.Context, args groupdb.GetMemberPendingGroupRequestsParams) ([]groupdb.GroupRequest, error) {
	requests, err := u.store.GetMemberPendingGroupRequests(pctx, args)
	if err != nil {
		return []groupdb.GroupRequest{}, err
	}
	return requests, nil
}

func (u *groupUsecase) GetGroupById(pctx context.Context, groupId, userId int) (group.GetGroupByIdRes, error) {
	groupData, err := u.store.GetGroupByID(pctx, int32(groupId))
	if err != nil {
		return group.GetGroupByIdRes{}, err
	}

	memberInfo := groupdb.GetMemberInfoRow{}

	if userId != -1 {
		memberInfo, err = u.store.GetMemberInfo(pctx, groupdb.GetMemberInfoParams{
			MemberID: int32(userId),
			GroupID:  int32(groupId),
		})
	}

	if err != nil {
		if err != sql.ErrNoRows {
			return group.GetGroupByIdRes{}, err
		}

		requestInfo, reqErr := u.store.GetRequestFromGroup(pctx, groupdb.GetRequestFromGroupParams{
			GroupID:  int32(groupId),
			MemberID: int32(userId),
		})

		if reqErr != nil {
			return group.GetGroupByIdRes{}, reqErr
		}

		if len(requestInfo) != 0 {
			return group.GetGroupByIdRes{
				GroupInfo: groupData,
				UserId:    int32(userId),
				Status:    "requested",
			}, nil
		}

		return group.GetGroupByIdRes{
			GroupInfo: groupData,
			UserId:    int32(userId),
			Status:    "non-member",
		}, nil
	}

	return group.GetGroupByIdRes{
		GroupInfo: groupData,
		UserId:    int32(userId),
		Status:    memberInfo.Role,
	}, nil
}

func (u *groupUsecase) GetGroupMembersByGroupId(pctx context.Context, args groupdb.GetMembersByGroupIDParams, grpcUrl string) ([]group.GroupMemberRes, error) {
	groupMembers, err := u.store.GetMembersByGroupID(pctx, args)
	if err != nil {
		return nil, err
	}

	var memberIds []int32
	for _, request := range groupMembers {
		memberIds = append(memberIds, request.MemberID)
	}

	userProfiles, err := u.GetBatchUserProfiles(pctx, grpcUrl, memberIds)
	if err != nil {
		return nil, err
	}

	GroupMemberRes := make([]group.GroupMemberRes, 0)
	for i, member := range groupMembers {
		if i < len(userProfiles.UserProfiles) {
			GroupMemberRes = append(GroupMemberRes, group.GroupMemberRes{
				GroupID:      member.GroupID,
				MemberID:     member.MemberID,
				Role:         member.Role,
				CreatedAt:    member.CreatedAt,
				UserID:       userProfiles.UserProfiles[i].UserId,
				Username:     userProfiles.UserProfiles[i].Username,
				UserPhotourl: userProfiles.UserProfiles[i].Photourl,
			})
		}
	}

	return GroupMemberRes, nil
}

func (u *groupUsecase) ListGroups(pctx context.Context, args groupdb.ListGroupsParams) ([]groupdb.ListGroupsRow, error) {
	groups, err := u.store.ListGroups(pctx, args)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (u *groupUsecase) EditGroupName(pctx context.Context, args groupdb.EditGroupNameParams, editorId int32) (groupdb.GetGroupByIDRow, error) {
	role, err := u.GetRole(pctx, editorId, args.GroupID)
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	if role != group.Admin && role != group.Moderator {
		return groupdb.GetGroupByIDRow{}, errors.New("no permission")
	}

	if err := u.store.EditGroupName(pctx, args); err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	updatedGroup, err := u.GetGroupById(pctx, int(args.GroupID), -1)
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	return updatedGroup.GroupInfo, nil
}

func (u *groupUsecase) EditGroupPhoto(pctx context.Context, args groupdb.EditGroupPhotoParams, editorId int32) (groupdb.GetGroupByIDRow, error) {
	role, err := u.GetRole(pctx, editorId, args.GroupID)
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	if role != group.Admin && role != group.Moderator {
		return groupdb.GetGroupByIDRow{}, errors.New("no permission")
	}
	if err := u.store.EditGroupPhoto(pctx, args); err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}
	groupData, err := u.GetGroupById(pctx, int(args.GroupID), -1)
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	return groupData.GroupInfo, nil
}

func (u *groupUsecase) EditGroupVisibility(pctx context.Context, args groupdb.EditGroupVisibilityParams, editorId int32) (groupdb.GetGroupByIDRow, error) {
	role, err := u.GetRole(pctx, editorId, args.GroupID)
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	if role != group.Admin && role != group.Moderator {
		return groupdb.GetGroupByIDRow{}, errors.New("no permission")
	}
	if err := u.store.EditGroupVisibility(pctx, args); err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	groupData, err := u.GetGroupById(pctx, int(args.GroupID), -1)
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}
	return groupData.GroupInfo, nil
}

func (u *groupUsecase) EditGroupDescription(pctx context.Context, args groupdb.EditGroupDescriptionParams, editorId int32) (groupdb.GetGroupByIDRow, error) {
	role, err := u.GetRole(pctx, editorId, args.GroupID)
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	if role != group.Admin && role != group.Moderator {
		return groupdb.GetGroupByIDRow{}, errors.New("no permission")
	}

	if err := u.store.EditGroupDescription(pctx, args); err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	updatedGroup, err := u.GetGroupById(pctx, int(args.GroupID), -1)
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	return updatedGroup.GroupInfo, nil
}

func (u *groupUsecase) EditMemberRole(pctx context.Context, args groupdb.EditMemberRoleParams, editorId int32) (groupdb.GetMemberInfoRow, error) {

	role, err := u.GetRole(pctx, editorId, args.GroupID)
	if err != nil {
		return groupdb.GetMemberInfoRow{}, err
	}

	if (role != group.Admin && role != group.Moderator) || (args.Role != string(group.Moderator)) {
		return groupdb.GetMemberInfoRow{}, errors.New("no permission")
	}

	if err := u.store.EditMemberRole(pctx, args); err != nil {
		return groupdb.GetMemberInfoRow{}, err
	}

	memberInfoParams := groupdb.GetMemberInfoParams{
		MemberID: int32(args.MemberID),
		GroupID:  int32(args.GroupID),
	}

	updatedMember, err := u.store.GetMemberInfo(pctx, memberInfoParams)
	if err != nil {
		return groupdb.GetMemberInfoRow{}, err
	}

	return updatedMember, nil
}

func (u *groupUsecase) DeleteGroup(pctx context.Context, groupId int, editorId int32) error {
	role, err := u.GetRole(pctx, editorId, int32(groupId))
	if err != nil {
		return err
	}

	if role != group.Admin {
		return errors.New("no permission")
	}

	if err := u.store.DeleteGroup(pctx, int32(groupId)); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) DeleteGroupMember(pctx context.Context, args groupdb.DeleteMemberParams, editorId int32) error {
	role, err := u.GetRole(pctx, editorId, args.GroupID)
	if err != nil {
		return err
	}

	if role != group.Admin && role != group.Moderator && editorId != args.MemberID {
		return errors.New("no permission")
	}
	if err := u.store.DeleteMember(pctx, args); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) CreatePost(pctx context.Context, args groupdb.CreatePostParams) (groupdb.Post, error) {
	newPost, err := u.store.CreatePost(pctx, args)
	if err != nil {
		return groupdb.Post{}, err
	}
	_, err = u.IncreaseStreak(pctx, groupdb.GetStreakByMemberIDandGroupIDParams{
		MemberID: args.MemberID,
		GroupID:  args.GroupID,
	})
	if err != nil {
		return groupdb.Post{}, err
	}

	return newPost, nil
}

func (u *groupUsecase) GetPostByPostId(pctx context.Context, postId, userId int, grpcUrl string) (groupdb.Post, *userPb.GetUserProfileRes, *groupdb.PostReaction, error) {
	postInfo, err := u.store.GetPostByPostID(pctx, int32(postId))
	if err != nil {
		return groupdb.Post{}, nil, nil, err
	}
	profile, err := u.GetUserProfile(pctx, grpcUrl, &userPb.GetUserProfileReq{
		UserId: postInfo.MemberID,
	})
	if err != nil {
		return groupdb.Post{}, nil, nil, err
	}
	userReaction, err := u.store.GetReactionByPostIDAndUserID(pctx, groupdb.GetReactionByPostIDAndUserIDParams{
		PostID:   postInfo.PostID,
		MemberID: int32(userId),
	})
	if err != nil && err != sql.ErrNoRows {
		return groupdb.Post{}, nil, nil, err
	}
	if err == sql.ErrNoRows {
		return postInfo, profile, nil, nil
	}

	return postInfo, profile, &userReaction, nil
}

func (u *groupUsecase) GetPostsByGroupId(pctx context.Context, args groupdb.GetPostsByGroupIDParams, grpcUrl string) ([]group.PostByGroupRes, error) {
	posts, err := u.store.GetPostsByGroupID(pctx, args)
	if err != nil {
		return nil, err
	}

	var memberIds []int32
	for _, post := range posts {
		memberIds = append(memberIds, post.MemberID)
	}

	userProfiles, err := u.GetBatchUserProfiles(pctx, grpcUrl, memberIds)
	if err != nil {
		return nil, err
	}

	PostsByGroupRes := make([]group.PostByGroupRes, 0)
	for i, post := range posts {
		if i < len(userProfiles.UserProfiles) {
			PostsByGroupRes = append(PostsByGroupRes, group.PostByGroupRes{
				PostID:       post.PostID,
				MemberID:     post.MemberID,
				GroupID:      post.GroupID,
				PhotoUrl:     post.PhotoUrl,
				Description:  post.Description,
				CreatedAt:    post.CreatedAt,
				UserID:       userProfiles.UserProfiles[i].UserId,
				Username:     userProfiles.UserProfiles[i].Username,
				UserPhotourl: userProfiles.UserProfiles[i].Photourl,
			})
		}
	}

	return PostsByGroupRes, nil
}

// รอแก้
func (u *groupUsecase) GetPostsByUserId(pctx context.Context, args groupdb.GetPostsByMemberIDParams) ([]groupdb.Post, error) {
	posts, err := u.store.GetPostsByMemberID(pctx, args)
	if err != nil {
		return []groupdb.Post{}, err
	}
	return posts, nil
}

func (u *groupUsecase) GetPostsByGroupAndMemberId(pctx context.Context, args groupdb.GetPostsByGroupAndMemberIDParams) ([]groupdb.Post, error) {
	posts, err := u.store.GetPostsByGroupAndMemberID(pctx, args)
	if err != nil {
		return []groupdb.Post{}, err
	}
	return posts, nil
}

func (u *groupUsecase) EditPost(pctx context.Context, args groupdb.EditPostParams) (groupdb.Post, error) {
	if err := u.store.EditPost(pctx, args); err != nil {
		return groupdb.Post{}, err
	}

	updatedPost, err := u.store.GetPostByPostID(pctx, int32(args.PostID))
	if err != nil {
		return groupdb.Post{}, err
	}

	return updatedPost, nil
}

func (u *groupUsecase) DeletePost(pctx context.Context, postId int) error {
	if err := u.store.DeletePost(pctx, int32(postId)); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) GetPostsForOngoingFeedByMemberId(pctx context.Context, args groupdb.GetPostsForOngoingFeedByMemberIDParams, grpcUrl string) ([]group.PostsForFeedRes, error) {
	posts, err := u.store.GetPostsForOngoingFeedByMemberID(pctx, args)
	if err != nil {
		return []group.PostsForFeedRes{}, err
	}

	var memberIds []int32
	for _, post := range posts {
		memberIds = append(memberIds, post.MemberID)
	}

	postProfiles, err := u.GetBatchUserProfiles(pctx, grpcUrl, memberIds)
	if err != nil {
		return nil, err
	}

	profileUsernames := make(map[int32]string)
	profilePhotoUrls := make(map[int32]string)
	for _, profile := range postProfiles.UserProfiles {
		profileUsernames[profile.UserId] = profile.Username
		profilePhotoUrls[profile.UserId] = profile.Photourl
	}

	result := make([]group.PostsForFeedRes, len(posts))
	for i, post := range posts {
		streak, err := u.GetStreakByMemberIDandGroupID(pctx,
			groupdb.GetStreakByMemberIDandGroupIDParams{
				MemberID: post.MemberID,
				GroupID:  post.GroupID,
			},
		)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		var reaction string
		userReaction, err := u.store.GetReactionByPostIDAndUserID(pctx, groupdb.GetReactionByPostIDAndUserIDParams{
			PostID:   post.PostID,
			MemberID: args.MemberID,
		})
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		if err == sql.ErrNoRows {
			reaction = ""
		} else {
			reaction = userReaction.Reaction
		}

		result[i] = group.PostsForFeedRes{
			PostID:            post.PostID,
			MemberID:          post.MemberID,
			GroupID:           post.GroupID,
			PhotoUrl:          post.PhotoUrl,
			Description:       post.Description,
			CreatedAt:         post.CreatedAt,
			GroupName:         post.GroupName,
			GroupPhotoUrl:     post.GroupPhotoUrl,
			PosterUsername:    profileUsernames[post.MemberID],
			PosterPhotoUrl:    profilePhotoUrls[post.MemberID],
			TotalStreakCount:  streak.TotalStreakCount,
			WeeklyStreakCount: streak.WeeklyStreakCount,
			UserReaction:      reaction,
		}
	}

	return result, nil
}

func (u *groupUsecase) CreateReaction(pctx context.Context, args groupdb.CreateReactionParams) (groupdb.PostReaction, error) {
	newReaction, err := u.store.CreateReaction(pctx, args)
	if err != nil {
		return groupdb.PostReaction{}, err
	}
	return newReaction, nil
}

func (u *groupUsecase) GetReactionsByPostId(pctx context.Context, args groupdb.GetReactionsByPostIDParams, grpcUrl string) ([]group.ReactionPostRes, error) {
	reactions, err := u.store.GetReactionsByPostID(pctx, args)
	if err != nil {
		return nil, err
	}
	var memberIds []int32
	for _, request := range reactions {
		memberIds = append(memberIds, request.MemberID)
	}

	userProfiles, err := u.GetBatchUserProfiles(pctx, grpcUrl, memberIds)
	if err != nil {
		return nil, err
	}

	ReactionRes := make([]group.ReactionPostRes, 0)
	for i, member := range reactions {
		if i < len(userProfiles.UserProfiles) {
			ReactionRes = append(ReactionRes, group.ReactionPostRes{
				ReactionID:   member.ReactionID,
				MemberID:     member.MemberID,
				PostID:       member.PostID,
				CreatedAt:    member.CreatedAt,
				Reaction:     member.Reaction,
				UserID:       userProfiles.UserProfiles[i].UserId,
				Username:     userProfiles.UserProfiles[i].Username,
				UserPhotourl: userProfiles.UserProfiles[i].Photourl,
			})
		}
	}

	return ReactionRes, nil
}

func (u *groupUsecase) GetReactionsCountByPostId(pctx context.Context, postId int) (map[string]int, int, error) {
	reactionTypes, err := u.store.GetReactionsWithTypeByPostID(pctx, int32(postId))
	if err != nil {
		return nil, -1, err
	}
	totalReactions := 0
	reactionCount := make(map[string]int)
	for _, reaction := range reactionTypes {
		reactionCount[reaction]++
		totalReactions++
	}

	return reactionCount, totalReactions, nil
}

func (u *groupUsecase) EditReaction(pctx context.Context, args groupdb.EditReactionParams) (groupdb.PostReaction, error) {
	if err := u.store.EditReaction(pctx, args); err != nil {
		return groupdb.PostReaction{}, err
	}

	updatedReaction, err := u.store.GetReactionById(pctx, args.ReactionID)
	if err != nil {
		return groupdb.PostReaction{}, err
	}

	return updatedReaction, nil
}

func (u *groupUsecase) DeleteReaction(pctx context.Context, args groupdb.DeleteReactionParams) error {
	if err := u.store.DeleteReaction(pctx, args); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) CreateComment(pctx context.Context, args groupdb.CreateCommentParams) (groupdb.PostComment, error) {
	newComment, err := u.store.CreateComment(pctx, args)
	if err != nil {
		return groupdb.PostComment{}, err
	}
	return newComment, nil
}

func (u *groupUsecase) GetCommentsByPostId(pctx context.Context, args groupdb.GetCommentsByPostIDParams, grpcUrl string) ([]group.CommentRes, error) {
	comments, err := u.store.GetCommentsByPostID(pctx, args)
	if err != nil {
		return []group.CommentRes{}, err
	}

	var memberIds []int32
	for _, request := range comments {
		memberIds = append(memberIds, request.MemberID)
	}

	userProfiles, err := u.GetBatchUserProfiles(pctx, grpcUrl, memberIds)
	if err != nil {
		return nil, err
	}

	CommentRes := make([]group.CommentRes, 0)
	for i, member := range comments {
		if i < len(userProfiles.UserProfiles) {
			CommentRes = append(CommentRes, group.CommentRes{
				CommentID:    member.CommentID,
				PostID:       member.PostID,
				MemberID:     member.MemberID,
				Comment:      member.Comment,
				CreatedAt:    member.CreatedAt,
				UserID:       userProfiles.UserProfiles[i].UserId,
				Username:     userProfiles.UserProfiles[i].Username,
				UserPhotourl: userProfiles.UserProfiles[i].Photourl,
			})
		}
	}
	return CommentRes, nil
}

func (u *groupUsecase) GetCommentCountByPostId(pctx context.Context, postId int) (int, error) {
	commentCount, err := u.store.GetCommentsCountByPostID(pctx, int32(postId))
	if err != nil {
		return -1, err
	}
	return int(commentCount), nil
}

func (u *groupUsecase) EditComment(pctx context.Context, args groupdb.EditCommentParams) (groupdb.PostComment, error) {
	if err := u.store.EditComment(pctx, args); err != nil {
		return groupdb.PostComment{}, err
	}

	updatedComment, err := u.store.GetCommentByCommentId(pctx, args.CommentID)
	if err != nil {
		return groupdb.PostComment{}, err
	}

	return updatedComment, nil
}

func (u *groupUsecase) DeleteComment(pctx context.Context, commentId int) error {
	if err := u.store.DeleteComment(pctx, int32(commentId)); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) SaveToFirebaseStorage(pctx context.Context, bucketName, objectPath, filename string, file io.Reader) (string, error) {
	bucket, _ := u.storageClient.Bucket(bucketName)
	obj := bucket.Object(objectPath + "/" + filename)

	wc := obj.NewWriter(pctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	url := "https://firebasestorage.googleapis.com/v0/b/" + bucketName + "/o/" + objectPath + "%" + "2F" + filename + "?alt=media"

	return url, nil
}

func (u *groupUsecase) GetGroupLatestId(pctx context.Context) (int, error) {
	latest, err := u.store.GetGroupLatestId(pctx)
	if err != nil {
		return -1, err
	}
	return int(latest), nil
}

func (u *groupUsecase) GetPostLatestId(pctx context.Context) (int, error) {
	latest, err := u.store.GetPostLatestId(pctx)
	if err != nil {
		return -1, err
	}
	return int(latest), nil
}

func (u *groupUsecase) CreateNewTag(pctx context.Context, args groupdb.CreateNewTagParams) (groupdb.Tag, error) {
	newTag, err := u.store.CreateNewTag(pctx, args)
	if err != nil {
		return groupdb.Tag{}, err
	}
	return newTag, nil
}

func (u *groupUsecase) GetAvailableTags(pctx context.Context) ([]groupdb.Tag, error) {
	tags, err := u.store.GetAvailableTags(pctx)
	if err != nil {
		return []groupdb.Tag{}, err
	}
	return tags, nil
}

func (u *groupUsecase) GetGroupsByTagID(pctx context.Context, tagId int) ([]groupdb.Group, error) {
	groups, err := u.store.GetGroupsByTagID(pctx, int32(tagId))
	if err != nil {
		return []groupdb.Group{}, err
	}
	return groups, nil
}

func (u *groupUsecase) GroupMockData(pctx context.Context, count int) error {
	ctx := context.Background()

	err := u.store.InitializeCategory(ctx)
	if err != nil {
		return err
	}

	tagCategories := map[string][]string{
		"Sports and Fitness": {
			"Football", "Rock Climbing", "Basketball", "Volleyball", "Golf",
			"Boxing", "Badminton", "Bowling", "Ice skating", "Racquet",
			"Tennis", "Table tennis", "Snooker", "Pool", "Swimming",
			"Running", "Yoga and Pilates", "Karate", "Taekwondo", "Hiking",
			"Cycling", "Hockey", "Figure Skating", "Skiing",
		},
		"Learning and Development": {
			"Online courses", "Exam prep", "Investing", "Programming",
			"Language", "Public speaking", "SAT", "IELTS", "Midterm exam",
			"Final exam",
		},
		"Health and Wellness": {
			"Fitness and gym", "Dietary", "Bulking", "Vegan and J", "Meditation",
		},
		"Entertainment and Media": {
			"Movies", "Series", "Music", "Theater", "Podcasts",
		},
		"Hobbies and Leisure": {
			"Cooking", "Baking", "Gardening", "Planting", "Knitting",
			"Pottery", "Caligraphy", "Travelling", "Board games",
		},
	}

	existingTags := make(map[string]int32)

	for _, tagNames := range tagCategories {
		for _, tagName := range tagNames {
			_, err := u.createOrGetTagByCategory(ctx, tagName, existingTags)
			if err != nil {
				return err
			}
		}
	}

	randomFrequency := func() int32 {
		frequencies := []int32{1, 2, 3, 4, 5}
		return frequencies[rand.Intn(len(frequencies))]
	}

	for i := 0; i < count; i++ {
		groupName := fmt.Sprintf("Group %d", i+1)
		creatorID := int32(rand.Intn(10) + 1)
		tagID := rand.Intn(len(tagCategories)) + 1

		fake := faker.NewWithSeed(rand.NewSource(time.Now().UnixNano() + int64(i)))

		group, err := u.store.CreateGroup(ctx, groupdb.CreateGroupParams{
			GroupName:      groupName,
			GroupCreatorID: creatorID,
			PhotoUrl:       utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2F15.webp?alt=media"),
			TagID:          int32(tagID),
			Frequency:      randomFrequency(),
			MaxMembers:     25,
			GroupType:      "social",
			Description:    utils.ConvertStringToSqlNullString(fake.Lorem().Sentence(5)),
			Visibility:     true,
		})
		if err != nil {
			return err
		}

		_, err = u.store.AddGroupMember(ctx, groupdb.AddGroupMemberParams{
			GroupID:  group.GroupID,
			MemberID: creatorID,
			Role:     "creator",
		})
		if err != nil {
			return err
		}

		coLeaderCount := rand.Intn(4) + 1
		for j := 0; j < coLeaderCount; j++ {
			memberID := int32(rand.Intn(10) + 1)
			role := "co_leader"
			if memberID == creatorID {
				continue
			}
			exists, err := u.memberExistsInGroup(ctx, group.GroupID, memberID)
			if err != nil {
				return err
			}
			if exists {
				continue
			}

			_, err = u.store.AddGroupMember(ctx, groupdb.AddGroupMemberParams{
				GroupID:  group.GroupID,
				MemberID: memberID,
				Role:     role,
			})
			if err != nil {
				return err
			}
		}

		for j := 0; j < 8; j++ {
			memberID := int32(rand.Intn(10) + 1)
			role := "member"
			if memberID == creatorID {
				continue
			}

			exists, err := u.memberExistsInGroup(ctx, group.GroupID, memberID)
			if err != nil {
				return err
			}
			if exists {
				continue
			}

			_, err = u.store.AddGroupMember(ctx, groupdb.AddGroupMemberParams{
				GroupID:  group.GroupID,
				MemberID: memberID,
				Role:     role,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (u *groupUsecase) createOrGetTagByCategory(ctx context.Context, tagName string, existingTags map[string]int32) (int32, error) {
	tagID, ok := existingTags[tagName]
	if ok {
		return tagID, nil
	}

	newTag, err := u.store.CreateNewTag(ctx, groupdb.CreateNewTagParams{
		TagName:    tagName,
		IconUrl:    utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2F3.jpeg?alt=media"),
		CategoryID: utils.ConvertIntToSqlNullInt32(rand.Intn(5) + 1),
	})
	if err != nil {
		return 0, err
	}

	existingTags[tagName] = newTag.TagID
	return newTag.TagID, nil
}

func (u *groupUsecase) memberExistsInGroup(ctx context.Context, groupID, memberID int32) (bool, error) {
	_, err := u.store.CheckMemberInGroup(ctx, groupdb.CheckMemberInGroupParams{
		GroupID:  groupID,
		MemberID: memberID,
	})
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (u *groupUsecase) PostMockData(ctx context.Context, count int) error {
	groups, err := u.store.ListGroups(ctx, groupdb.ListGroupsParams{
		Limit:  int32(count),
		Offset: 0,
	})

	if err != nil {
		return err
	}

	for _, group := range groups {
		for i := 0; i < count; i++ {
			postID, err := u.createPost(ctx, group.GroupID)
			if err != nil {
				return err
			}

			numReactions := rand.Intn(10)
			for j := 0; j < numReactions; j++ {
				err := u.createReaction(ctx, postID)
				if err != nil {
					return err
				}
			}

			numComments := rand.Intn(10)
			for j := 0; j < numComments; j++ {
				err := u.createComment(ctx, postID)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (u *groupUsecase) createPost(ctx context.Context, groupID int32) (int32, error) {
	photoURL := sql.NullString{String: "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2F1.webp?alt=media", Valid: true}
	description := sql.NullString{String: "Lorem ipsum dolor sit amet", Valid: true}

	post, err := u.store.CreatePost(ctx, groupdb.CreatePostParams{
		MemberID:    1,
		GroupID:     groupID,
		PhotoUrl:    photoURL,
		Description: description,
	})
	if err != nil {
		return 0, err
	}

	// _, err := u.store.IncreaseStreak(ctx, groupdb.IncreaseStreakParams{

	// })

	return post.PostID, nil
}

func (u *groupUsecase) createReaction(ctx context.Context, postID int32) error {
	memberID := rand.Int31n(10) + 1
	reactions := []string{"like", "love", "haha", "wow", "sad", "angry"}
	reaction := reactions[rand.Intn(len(reactions))]

	_, err := u.store.CreateReaction(ctx, groupdb.CreateReactionParams{
		PostID:   postID,
		MemberID: memberID,
		Reaction: reaction,
	})
	return err
}

func (u *groupUsecase) createComment(ctx context.Context, postID int32) error {
	fake := faker.New()
	memberID := rand.Int31n(10) + 1
	comment := fake.Lorem().Sentence(5)

	_, err := u.store.CreateComment(ctx, groupdb.CreateCommentParams{
		PostID:   postID,
		MemberID: memberID,
		Comment:  comment,
	})
	return err
}

func (u *groupUsecase) GetTagByCategory(pctx context.Context, categoryID int32) ([]groupdb.Tag, error) {
	categoryId := utils.ConvertIntToSqlNullInt32(int(categoryID))
	tags, err := u.store.GetTagByCategory(pctx, categoryId)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (u *groupUsecase) GetTagByGroupId(pctx context.Context, groupId int32) (groupdb.GetTagByGroupIdRow, error) {
	tag, err := u.store.GetTagByGroupId(pctx, groupId)
	if err != nil {
		return groupdb.GetTagByGroupIdRow{}, err
	}
	return tag, nil
}

func (u *groupUsecase) GetGroupsByCategoryID(pctx context.Context, categoryId int32) ([]groupdb.GetGroupsByCategoryIDRow, error) {
	categoryIdNullInt := utils.ConvertIntToSqlNullInt32(int(categoryId))
	groups, err := u.store.GetGroupsByCategoryID(pctx, categoryIdNullInt)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (u *groupUsecase) GetTagByTagID(pctx context.Context, tagId int32) (groupdb.Tag, error) {
	tag, err := u.store.GetTagByTagID(pctx, tagId)
	if err != nil {
		return groupdb.Tag{}, err
	}
	return tag, nil
}

func (u *groupUsecase) EditTagName(pctx context.Context, args groupdb.EditTagNameParams) (groupdb.Tag, error) {
	// role, err := u.GetRole(pctx, editorId, args.TagID)
	// if err != nil {
	// 	return groupdb.Tag{}, err
	// }

	// if role != group.Admin && role != group.Moderator {
	// 	return groupdb.Tag{}, errors.New("no permission")
	// }

	if err := u.store.EditTagName(pctx, args); err != nil {
		return groupdb.Tag{}, err
	}

	updatedTag, err := u.GetTagByTagID(pctx, int32(args.TagID))
	if err != nil {
		return groupdb.Tag{}, err
	}

	return updatedTag, nil
}

func (u *groupUsecase) EditTagCategory(pctx context.Context, args groupdb.EditTagCategoryParams) (groupdb.Tag, error) {

	if err := u.store.EditTagCategory(pctx, args); err != nil {
		return groupdb.Tag{}, err
	}

	updatedTag, err := u.GetTagByTagID(pctx, int32(args.TagID))
	if err != nil {
		return groupdb.Tag{}, err
	}

	return updatedTag, nil
}

func (u *groupUsecase) EditTagIcon(pctx context.Context, args groupdb.EditTagIconParams) (groupdb.Tag, error) {
	if err := u.store.EditTagIcon(pctx, args); err != nil {
		return groupdb.Tag{}, err
	}

	updatedTag, err := u.GetTagByTagID(pctx, int32(args.TagID))
	if err != nil {
		return groupdb.Tag{}, err
	}

	return updatedTag, nil
}

func (u *groupUsecase) DeleteTag(pctx context.Context, tagId int) error {
	if err := u.store.DeleteTag(pctx, int32(tagId)); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) EditGroupTag(pctx context.Context, args groupdb.EditGroupTagParams, editorId int32) (groupdb.Group, error) {
	role, err := u.GetRole(pctx, editorId, args.GroupID)
	if err != nil {
		return groupdb.Group{}, err
	}

	if role != group.Admin && role != group.Moderator {
		return groupdb.Group{}, errors.New("no permission")
	}

	updatedGroup, err := u.store.EditGroupTag(pctx, args)
	if err != nil {
		return groupdb.Group{}, err
	}

	return updatedGroup, nil
}

func (u *groupUsecase) GetPostsForFollowingFeedByMemberId(pctx context.Context, userId, limit, offset int, grpcUrl string) ([]group.PostsForFeedRes, error) {
	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("error - gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	friends, err := conn.User().GetUserFriends(pctx, &userPb.GetUserFriendsReq{UserId: int32(userId)})
	if err != nil {
		log.Printf("error - gRPC GetUserFriends failed: %s", err.Error())
		return nil, errors.New("error: gRPC GetUserFriends failed")
	}

	friendIdArr := make([]int32, len(friends.Friends))
	friendUsernames := make(map[int32]string)
	friendPhotoUrls := make(map[int32]string)

	for i, friend := range friends.Friends {
		friendIdArr[i] = friend.UserId
		friendUsernames[friend.UserId] = friend.Username
		friendPhotoUrls[friend.UserId] = friend.Photourl
	}

	posts, err := u.store.GetPostsForFollowingFeedByMemberId(pctx, groupdb.GetPostsForFollowingFeedByMemberIdParams{
		Column1: friendIdArr,
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, err
	}

	result := make([]group.PostsForFeedRes, len(posts))
	for i, post := range posts {
		streak, err := u.GetStreakByMemberIDandGroupID(pctx,
			groupdb.GetStreakByMemberIDandGroupIDParams{
				MemberID: post.MemberID,
				GroupID:  post.GroupID,
			},
		)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		var reaction string
		userReaction, err := u.store.GetReactionByPostIDAndUserID(pctx, groupdb.GetReactionByPostIDAndUserIDParams{
			PostID:   post.PostID,
			MemberID: int32(userId),
		})
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		if err == sql.ErrNoRows {
			reaction = ""
		} else {
			reaction = userReaction.Reaction
		}
		result[i] = group.PostsForFeedRes{
			PostID:            post.PostID,
			MemberID:          post.MemberID,
			GroupID:           post.GroupID,
			PhotoUrl:          post.PhotoUrl,
			Description:       post.Description,
			CreatedAt:         post.CreatedAt,
			GroupName:         post.GroupName,
			GroupPhotoUrl:     post.GroupPhotoUrl,
			PosterUsername:    friendUsernames[post.MemberID],
			PosterPhotoUrl:    friendPhotoUrls[post.MemberID],
			TotalStreakCount:  streak.TotalStreakCount,
			WeeklyStreakCount: streak.WeeklyStreakCount,
			UserReaction:      reaction,
		}
	}

	return result, nil
}

func (u *groupUsecase) SearchGroupByGroupName(ctx context.Context, args groupdb.SearchGroupByGroupNameParams) ([]groupdb.SearchGroupByGroupNameRow, error) {
	groups, err := u.store.SearchGroupByGroupName(ctx, args)
	if err != nil {
		return nil, err
	}

	var visibleGroups []groupdb.SearchGroupByGroupNameRow
	for _, g := range groups {
		if g.Visibility && g.GroupType == "social" {
			visibleGroups = append(visibleGroups, g)
		}
	}

	return visibleGroups, nil
}

func (u *groupUsecase) DeleteUserData(ctx context.Context, userID int32) error {
	if err := u.store.DeleteGroupMembers(ctx, userID); err != nil {
		return fmt.Errorf("failed to delete group members: %v", err)
	}

	if err := u.store.DeleteGroupRequests(ctx, userID); err != nil {
		return fmt.Errorf("failed to delete group requests: %v", err)
	}

	if err := u.store.DeletePosts(ctx, userID); err != nil {
		return fmt.Errorf("failed to delete posts: %v", err)
	}

	if err := u.store.DeletePostReactions(ctx, userID); err != nil {
		return fmt.Errorf("failed to delete post reactions: %v", err)
	}

	if err := u.store.DeletePostComments(ctx, userID); err != nil {
		return fmt.Errorf("failed to delete post comments: %v", err)
	}

	return nil
}

func (u *groupUsecase) GetAcceptorGroupRequests(ctx context.Context, userId int32) ([]groupdb.GetAcceptorGroupRequestsRow, error) {
	requests, err := u.store.GetAcceptorGroupRequests(ctx, userId)
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (u *groupUsecase) GetAcceptorGroupRequestsCount(ctx context.Context, userId int32) (groupdb.GetAcceptorGroupRequestsCountRow, error) {
	count, err := u.store.GetAcceptorGroupRequestsCount(ctx, userId)
	if err != nil && err != sql.ErrNoRows {
		return groupdb.GetAcceptorGroupRequestsCountRow{}, err
	}
	if err == sql.ErrNoRows {
		return groupdb.GetAcceptorGroupRequestsCountRow{
			MemberID:     userId,
			RequestCount: 0,
		}, nil
	}
	return count, nil
}

func (u *groupUsecase) GetUserGroups(ctx context.Context, userId int32) (group.GetUserGroupRes, error) {
	groups, err := u.store.GetUserGroups(ctx, userId)
	if err != nil {
		return group.GetUserGroupRes{}, err
	}

	social := make([]group.GetUserGroupsModel, 0)
	personal := make([]group.GetUserGroupsModel, 0)

	for _, g := range groups {
		streak, err := u.GetStreakByMemberIDandGroupID(ctx, groupdb.GetStreakByMemberIDandGroupIDParams{
			MemberID: userId,
			GroupID:  g.GroupID,
		})
		if err != nil && err != sql.ErrNoRows {
			return group.GetUserGroupRes{}, err
		}
		var streakCount int32
		if err == sql.ErrNoRows {
			streakCount = 0
		} else {
			streakCount = streak.TotalStreakCount
		}
		res := group.GetUserGroupsModel{
			GroupID:    g.GroupID,
			GroupName:  g.GroupName,
			PhotoUrl:   g.PhotoUrl,
			GroupType:  g.GroupType,
			UserStreak: streakCount,
		}
		if g.GroupType == "social" {
			social = append(social, res)
		} else {
			personal = append(personal, res)
		}
	}

	return group.GetUserGroupRes{
		SocialGroups:   social,
		PersonalGroups: personal,
	}, nil
}

func (u *groupUsecase) GetStreakSetByGroupIdAndMemberId(ctx context.Context, args groupdb.GetStreakSetByGroupIdandUserIdParams) ([]groupdb.StreakSet, error) {
	streaks, err := u.store.GetStreakSetByGroupIdandUserId(ctx, args)
	if err != nil {
		return nil, err
	}
	return streaks, nil
}

// func (u *groupUsecase) EditStreakSetEndDate(ctx context.Context, args groupdb.EditStreakSetEndDateParams) ([]groupdb.StreakSet, error) {
// 	err := u.store.EditStreakSetEndDate(ctx, args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return u.store.GetStreakSetByGroupId(ctx, args.GroupID)
// }

func (u *groupUsecase) DeleteStreakSet(ctx context.Context, args groupdb.DeleteStreakSetParams) error {
	err := u.store.DeleteStreakSet(ctx, args)
	if err != nil {
		return err
	}
	return err
}

func (u *groupUsecase) CreateStreak(ctx context.Context, GroupId int32, MemberId int32) (groupdb.StreakSet, error) {
	today := time.Now().Truncate(24 * time.Hour)

	nextWeek := today.AddDate(0, 0, 7)
	endOfNextWeek := nextWeek.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	args := groupdb.CreateStreakSetParams{
		GroupID:  GroupId,
		MemberID: MemberId,
		EndDate:  endOfNextWeek,
	}
	streakSet, err := u.store.CreateStreakSet(ctx, args)
	if err != nil {
		return groupdb.StreakSet{}, err
	}

	_, err = u.store.CreateStreak(ctx, streakSet.StreakSetID)
	if err != nil {
		return groupdb.StreakSet{}, err
	}

	return streakSet, nil
}

func (u *groupUsecase) GetStreakByMemberId(ctx context.Context, MemberId int32) ([]groupdb.GetStreakByMemberIdRow, error) {
	streak, err := u.store.GetStreakByMemberId(ctx, MemberId)
	if err != nil {
		return nil, err
	}
	return streak, nil
}

func (u *groupUsecase) GetStreakByMemberIDandGroupID(ctx context.Context, args groupdb.GetStreakByMemberIDandGroupIDParams) (groupdb.GetStreakByMemberIDandGroupIDRow, error) {
	streak, err := u.store.GetStreakByMemberIDandGroupID(ctx, args)
	if err != nil {
		return groupdb.GetStreakByMemberIDandGroupIDRow{}, err
	}
	return streak, nil
}

func (u *groupUsecase) GetIncompletedStreakByUserID(ctx context.Context, memberId int32) ([]groupdb.GetIncompletedStreakByUserIDRow, error) {
	streak, err := u.store.GetIncompletedStreakByUserID(ctx, memberId)
	if err != nil {
		return nil, err
	}
	return streak, nil
}

func (u *groupUsecase) IncreaseStreak(ctx context.Context, args groupdb.GetStreakByMemberIDandGroupIDParams) (groupdb.Streak, error) {
	data, err := u.store.GetStreakByMemberIDandGroupID(ctx, args)
	if err != nil {
		return groupdb.Streak{}, err
	}

	if data.RecentDateAdded.Valid {
		if time.Since(data.RecentDateAdded.Time) < 24*time.Hour {
			return groupdb.Streak{}, errors.New("recentDateAdded is within 1 day, streak cannot be increased")
		}
	}

	streak, err := u.store.IncreaseStreak(ctx, groupdb.IncreaseStreakParams{
		MemberID: args.MemberID,
		GroupID:  args.GroupID,
	})
	if err != nil {
		return groupdb.Streak{}, err
	}

	FreqGroup, err := u.store.GetGroupByID(ctx, args.GroupID)
	if err != nil {
		return groupdb.Streak{}, err
	}

	if int32(FreqGroup.Frequency) <= streak.WeeklyStreakCount {
		u.store.EditCompleteStatus(ctx, groupdb.EditCompleteStatusParams{
			MemberID:  args.MemberID,
			GroupID:   args.GroupID,
			Completed: true,
		})
	}

	return streak, nil
}

func (u *groupUsecase) GetMaxStreakByMemberId(ctx context.Context, MemberId int32) (int32, error) {
	streak, err := u.store.GetMaxStreakUser(ctx, MemberId)
	if err != nil && err != sql.ErrNoRows {
		return -1, err
	}

	if err == sql.ErrNoRows {
		return 0, nil
	}

	return streak, nil
}

func (u *groupUsecase) MockupTag(ctx context.Context) error {
	// Insert tag categories
	err := u.store.InitializeCategory(ctx)
	if err != nil {
		return err
	}

	// Map of tag categories
	tagCategories := map[string][]string{
		"Sports and Fitness": {
			"Football", "Rock Climbing", "Basketball", "Volleyball", "Golf",
			"Boxing", "Badminton", "Bowling", "Ice skating", "Racquet",
			"Tennis", "Table tennis", "Snooker", "Pool", "Swimming",
			"Running", "Yoga and Pilates", "Karate", "Taekwondo", "Hiking",
			"Cycling", "Hockey", "Figure Skating", "Skiing",
		},
		"Learning and Development": {
			"Online courses", "Exam prep", "Investing", "Programming",
			"Language", "Public speaking", "SAT", "IELTS", "Midterm exam",
			"Final exam",
		},
		"Health and Wellness": {
			"Dietary", "Bulking", "Vegan and J", "Meditation",
		},
		"Entertainment and Media": {
			"Movies", "Series", "Music", "Theater", "Podcasts",
		},
		"Hobbies and Leisure": {
			"Cooking", "Baking", "Gardening", "Planting", "Knitting",
			"Pottery", "Caligraphy", "Travelling", "Board games",
		},
	}

	// Iterate over tag categories
	for category, tags := range tagCategories {
		// Get category ID from database
		categoryID, err := u.store.GetCategoryIDByName(ctx, category)
		if err != nil {
			return err
		}

		// Insert tags for each category
		for _, tagName := range tags {
			// Insert tag into the database
			_, err := u.store.CreateNewTag(ctx, groupdb.CreateNewTagParams{
				TagName:    tagName,
				IconUrl:    utils.ConvertStringToSqlNullString(""),
				CategoryID: utils.ConvertIntToSqlNullInt32(int(categoryID)),
			})
			if err != nil {
				return err
			}
		}
	}

	// Rest of the function to create mock post, comment, reaction, etc.

	return nil
}

func (u *groupUsecase) MockupGroup(ctx context.Context) error {

	groupDetails := map[string]map[string]interface{}{
		"group1": {
			"group_name":       "ISE football club",
			"group_creator_id": int32(1),
			"photo_url":        "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Ffootball.jpeg?alt=media&token=2ab15bda-2a4f-47e2-88c7-00c7a8597290",
			"frequency":        int32(1),
			"max_members":      int32(50),
			"group_type":       "social",
			"description":      "Weekly football at BBB football club, ma join gunn",
			"visibility":       true,
			"tag_id":           int32(19),
		},
		"group2": {
			"group_name":       "Intania Badminton",
			"group_creator_id": int32(2),
			"photo_url":        "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fbadminton.png?alt=media&token=2a532ca2-e441-4dd1-a1f5-b9692e2915e6",
			"frequency":        int32(1),
			"max_members":      int32(30),
			"group_type":       "Private",
			"description":      "Let's join our badminton squad from Engineering Faculty! We encourage 2 times a week!",
			"visibility":       true,
			"tag_id":           int32(25),
		},
		"group3": {
			"group_name":       "Coursera ganag",
			"group_creator_id": int32(5),
			"photo_url":        "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fonlinecourse.jpeg?alt=media&token=52d021cc-d775-4985-beea-226c5a7afd4f",
			"frequency":        int32(5),
			"max_members":      int32(30),
			"group_type":       "social",
			"description":      "Enhance your soft skills with us!",
			"visibility":       true,
			"tag_id":           int32(43),
		},
		"group4": {
			"group_name":       "Code nerdyy",
			"group_creator_id": int32(10),
			"photo_url":        "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fprogramming.png?alt=media&token=210e448e-6764-4a7b-8934-71380dc6c26f",
			"frequency":        int32(5),
			"max_members":      int32(30),
			"group_type":       "social",
			"description":      "Commit the code 5 times a week and you will receive nerdy trophy from us",
			"visibility":       true,
			"tag_id":           int32(46),
		},
		"group5": {
			"group_name":       "Midterm try hard gang",
			"group_creator_id": int32(5),
			"photo_url":        "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fmidtermexam.jpeg?alt=media&token=4e7dbe16-8ff8-4db4-aed0-50349b46b969",
			"frequency":        int32(6),
			"max_members":      int32(30),
			"group_type":       "social",
			"description":      "Study hard, no fail, no cry, no F, happy life",
			"visibility":       true,
			"tag_id":           int32(51),
		},
		"group6": {
			"group_name":       "ISE bodybuilder",
			"group_creator_id": int32(8),
			"photo_url":        "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fgym.png?alt=media&token=1c78b800-1c3e-407e-b038-f82873c614c1",
			"frequency":        int32(5),
			"max_members":      int32(30),
			"group_type":       "social",
			"description":      "Get up from your bed and start working out <3",
			"visibility":       true,
			"tag_id":           int32(53),
		},
		"group7": {
			"group_name":       "Serious movie discussion",
			"group_creator_id": int32(9),
			"photo_url":        "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fmovies.png?alt=media&token=02e85840-0d42-4065-a456-01bed8bd1015",
			"frequency":        int32(2),
			"max_members":      int32(30),
			"group_type":       "social",
			"description":      "Watch movie twice a week and share them here!",
			"visibility":       true,
			"tag_id":           int32(2),
		},
		"group8": {
			"group_name":       "K-drama stands",
			"group_creator_id": int32(4),
			"photo_url":        "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fseries.jpeg?alt=media&token=6758213e-394f-4df3-8d2c-f222080ffd8b",
			"frequency":        int32(1),
			"max_members":      int32(30),
			"group_type":       "social",
			"description":      "Put down all the work and enjoy K-drama once a week!",
			"visibility":       true,
			"tag_id":           int32(5),
		},
		"group9": {
			"group_name":       "Intania Music Club",
			"group_creator_id": int32(6),
			"photo_url":        "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fmusic.jpeg?alt=media&token=4a034376-5361-4cab-a09b-c59d3027083c",
			"frequency":        int32(5),
			"max_members":      int32(100),
			"group_type":       "social",
			"description":      "Post your music here or you will be cursed by spotify devil",
			"visibility":       true,
			"tag_id":           int32(7),
		},
		"group10": {
			"group_name":       "Travelling & Hanging out",
			"group_creator_id": int32(11),
			"photo_url":        "https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Ftravelling.png?alt=media&token=03b1961f-a6a9-4143-ae20-5038d79ca813",
			"frequency":        int32(1),
			"max_members":      int32(100),
			"group_type":       "social",
			"description":      "Share your beautiful journey here!",
			"visibility":       true,
			"tag_id":           int32(17),
		},
	}

	// Create mock groups
	for i := 1; i <= 10; i++ {
		groupDetail, ok := groupDetails[fmt.Sprintf("group%d", i)]
		if !ok {
			return errors.New(fmt.Sprintf("group%d not found", i))
		}

		frequencyStr, _ := groupDetail["frequency"].(string)
		frequency, _ := strconv.Atoi(frequencyStr)

		// Create the group
		createdGroup, err := u.store.CreateGroup(ctx, groupdb.CreateGroupParams{
			GroupName:      groupDetail["group_name"].(string),
			GroupCreatorID: groupDetail["group_creator_id"].(int32),
			PhotoUrl:       utils.ConvertStringToSqlNullString(groupDetail["photo_url"].(string)),
			Frequency:      int32(frequency),
			MaxMembers:     groupDetail["max_members"].(int32),
			GroupType:      groupDetail["group_type"].(string),
			Description:    utils.ConvertStringToSqlNullString(groupDetail["description"].(string)),
			Visibility:     groupDetail["visibility"].(bool),
			TagID:          groupDetail["tag_id"].(int32),
		})
		if err != nil {
			return err
		}

		// Use the created group's ID to insert into the group_members table
		_, err = u.store.AddGroupMember(ctx, groupdb.AddGroupMemberParams{
			GroupID:  createdGroup.GroupID,
			MemberID: groupDetail["group_creator_id"].(int32), // You need to define the creatorID variable
			Role:     "creator",
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *groupUsecase) MockupMember(ctx context.Context) error {
	for groupID := 1; groupID <= 10; groupID++ {
		group, err := u.store.GetGroupByID(ctx, int32(groupID))
		if err != nil {
			return err
		}

		requests, err := u.store.GetGroupRequests(ctx, groupdb.GetGroupRequestsParams{
			GroupID: int32(groupID),
			Limit:   1000,
			Offset:  0,
		})
		if err != nil {
			return err
		}

		for _, request := range requests {
			_, err := u.AcceptGroupRequest(ctx, groupdb.AcceptGroupRequestParams{
				GroupID:    groupID,
				MemberID:   int(request.MemberID),
				AcceptorId: int(group.GroupCreatorID),
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// create mock post
func (u *groupUsecase) MockupPost(ctx context.Context) error {
	postDetails := []struct {
		MemberID    int32
		GroupID     int32
		PhotoURL    sql.NullString
		Description sql.NullString
	}{
		// Mock post details
		{MemberID: 8, GroupID: 1, PhotoURL: utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Ffootball.jpg?alt=media&token=47e76d94-1a3c-4c64-ac93-be9c20bdc15f"), Description: utils.ConvertStringToSqlNullString("1st day practise")},
		{MemberID: 3, GroupID: 2, PhotoURL: utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fbadminton.jpg?alt=media&token=65167afe-1e56-4e1b-8fe7-9cfae50ae484"), Description: utils.ConvertStringToSqlNullString("Enjoy makk")},
		{MemberID: 1, GroupID: 3, PhotoURL: utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fonlinecourse.jpeg?alt=media&token=717da4a4-9049-41d2-8ffc-448d81f31c86"), Description: utils.ConvertStringToSqlNullString("Wanna sleep T_T")},
		{MemberID: 1, GroupID: 4, PhotoURL: utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fprogramming.JPG?alt=media&token=dbd8f503-2fea-43ec-9599-c1df1ac10515"), Description: utils.ConvertStringToSqlNullString("Grinding")},
		{MemberID: 6, GroupID: 5, PhotoURL: utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fmidterm.jpeg?alt=media&token=2a767e05-667c-4898-9175-e8729fe20c9b"), Description: utils.ConvertStringToSqlNullString("Cheat sheet done!")},
		{MemberID: 10, GroupID: 6, PhotoURL: utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Ffitness.JPG?alt=media&token=73336d8e-a753-44dc-89ff-d751d7aacf69"), Description: utils.ConvertStringToSqlNullString("Six packs is coming")},
		{MemberID: 3, GroupID: 7, PhotoURL: utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fmovies.jpeg?alt=media&token=44906b86-8c27-4ab1-b46f-96dcfe9cfab3"), Description: utils.ConvertStringToSqlNullString("I cried so hard T_T")},
		{MemberID: 5, GroupID: 8, PhotoURL: utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fseries.jpeg?alt=media&token=84f76ffd-fdc7-4af1-9ede-c12dc425b9cb"), Description: utils.ConvertStringToSqlNullString("Today I skipped K-drama na")},
		{MemberID: 8, GroupID: 9, PhotoURL: utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2FMusic.jpg?alt=media&token=2184d287-59d2-4cef-8c31-33b92e44a8b5"), Description: utils.ConvertStringToSqlNullString("SRV is my tune")},
		{MemberID: 13, GroupID: 10, PhotoURL: utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Ftravel.jpeg?alt=media&token=6e02713a-d086-4eca-bf69-257a3ebab442"), Description: utils.ConvertStringToSqlNullString("The weather is so fresh here!")},
	}

	for _, post := range postDetails {
		_, err := u.CreatePost(ctx, groupdb.CreatePostParams{
			MemberID:    post.MemberID,
			GroupID:     post.GroupID,
			PhotoUrl:    post.PhotoURL,
			Description: post.Description,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

//create mock comment

func (u *groupUsecase) MockupComment(ctx context.Context) error {
	// Define mock comment details
	commentDetails := []struct {
		PostID   int32
		MemberID int32
		Comment  string
	}{
		// Mock comment details

		{PostID: 1, MemberID: 1, Comment: "Goodjob"},
		{PostID: 2, MemberID: 3, Comment: "Impressive work!"},
		{PostID: 2, MemberID: 10, Comment: "Fantastic!"},
		{PostID: 10, MemberID: 4, Comment: "Outstanding!"},
		{PostID: 4, MemberID: 5, Comment: "Brilliant!"},
		{PostID: 2, MemberID: 2, Comment: "Excellent!"},
		{PostID: 5, MemberID: 4, Comment: "Amazing!"},
		{PostID: 7, MemberID: 6, Comment: "Awesome!"},
		{PostID: 3, MemberID: 2, Comment: "Goodjob"},
		{PostID: 4, MemberID: 1, Comment: "Impressive work!"},
		{PostID: 2, MemberID: 2, Comment: "Fantastic!"},
		{PostID: 2, MemberID: 3, Comment: "Outstanding!"},
		{PostID: 2, MemberID: 7, Comment: "Brilliant!"},
		{PostID: 1, MemberID: 3, Comment: "Excellent!"},
		{PostID: 2, MemberID: 2, Comment: "Amazing!"},
		{PostID: 5, MemberID: 14, Comment: "Awesome!"},
		{PostID: 1, MemberID: 1, Comment: "Goodjob"},
		{PostID: 2, MemberID: 2, Comment: "Impressive work!"},
		{PostID: 2, MemberID: 10, Comment: "Fantastic!"},
		{PostID: 10, MemberID: 4, Comment: "Outstanding!"},
		{PostID: 4, MemberID: 1, Comment: "Brilliant!"},
		{PostID: 2, MemberID: 2, Comment: "Excellent!"},
		{PostID: 5, MemberID: 5, Comment: "Amazing!"},
		{PostID: 7, MemberID: 3, Comment: "Awesome!"},
		{PostID: 3, MemberID: 2, Comment: "Goodjob"},
		{PostID: 4, MemberID: 1, Comment: "Impressive work!"},
		{PostID: 2, MemberID: 4, Comment: "Fantastic!"},
		{PostID: 2, MemberID: 3, Comment: "Outstanding!"},
		{PostID: 2, MemberID: 7, Comment: "Brilliant!"},
		{PostID: 1, MemberID: 9, Comment: "Excellent!"},
		{PostID: 2, MemberID: 2, Comment: "Amazing!"},
		{PostID: 5, MemberID: 8, Comment: "Awesome!"},
		{PostID: 1, MemberID: 1, Comment: "Goodjob"},
		{PostID: 2, MemberID: 2, Comment: "Impressive work!"},
		{PostID: 2, MemberID: 10, Comment: "Fantastic!"},
		{PostID: 10, MemberID: 4, Comment: "Outstanding!"},
		{PostID: 4, MemberID: 1, Comment: "Brilliant!"},
		{PostID: 3, MemberID: 2, Comment: "Excellent!"},
		{PostID: 5, MemberID: 5, Comment: "Amazing!"},
		{PostID: 7, MemberID: 3, Comment: "Awesome!"},
		{PostID: 2, MemberID: 2, Comment: "Goodjob"},
		{PostID: 4, MemberID: 1, Comment: "Impressive work!"},
		{PostID: 9, MemberID: 4, Comment: "Fantastic!"},
		{PostID: 2, MemberID: 3, Comment: "Outstanding!"},
		{PostID: 7, MemberID: 7, Comment: "Brilliant!"},
		{PostID: 1, MemberID: 9, Comment: "Excellent!"},
		{PostID: 2, MemberID: 2, Comment: "Amazing!"},
		{PostID: 5, MemberID: 8, Comment: "Awesome!"},

		// Add more mock comment details as needed
	}

	for _, comment := range commentDetails {
		_, err := u.store.CreateComment(ctx, groupdb.CreateCommentParams{
			PostID:   comment.PostID,
			MemberID: comment.MemberID,
			Comment:  comment.Comment,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// create mock reaction
func (u *groupUsecase) MockupReactions(ctx context.Context) error {
	// Define reactions
	reactions := []string{"like", "love", "haha", "wow", "sad", "angry"}
	postIDs := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for groupID := int32(1); groupID <= 10; groupID++ {
		memberIDs, err := u.store.GetMembersByGroupID(ctx, groupdb.GetMembersByGroupIDParams{
			GroupID: groupID,
			Limit:   10000,
			Offset:  0,
		})
		if err != nil {
			return err
		}

		for _, postID := range postIDs {
			for _, memberID := range memberIDs {
				reaction := reactions[rand.Intn(len(reactions))]
				_, err := u.store.CreateReaction(ctx, groupdb.CreateReactionParams{
					PostID:   postID,
					MemberID: memberID.MemberID,
					Reaction: reaction,
				})
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
