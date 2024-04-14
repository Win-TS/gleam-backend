package groupHandler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/modules/group"
	"github.com/Win-TS/gleam-backend.git/modules/group/groupUsecase"

	userPb "github.com/Win-TS/gleam-backend.git/modules/user/userPb"
	groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
	"github.com/Win-TS/gleam-backend.git/pkg/request"
	"github.com/Win-TS/gleam-backend.git/pkg/response"
	"github.com/Win-TS/gleam-backend.git/pkg/utils"
	"github.com/labstack/echo/v4"
)

type (
	GroupHttpHandlerService interface {
		CreateNewGroup(c echo.Context) error
		SendRequestToJoinGroup(c echo.Context) error
		AcceptGroupRequest(c echo.Context) error
		DeclineGroupRequest(c echo.Context) error
		GetGroupJoinRequests(c echo.Context) error
		GetUserJoinRequests(c echo.Context) error
		GetGroupById(c echo.Context) error
		GetGroupMembersByGroupId(c echo.Context) error
		ListGroups(c echo.Context) error
		EditGroupName(c echo.Context) error
		EditGroupPhoto(c echo.Context) error
		EditMemberRole(c echo.Context) error
		EditGroupVisibility(c echo.Context) error
		EditGroupDescription(c echo.Context) error
		DeleteGroup(c echo.Context) error
		DeleteGroupMember(c echo.Context) error
		CreatePost(c echo.Context) error
		GetPostByPostId(c echo.Context) error
		GetPostsByGroupId(c echo.Context) error
		GetPostsByUserId(c echo.Context) error
		GetPostsByGroupAndMemberId(c echo.Context) error
		EditPost(c echo.Context) error
		DeletePost(c echo.Context) error
		GetPostsForOngoingFeedByMemberId(c echo.Context) error
		CreateReaction(c echo.Context) error
		GetReactionsByPostId(c echo.Context) error
		GetReactionsCountByPostId(c echo.Context) error
		EditReaction(c echo.Context) error
		DeleteReaction(c echo.Context) error
		CreateComment(c echo.Context) error
		GetCommentsByPostId(c echo.Context) error
		GetCommentCountByPostId(c echo.Context) error
		EditComment(c echo.Context) error
		DeleteComment(c echo.Context) error
		CreateTag(c echo.Context) error
		GetAvailableTags(c echo.Context) error
		GetGroupsByTagID(c echo.Context) error
		GetTagByCategory(c echo.Context) error
		GetTagByGroupId(c echo.Context) error
		GetGroupsByCategoryID(c echo.Context) error
		EditTagName(c echo.Context) error
		EditTagCategory(c echo.Context) error
		EditTagIcon(c echo.Context) error
		DeleteTag(c echo.Context) error
		EditGroupTag(c echo.Context) error
		GroupMockData(c echo.Context) error
		PostMockData(c echo.Context) error
		GetPostsForFollowingFeedByMemberId(c echo.Context) error
		SearchGroupByGroupName(c echo.Context) error
		GetAcceptorGroupRequests(c echo.Context) error
		GetAcceptorGroupRequestsCount(c echo.Context) error
		GetUserGroups(c echo.Context) error
	}

	groupHttpHandler struct {
		cfg          *config.Config
		groupUsecase groupUsecase.GroupUsecaseService
	}
)

func NewGroupHttpHandler(cfg *config.Config, groupUsecase groupUsecase.GroupUsecaseService) GroupHttpHandlerService {
	return &groupHttpHandler{cfg, groupUsecase}
}

