package groupUsecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
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
		GetGroupJoinRequests(pctx context.Context, groupId int, grpcUrl string) ([]group.GroupRequestRes, error)
		GetUserJoinRequests(pctx context.Context, userId int) ([]groupdb.GroupRequest, error)
		GetGroupById(pctx context.Context, groupId int) (groupdb.GetGroupByIDRow, error)
		GetGroupMembersByGroupId(pctx context.Context, groupId int) ([]groupdb.GroupMember, error)
		ListGroups(pctx context.Context, args groupdb.ListGroupsParams) ([]groupdb.Group, error)
		EditGroupName(pctx context.Context, args groupdb.EditGroupNameParams, editorId int32) (groupdb.GetGroupByIDRow, error)
		EditGroupPhoto(pctx context.Context, args groupdb.EditGroupPhotoParams, editorId int32) (groupdb.GetGroupByIDRow, error)
		EditGroupVisibility(pctx context.Context, args groupdb.EditGroupVisibilityParams, editorId int32) (groupdb.GetGroupByIDRow, error)
		EditGroupDescription(pctx context.Context, args groupdb.EditGroupDescriptionParams, editorId int32) (groupdb.GetGroupByIDRow, error)
		EditMemberRole(pctx context.Context, args groupdb.EditMemberRoleParams, editorId int32) (groupdb.GetMemberInfoRow, error)
		DeleteGroup(pctx context.Context, groupId int, editorId int32) error
		DeleteGroupMember(pctx context.Context, args groupdb.DeleteMemberParams, editorId int32) error
		CreatePost(pctx context.Context, args groupdb.CreatePostParams) (groupdb.Post, error)
		GetPostByPostId(pctx context.Context, postId int) (groupdb.Post, error)
		GetPostsByGroupId(pctx context.Context, groupId int) ([]groupdb.Post, error)
		GetPostsByUserId(pctx context.Context, userId int) ([]groupdb.Post, error)
		GetPostsByGroupAndMemberId(pctx context.Context, args groupdb.GetPostsByGroupAndMemberIDParams) ([]groupdb.Post, error)
		EditPost(pctx context.Context, args groupdb.EditPostParams) (groupdb.Post, error)
		DeletePost(pctx context.Context, postId int) error
		GetPostsForOngoingFeedByMemberId(pctx context.Context, userId int) ([]groupdb.GetPostsForOngoingFeedByMemberIDRow, error)
		CreateReaction(pctx context.Context, args groupdb.CreateReactionParams) (groupdb.PostReaction, error)
		GetReactionsByPostId(pctx context.Context, postId int) ([]groupdb.PostReaction, error)
		GetReactionsCountByPostId(pctx context.Context, postId int) (int, error)
		EditReaction(pctx context.Context, args groupdb.EditReactionParams) (groupdb.PostReaction, error)
		DeleteReaction(pctx context.Context, reactionId int) error
		CreateComment(pctx context.Context, args groupdb.CreateCommentParams) (groupdb.PostComment, error)
		GetCommentsByPostId(pctx context.Context, postId int) ([]groupdb.PostComment, error)
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
	return newGroup, nil
}

func (u *groupUsecase) SendRequestToJoinGroup(pctx context.Context, args groupdb.SendRequestToJoinGroupParams) (groupdb.GroupRequest, error) {
	newRequest, err := u.store.SendRequestToJoinGroup(pctx, args)
	if err != nil {
		return groupdb.GroupRequest{}, err
	}
	return newRequest, nil
}

func (u *groupUsecase) AcceptGroupRequest(pctx context.Context, args groupdb.AcceptGroupRequestParams) (groupdb.GroupMember, error) {
	role, err := u.GetRole(pctx, int32(args.AcceptorId), int32(args.GroupID))
	if err != nil {
		return groupdb.GroupMember{}, err
	}
	if role != group.Admin && role != group.Moderator {
		return groupdb.GroupMember{}, errors.New("no permission")
	}

	newMember, err := u.store.AcceptGroupRequest(pctx, args)
	if err != nil {
		return groupdb.GroupMember{}, err
	}
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

func (u *groupUsecase) GetGroupJoinRequests(pctx context.Context, groupId int, grpcUrl string) ([]group.GroupRequestRes, error) {
	requests, err := u.store.GetGroupRequests(pctx, int32(groupId))
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
				Username:     userProfiles.UserProfiles[i].Username,
				UserPhotourl: userProfiles.UserProfiles[i].Photourl,
			})
		}
	}

	return groupRequestRes, nil
}

func (u *groupUsecase) GetUserJoinRequests(pctx context.Context, userId int) ([]groupdb.GroupRequest, error) {
	requests, err := u.store.GetMemberPendingGroupRequests(pctx, int32(userId))
	if err != nil {
		return []groupdb.GroupRequest{}, err
	}
	return requests, nil
}

func (u *groupUsecase) GetGroupById(pctx context.Context, groupId int) (groupdb.GetGroupByIDRow, error) {
	groupData, err := u.store.GetGroupByID(pctx, int32(groupId))
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}
	return groupData, nil
}

func (u *groupUsecase) GetGroupMembersByGroupId(pctx context.Context, groupId int) ([]groupdb.GroupMember, error) {
	groupMembers, err := u.store.GetMembersByGroupID(pctx, int32(groupId))
	if err != nil {
		return []groupdb.GroupMember{}, err
	}
	return groupMembers, nil
}