func (h *groupHttpHandler) CreateNewGroup(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(group.NewGroupReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err := h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(req.GroupCreatorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	file, _ := c.FormFile("photo")
	var url string

	fileid, err := h.groupUsecase.GetGroupLatestId(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}
	fileidStr := strconv.Itoa(fileid)

	if file != nil {
		src, err := file.Open()
		if err != nil {
			return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		bucketName := h.cfg.Firebase.StorageBucket
		objectPath := "groupphoto"

		url, err = h.groupUsecase.SaveToFirebaseStorage(c.Request().Context(), bucketName, objectPath, (fileidStr + utils.GetFileExtension(file.Filename)), src)
		if err != nil {
			return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	args := &groupdb.CreateGroupParams{
		GroupName:      req.GroupName,
		GroupCreatorID: int32(req.GroupCreatorId),
		PhotoUrl:       utils.ConvertStringToSqlNullString(url),
		TagID:          int32(req.TagID),
		Description:    utils.ConvertStringToSqlNullString(req.Description),
		Frequency:      utils.ConvertIntToSqlNullInt32(req.Frequency),
		MaxMembers:     int32(req.MaxMembers),
		GroupType:      req.GroupType,
		Visibility:     req.Visibility,
	}

	newGroup, err := h.groupUsecase.CreateNewGroup(ctx, *args)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	user, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(req.GroupCreatorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Success: group created",
		"data":    newGroup,
		"creator": user,
	})
}

func (h *groupHttpHandler) SendRequestToJoinGroup(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(group.RequestToJoinGroupReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err := h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(req.MemberID),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	newRequest, err := h.groupUsecase.SendRequestToJoinGroup(ctx, groupdb.SendRequestToJoinGroupParams{
		GroupID:     int32(req.GroupID),
		MemberID:    int32(req.MemberID),
		Description: utils.ConvertStringToSqlNullString(req.Description),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	user, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(req.MemberID),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success":            true,
		"message":            "Success: send join group request success",
		"data":               newRequest,
		"requestUserProfile": user,
	})
}

func (h *groupHttpHandler) AcceptGroupRequest(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.AcceptGroupRequestParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	idToSearch := []int32{int32(req.AcceptorId), int32(req.MemberID)}

	_, err := h.groupUsecase.GetBatchUserProfiles(ctx, h.cfg.Grpc.UserUrl, idToSearch)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	newMember, err := h.groupUsecase.AcceptGroupRequest(ctx, *req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	// return response.SuccessResponse(c, http.StatusCreated, newMember)
	user, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(req.MemberID),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	acceptor, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(req.AcceptorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success":  true,
		"message":  "Success: Accepted join request",
		"data":     newMember,
		"member":   user,
		"acceptor": acceptor,
	})
}

func (h *groupHttpHandler) DeclineGroupRequest(c echo.Context) error {
	ctx := context.Background()

	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	memberId, err := strconv.Atoi(c.QueryParam("member_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	declinerId, err := strconv.Atoi(c.QueryParam("decliner_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	idToSearch := []int32{int32(memberId), int32(declinerId)}

	_, err = h.groupUsecase.GetBatchUserProfiles(ctx, h.cfg.Grpc.UserUrl, idToSearch)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := h.groupUsecase.DeclineGroupRequest(ctx, groupdb.DeleteRequestToJoinGroupParams{
		GroupID:  int32(groupId),
		MemberID: int32(memberId),
	}, declinerId); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	user, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(memberId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	acceptor, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(declinerId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	// return response.SuccessResponse(c, http.StatusOK, "Success: group request declined")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: Declined group request",
		// "data":    updatedComment,
		"member":   user,
		"decliner": acceptor,
	})
}

func (h *groupHttpHandler) GetGroupJoinRequests(c echo.Context) error {
	ctx := context.Background()
	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	joinRequests, err := h.groupUsecase.GetGroupJoinRequests(ctx, groupdb.GetGroupRequestsParams{
		GroupID: int32(groupId),
		Limit:   int32(limit),
		Offset:  int32(offset),
	}, h.cfg.Grpc.UserUrl)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, joinRequests)
}

func (h *groupHttpHandler) GetUserJoinRequests(c echo.Context) error {
	ctx := context.Background()
	memberId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(memberId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	userRequests, err := h.groupUsecase.GetUserJoinRequests(ctx, groupdb.GetMemberPendingGroupRequestsParams{
		MemberID: int32(memberId),
		Limit:    int32(limit),
		Offset:   int32(offset),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, userRequests)
}

func (h *groupHttpHandler) GetGroupById(c echo.Context) error {
	ctx := context.Background()
	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	groupInfo, err := h.groupUsecase.GetGroupById(ctx, groupId, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, groupInfo)
}

func (h *groupHttpHandler) GetGroupMembersByGroupId(c echo.Context) error {
	ctx := context.Background()
	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	groupMembers, err := h.groupUsecase.GetGroupMembersByGroupId(ctx, groupdb.GetMembersByGroupIDParams{
		GroupID: int32(groupId),
		Limit:   int32(limit),
		Offset:  int32(offset),
	}, h.cfg.Grpc.UserUrl)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, groupMembers)
}

func (h *groupHttpHandler) ListGroups(c echo.Context) error {
	ctx := context.Background()
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	Offset, err := strconv.Atoi(c.QueryParam("offset"))

	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	args := &groupdb.ListGroupsParams{
		Limit:  int32(limit),
		Offset: int32(Offset),
	}

	groupList, err := h.groupUsecase.ListGroups(ctx, *args)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, groupList)
}

func (h *groupHttpHandler) EditGroupName(c echo.Context) error {
	ctx := context.Background()

	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	newGroupName, ok := requestBody["group_name"].(string)
	if !ok {
		return response.ErrResponse(c, http.StatusBadRequest, "New group name is missing or invalid in request body")
	}

	memberID, ok := requestBody["member_id"].(float64)
	if !ok {
		return response.ErrResponse(c, http.StatusBadRequest, "member_id is missing or invalid in request body")
	}

	memberIdInt := int(memberID)

	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(memberIdInt),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	args := &groupdb.EditGroupNameParams{
		GroupID:   int32(groupId),
		GroupName: newGroupName,
	}

	updatedGroup, err := h.groupUsecase.EditGroupName(ctx, *args, int32(memberIdInt))

	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	Editor, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(memberID),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: group name edited",
		"data":    updatedGroup,
		"editor":  Editor,
	})
}

func (h *groupHttpHandler) EditGroupPhoto(c echo.Context) error {
	ctx := c.Request().Context()
	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Invalid group ID")
	}

	file, err := c.FormFile("photo")
	if err != nil && err != http.ErrMissingFile {
		return response.ErrResponse(c, http.StatusBadRequest, "Invalid file")
	}

	editorIdStr := c.FormValue("editor_id")
	editorId, err := strconv.Atoi(editorIdStr)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Invalid editor ID")
	}

	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(editorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	var url string
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return response.ErrResponse(c, http.StatusInternalServerError, "Failed to open file")
		}
		defer src.Close()
		bucketName := h.cfg.Firebase.StorageBucket
		objectPath := "groupphoto"
		url, err = h.groupUsecase.SaveToFirebaseStorage(ctx, bucketName, objectPath, strconv.Itoa(groupId)+utils.GetFileExtension(file.Filename), src)
		if err != nil {
			return response.ErrResponse(c, http.StatusInternalServerError, "Failed to upload file to storage")
		}
	}

	req := &groupdb.EditGroupPhotoParams{
		GroupID:  int32(groupId),
		PhotoUrl: utils.ConvertStringToSqlNullString(url),
	}

	updatedGroup, err := h.groupUsecase.EditGroupPhoto(ctx, *req, int32(editorId))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, "Failed to update group photo")
	}
	Editor, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(editorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: group photo edited",
		"data":    updatedGroup,
		"editor":  Editor,
	})
}

func (h *groupHttpHandler) EditGroupVisibility(c echo.Context) error {
	ctx := context.Background()

	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	visibility_str := c.QueryParam("visibility")
	visibility := false
	if visibility_str != "true" && visibility_str != "false" {
		return response.ErrResponse(c, http.StatusBadRequest, "Invalid visibility value")
	}
	if visibility_str == "true" {
		visibility = true
	}

	editorId, err := strconv.Atoi(c.QueryParam("editor_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(editorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	args := &groupdb.EditGroupVisibilityParams{
		GroupID:    int32(groupId),
		Visibility: visibility,
	}

	updatedGroup, err := h.groupUsecase.EditGroupVisibility(ctx, *args, int32(editorId))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	Editor, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(editorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: group visibility edited",
		"data":    updatedGroup,
		"editor":  Editor,
	})
}

func (h *groupHttpHandler) EditGroupDescription(c echo.Context) error {
	ctx := context.Background()

	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	newDescription := c.QueryParam("description")
	if newDescription == "" {
		return response.ErrResponse(c, http.StatusBadRequest, "Description is missing")
	}

	editorId, err := strconv.Atoi(c.QueryParam("editor_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(editorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	args := &groupdb.EditGroupDescriptionParams{
		GroupID:     int32(groupId),
		Description: utils.ConvertStringToSqlNullString(newDescription),
	}

	updatedGroup, err := h.groupUsecase.EditGroupDescription(ctx, *args, int32(editorId))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	Editor, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(editorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: group description edited",
		"data":    updatedGroup,
		"editor":  Editor,
	})
}