func (u *groupUsecase) ListGroups(pctx context.Context, args groupdb.ListGroupsParams) ([]groupdb.Group, error) {
	groups, err := u.store.ListGroups(pctx, args)
	if err != nil {
		return []groupdb.Group{}, err
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

	updatedGroup, err := u.GetGroupById(pctx, int(args.GroupID))
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	return updatedGroup, nil
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
	groupData, err := u.GetGroupById(pctx, int(args.GroupID))
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	return groupData, nil
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

	groupData, err := u.GetGroupById(pctx, int(args.GroupID))
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}
	return groupData, nil
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

	updatedGroup, err := u.GetGroupById(pctx, int(args.GroupID))
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}

	return updatedGroup, nil
}

func (u *groupUsecase) EditMemberRole(pctx context.Context, args groupdb.EditMemberRoleParams, editorId int32) (groupdb.GetMemberInfoRow, error) {

	role, err := u.GetRole(pctx, editorId, args.GroupID)
	if err != nil {
		return groupdb.GetMemberInfoRow{}, err
	}

	if role != group.Admin && role != group.Moderator || args.Role != string(group.Moderator) {
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

	if role != group.Admin && role != group.Moderator {
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
	return newPost, nil
}

func (u *groupUsecase) GetPostByPostId(pctx context.Context, postId int) (groupdb.Post, error) {
	postInfo, err := u.store.GetPostByPostID(pctx, int32(postId))
	if err != nil {
		return groupdb.Post{}, err
	}
	return postInfo, nil
}

func (u *groupUsecase) GetPostsByGroupId(pctx context.Context, groupId int) ([]groupdb.Post, error) {
	posts, err := u.store.GetPostsByGroupID(pctx, int32(groupId))
	if err != nil {
		return []groupdb.Post{}, err
	}
	return posts, nil
}

func (u *groupUsecase) GetPostsByUserId(pctx context.Context, userId int) ([]groupdb.Post, error) {
	posts, err := u.store.GetPostsByMemberID(pctx, int32(userId))
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

func (u *groupUsecase) GetPostsForOngoingFeedByMemberId(pctx context.Context, userId int) ([]groupdb.GetPostsForOngoingFeedByMemberIDRow, error) {
	posts, err := u.store.GetPostsForOngoingFeedByMemberID(pctx, int32(userId))
	if err != nil {
		return []groupdb.GetPostsForOngoingFeedByMemberIDRow{}, err
	}
	return posts, nil
}

func (u *groupUsecase) CreateReaction(pctx context.Context, args groupdb.CreateReactionParams) (groupdb.PostReaction, error) {
	newReaction, err := u.store.CreateReaction(pctx, args)
	if err != nil {
		return groupdb.PostReaction{}, err
	}
	return newReaction, nil
}

func (u *groupUsecase) GetReactionsByPostId(pctx context.Context, postId int) ([]groupdb.PostReaction, error) {
	reactions, err := u.store.GetReactionsByPostID(pctx, int32(postId))
	if err != nil {
		return []groupdb.PostReaction{}, err
	}
	return reactions, nil
}

func (u *groupUsecase) GetReactionsCountByPostId(pctx context.Context, postId int) (int, error) {
	reactionCount, err := u.store.GetReactionsCountByPostID(pctx, int32(postId))
	if err != nil {
		return -1, err
	}
	return int(reactionCount), nil
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

func (u *groupUsecase) DeleteReaction(pctx context.Context, reactionId int) error {
	if err := u.store.DeleteReaction(pctx, int32(reactionId)); err != nil {
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

func (u *groupUsecase) GetCommentsByPostId(pctx context.Context, postId int) ([]groupdb.PostComment, error) {
	comments, err := u.store.GetCommentsByPostID(pctx, int32(postId))
	if err != nil {
		return []groupdb.PostComment{}, err
	}
	return comments, nil
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

	tagCategories := []string{"tag1", "tag2", "tag5", "tag4", "tag6", "tag7"}

	existingTags := make(map[string]int32)

	for _, category := range tagCategories {
		_, err := u.createOrGetTagByCategory(ctx, category, existingTags)
		if err != nil {
			return err
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

		group, err := u.store.CreateGroup(ctx, groupdb.CreateGroupParams{
			GroupName:      groupName,
			GroupCreatorID: creatorID,
			PhotoUrl:       utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2F1.jpg?alt=media"),
			TagID:          int32(tagID),
			Frequency:      sql.NullInt32{Int32: randomFrequency(), Valid: true},
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

func (u *groupUsecase) createOrGetTagByCategory(ctx context.Context, category string, existingTags map[string]int32) (int32, error) {
	tagID, ok := existingTags[category]
	if ok {
		return tagID, nil
	}

	newTag, err := u.store.CreateNewTag(ctx, groupdb.CreateNewTagParams{
		TagName:    category,
		IconUrl:    utils.ConvertStringToSqlNullString("https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2F3.jpeg?alt=media"),
		CategoryID: utils.ConvertIntToSqlNullInt32(rand.Intn(5) + 1),
	})
	if err != nil {
		return 0, err
	}

	existingTags[category] = newTag.TagID
	return newTag.TagID, nil
}

func (u *groupUsecase) memberExistsInGroup(ctx context.Context, groupID, memberID int32) (bool, error) {
	members, err := u.store.GetMembersByGroupID(ctx, groupID)
	if err != nil {
		return false, err
	}
	for _, member := range members {
		if member.MemberID == memberID {
			return true, nil
		}
	}
	return false, nil
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