func (h *groupHttpHandler) EditMemberRole(c echo.Context) error {
	ctx := context.Background()

	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	memberId, err := strconv.Atoi(c.QueryParam("member_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	editorId, ok := requestBody["editor_id"].(float64)
	if !ok {
		return response.ErrResponse(c, http.StatusBadRequest, "EditorId is missing")
	}
	newGroupRole, ok := requestBody["role"].(string)
	if !ok {
		return response.ErrResponse(c, http.StatusBadRequest, "New group role is missing or invalid in request body")
	}

	targetId := int32(memberId)

	idToSearch := []int32{targetId, int32(editorId)}

	_, err = h.groupUsecase.GetBatchUserProfiles(ctx, h.cfg.Grpc.UserUrl, idToSearch)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	args := &groupdb.EditMemberRoleParams{
		GroupID:  int32(groupId),
		MemberID: targetId,
		Role:     newGroupRole,
	}

	updatedMember, err := h.groupUsecase.EditMemberRole(ctx, *args, int32(editorId))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	Editor, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(editorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	Target, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(targetId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: group role edited",
		"data":    updatedMember,
		"editor":  Editor,
		"target":  Target,
	})
}

func (h *groupHttpHandler) DeleteGroup(c echo.Context) error {
	ctx := context.Background()
	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	editorId, err := strconv.Atoi(c.QueryParam("editor_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(editorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := h.groupUsecase.DeleteGroup(ctx, groupId, int32(editorId)); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	Editor, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(editorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: group deleted",
		"editor":  Editor,
	})
}

func (h *groupHttpHandler) DeleteGroupMember(c echo.Context) error {
	ctx := context.Background()

	editorId, err := strconv.Atoi(c.QueryParam("editor_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	memberId, err := strconv.Atoi(c.QueryParam("member_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	idToSearch := []int32{int32(memberId), int32(editorId)}

	_, err = h.groupUsecase.GetBatchUserProfiles(ctx, h.cfg.Grpc.UserUrl, idToSearch)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	req := &groupdb.DeleteMemberParams{
		MemberID: int32(memberId),
		GroupID:  int32(groupId),
	}

	if err := h.groupUsecase.DeleteGroupMember(ctx, *req, int32(editorId)); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	// return response.SuccessResponse(c, http.StatusOK, "Success: member deleted from group")
	Editor, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(editorId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	Target, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(memberId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: member deleted from group",
		"group":   groupId,
		"editor":  Editor,
		"target":  Target,
	})
}

func (h *groupHttpHandler) CreatePost(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(group.NewPostReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	file, _ := c.FormFile("photo")
	var url string

	fileid, err := h.groupUsecase.GetPostLatestId(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}
	fileidStr := strconv.Itoa(fileid)

	if file != nil {
		src, err := file.Open()
		if err != nil {
			return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		bucketName := h.cfg.Firebase.StorageBucket
		objectPath := "postphoto"

		url, err = h.groupUsecase.SaveToFirebaseStorage(c.Request().Context(), bucketName, objectPath, (fileidStr + utils.GetFileExtension(file.Filename)), src)
		if err != nil {
			return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(req.MemberID),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	args := &groupdb.CreatePostParams{
		MemberID:    int32(req.MemberID),
		GroupID:     int32(req.GroupID),
		Description: utils.ConvertStringToSqlNullString(req.Description),
		PhotoUrl:    utils.ConvertStringToSqlNullString(url),
	}

	newPost, err := h.groupUsecase.CreatePost(ctx, *args)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	// return response.SuccessResponse(c, http.StatusCreated, newPost)
	Member, err := h.groupUsecase.GetUserProfile(ctx, h.cfg.Grpc.UserUrl, &userPb.GetUserProfileReq{
		UserId: int32(req.MemberID),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: group role edited",
		"data":    newPost,
		"member":  Member,
	})
}

func (h *groupHttpHandler) GetPostByPostId(c echo.Context) error {
	ctx := context.Background()
	postId, err := strconv.Atoi(c.QueryParam("post_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	postInfo, Member, err := h.groupUsecase.GetPostByPostId(ctx, postId, h.cfg.Grpc.UserUrl)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: get post by id success",
		"data":    postInfo,
		"member":  Member,
	})
}

func (h *groupHttpHandler) GetPostsByGroupId(c echo.Context) error {
	ctx := context.Background()
	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	postsInGroup, err := h.groupUsecase.GetPostsByGroupId(ctx, groupdb.GetPostsByGroupIDParams{
		GroupID: int32(groupId),
		Limit:   int32(limit),
		Offset:  int32(offset),
	}, h.cfg.Grpc.UserUrl)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, postsInGroup)
}

func (h *groupHttpHandler) GetPostsByUserId(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(userId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	userPosts, err := h.groupUsecase.GetPostsByUserId(ctx, groupdb.GetPostsByMemberIDParams{
		MemberID: int32(userId),
		Limit:    int32(limit),
		Offset:   int32(offset),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, userPosts)
}

func (h *groupHttpHandler) GetPostsByGroupAndMemberId(c echo.Context) error {
	ctx := context.Background()
	memberId, err := strconv.Atoi(c.QueryParam("member_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}
	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	req := &groupdb.GetPostsByGroupAndMemberIDParams{
		GroupID:  int32(groupId),
		MemberID: int32(memberId),
	}
	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(memberId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	posts, err := h.groupUsecase.GetPostsByGroupAndMemberId(ctx, *req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, posts)
}

func (h *groupHttpHandler) EditPost(c echo.Context) error {
	ctx := context.Background()
	postIdStr := c.QueryParam("post_id")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	requestBody := make(map[string]string)

	if err := c.Bind(&requestBody); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Failed to parse request body")
	}

	if len(requestBody) == 0 {
		return response.ErrResponse(c, http.StatusBadRequest, "Request body is empty")
	}

	description, ok := requestBody["description"]
	if !ok {
		return response.ErrResponse(c, http.StatusBadRequest, "New description is missing in request body")
	}

	args := &groupdb.EditPostParams{
		PostID:      int32(postId),
		Description: utils.ConvertStringToSqlNullString(description),
	}

	updatedPost, err := h.groupUsecase.EditPost(ctx, *args)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: post edited",
		"data":    updatedPost,
	})
}

func (h *groupHttpHandler) DeletePost(c echo.Context) error {
	ctx := context.Background()
	postId, err := strconv.Atoi(c.QueryParam("post_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(postId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := h.groupUsecase.DeletePost(ctx, postId); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: post deleted")
}

func (h *groupHttpHandler) GetPostsForOngoingFeedByMemberId(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(userId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	feedPosts, err := h.groupUsecase.GetPostsForOngoingFeedByMemberId(ctx, groupdb.GetPostsForOngoingFeedByMemberIDParams{
		MemberID: int32(userId),
		Limit:    int32(limit),
		Offset:   int32(offset),
	}, h.cfg.Grpc.UserUrl)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, feedPosts)
}

func (h *groupHttpHandler) CreateReaction(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.CreateReactionParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err := h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(req.MemberID),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	newReaction, err := h.groupUsecase.CreateReaction(ctx, *req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, newReaction)
}

func (h *groupHttpHandler) GetReactionsByPostId(c echo.Context) error {
	ctx := context.Background()
	postId, err := strconv.Atoi(c.QueryParam("post_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	reactions, err := h.groupUsecase.GetReactionsByPostId(ctx, groupdb.GetReactionsByPostIDParams{
		PostID: int32(postId),
		Limit:  int32(limit),
		Offset: int32(offset),
	}, h.cfg.Grpc.UserUrl)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, reactions)
}

func (h *groupHttpHandler) GetReactionsCountByPostId(c echo.Context) error {
	ctx := context.Background()
	postId, err := strconv.Atoi(c.QueryParam("post_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	reactionsCount, err := h.groupUsecase.GetReactionsCountByPostId(ctx, postId)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	// return response.SuccessResponse(c, http.StatusOK, reactionsCount)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"post_id": postId,
		"data":    reactionsCount,
	})
}

func (h *groupHttpHandler) EditReaction(c echo.Context) error {
	ctx := context.Background()
	reactionId, err := strconv.Atoi(c.QueryParam("reaction_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	requestBody := make(map[string]string)

	if err := c.Bind(&requestBody); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Failed to parse request body")
	}

	if len(requestBody) == 0 {
		return response.ErrResponse(c, http.StatusBadRequest, "Request body is empty")
	}

	reaction, ok := requestBody["reaction"]
	if !ok {
		return response.ErrResponse(c, http.StatusBadRequest, "New reaction is missing in request body")
	}
	args := &groupdb.EditReactionParams{
		ReactionID: int32(reactionId),
		Reaction:   reaction,
	}

	updatedReaction, err := h.groupUsecase.EditReaction(ctx, *args)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: reaction edited",
		"data":    updatedReaction,
	})
}

func (h *groupHttpHandler) DeleteReaction(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.DeleteReactionParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.groupUsecase.DeleteReaction(ctx, *req); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: reaction deleted")
}

func (h *groupHttpHandler) CreateComment(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.CreateCommentParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	_, err := h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(req.MemberID),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	newComment, err := h.groupUsecase.CreateComment(ctx, *req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, newComment)
}

// bug
func (h *groupHttpHandler) GetCommentsByPostId(c echo.Context) error {
	ctx := context.Background()
	postId, err := strconv.Atoi(c.QueryParam("post_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	comments, err := h.groupUsecase.GetCommentsByPostId(ctx, groupdb.GetCommentsByPostIDParams{
		PostID: int32(postId),
		Limit:  int32(limit),
		Offset: int32(offset),
	}, h.cfg.Grpc.UserUrl)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, comments)
}

func (h *groupHttpHandler) GetCommentCountByPostId(c echo.Context) error {
	ctx := context.Background()
	postId, err := strconv.Atoi(c.QueryParam("post_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	commentCount, err := h.groupUsecase.GetCommentCountByPostId(ctx, postId)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, commentCount)
}

func (h *groupHttpHandler) EditComment(c echo.Context) error {
	ctx := context.Background()
	comment_id, err := strconv.Atoi(c.QueryParam("comment_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	requestBody := make(map[string]string)

	if err := c.Bind(&requestBody); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Failed to parse request body")
	}

	if len(requestBody) == 0 {
		return response.ErrResponse(c, http.StatusBadRequest, "Request body is empty")
	}

	comment, ok := requestBody["comment"]
	if !ok {
		return response.ErrResponse(c, http.StatusBadRequest, "New comment is missing in request body")
	}

	args := &groupdb.EditCommentParams{
		CommentID: int32(comment_id),
		Comment:   comment,
	}

	updatedComment, err := h.groupUsecase.EditComment(ctx, *args)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: group photo edited",
		"data":    updatedComment,
	})

}

func (h *groupHttpHandler) DeleteComment(c echo.Context) error {
	ctx := context.Background()
	commentId, err := strconv.Atoi(c.QueryParam("comment_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.groupUsecase.DeleteComment(ctx, commentId); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: comment deleted")
}

func (h *groupHttpHandler) CreateTag(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(group.NewTagReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	newTag, err := h.groupUsecase.CreateNewTag(ctx, groupdb.CreateNewTagParams{
		TagName:    req.TagName,
		IconUrl:    utils.ConvertStringToSqlNullString(req.IconUrl),
		CategoryID: utils.ConvertIntToSqlNullInt32(req.CategoryId),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, newTag)
}

func (h *groupHttpHandler) GetAvailableTags(c echo.Context) error {
	ctx := context.Background()

	tags, err := h.groupUsecase.GetAvailableTags(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, tags)
}

func (h *groupHttpHandler) GetGroupsByTagID(c echo.Context) error {
	ctx := context.Background()

	tagId, err := strconv.Atoi(c.QueryParam("tag_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	groups, err := h.groupUsecase.GetGroupsByTagID(ctx, tagId)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, groups)
}

func (h *groupHttpHandler) GetTagByCategory(c echo.Context) error {
	ctx := context.Background()
	categoryId, err := strconv.Atoi(c.QueryParam("category_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	tag, err := h.groupUsecase.GetTagByCategory(ctx, int32(categoryId))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, tag)
}

func (h *groupHttpHandler) GetTagByGroupId(c echo.Context) error {
	ctx := context.Background()
	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	tag, err := h.groupUsecase.GetTagByGroupId(ctx, int32(groupId))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, tag)
}

func (h *groupHttpHandler) GetGroupsByCategoryID(c echo.Context) error {
	ctx := context.Background()
	categoryID, err := strconv.Atoi(c.QueryParam("category_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	reactions, err := h.groupUsecase.GetGroupsByCategoryID(ctx, int32(categoryID))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, reactions)
}

func (h *groupHttpHandler) EditTagName(c echo.Context) error {
	ctx := context.Background()

	tagId, err := strconv.Atoi(c.QueryParam("tag_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	newTagName, ok := requestBody["tag_name"].(string)
	if !ok {
		return response.ErrResponse(c, http.StatusBadRequest, "New tag name is missing or invalid in request body")
	}

	args := &groupdb.EditTagNameParams{
		TagID:   int32(tagId),
		TagName: newTagName,
	}

	updatedTag, err := h.groupUsecase.EditTagName(ctx, *args)

	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: tag name edited",
		"data":    updatedTag,
	})
}

func (h *groupHttpHandler) EditTagCategory(c echo.Context) error {
	ctx := context.Background()

	tagId, err := strconv.Atoi(c.QueryParam("tag_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	tagCategoryFloat, ok := requestBody["category_id"].(float64)
	if !ok {
		return response.ErrResponse(c, http.StatusBadRequest, "New tag category is missing or invalid in request body")
	}

	newTagCategory := int(tagCategoryFloat)

	args := &groupdb.EditTagCategoryParams{
		TagID:      int32(tagId),
		CategoryID: utils.ConvertIntToSqlNullInt32(newTagCategory),
	}

	updatedTag, err := h.groupUsecase.EditTagCategory(ctx, *args)

	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: tag category edited",
		"data":    updatedTag,
	})
}

func (h *groupHttpHandler) EditTagIcon(c echo.Context) error {
	ctx := c.Request().Context()
	tagId, err := strconv.Atoi(c.QueryParam("tag_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Invalid tag ID")
	}

	IconStr := c.FormValue("icon_url")

	req := &groupdb.EditTagIconParams{
		TagID:   int32(tagId),
		IconUrl: utils.ConvertStringToSqlNullString(IconStr),
	}

	updatedTag, err := h.groupUsecase.EditTagIcon(ctx, *req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, "Failed to update tag icon")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: tag icon edited",
		"data":    updatedTag,
	})
}

func (h *groupHttpHandler) DeleteTag(c echo.Context) error {
	ctx := context.Background()
	tagId, err := strconv.Atoi(c.QueryParam("tag_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.groupUsecase.DeleteTag(ctx, tagId); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: Tag deleted")
}

func (h *groupHttpHandler) EditGroupTag(c echo.Context) error {
	ctx := context.Background()

	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	newTag, ok := requestBody["tag_id"].(float64)
	if !ok {
		return response.ErrResponse(c, http.StatusBadRequest, "New tag is missing or invalid in request body")
	}

	editorID, ok := requestBody["editor_id"].(float64)
	if !ok {
		return response.ErrResponse(c, http.StatusBadRequest, "member_id is missing or invalid in request body")
	}

	_, err = h.groupUsecase.SearchUser(ctx, h.cfg.Grpc.UserUrl, &userPb.SearchUserReq{
		UserId: int32(editorID),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	args := &groupdb.EditGroupTagParams{
		GroupID: int32(groupId),
		TagID:   int32(newTag),
	}

	updatedGroup, err := h.groupUsecase.EditGroupTag(ctx, *args, int32(editorID))

	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Success: group name edited",
		"data":    updatedGroup,
	})
}

func (h *groupHttpHandler) GroupMockData(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := struct {
		Count int `json:"count"`
	}{}

	if err := wrapper.Bind(&req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	err := h.groupUsecase.GroupMockData(ctx, req.Count)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, fmt.Sprintf("%d group data created", req.Count))

}

func (h *groupHttpHandler) PostMockData(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := struct {
		Count int `json:"count"`
	}{}

	if err := wrapper.Bind(&req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	err := h.groupUsecase.PostMockData(ctx, req.Count)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, fmt.Sprintf("%d Post data created", req.Count))
}

func (h *groupHttpHandler) GetPostsForFollowingFeedByMemberId(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Invalid userId")
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	feedPosts, err := h.groupUsecase.GetPostsForFollowingFeedByMemberId(ctx, userId, limit, offset, h.cfg.Grpc.UserUrl)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, feedPosts)
}

func (h *groupHttpHandler) SearchGroupByGroupName(c echo.Context) error {
	ctx := context.Background()
	groupname := c.QueryParam("groupname")

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	groupInfo, err := h.groupUsecase.SearchGroupByGroupName(ctx, groupdb.SearchGroupByGroupNameParams{
		Column1: utils.ConvertStringToSqlNullString(groupname),
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, groupInfo)
}

func (h *groupHttpHandler) GetAcceptorGroupRequests(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	groupRequests, err := h.groupUsecase.GetAcceptorGroupRequests(ctx, int32(userId))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, groupRequests)
}

func (h *groupHttpHandler) GetAcceptorGroupRequestsCount(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	groupRequestsCount, err := h.groupUsecase.GetAcceptorGroupRequestsCount(ctx, int32(userId))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, groupRequestsCount)
}

func (h *groupHttpHandler) GetUserGroups(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	groups, err := h.groupUsecase.GetUserGroups(ctx, int32(userId))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, groups)
}
