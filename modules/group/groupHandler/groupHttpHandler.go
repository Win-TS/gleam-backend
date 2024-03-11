package groupHandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Win-TS/gleam-backend.git/config"
	"github.com/Win-TS/gleam-backend.git/modules/group"
	"github.com/Win-TS/gleam-backend.git/modules/group/groupUsecase"
	groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
	"github.com/Win-TS/gleam-backend.git/pkg/request"
	"github.com/Win-TS/gleam-backend.git/pkg/response"
	"github.com/Win-TS/gleam-backend.git/pkg/utils"
	"github.com/labstack/echo/v4"
)

type (
	GroupHttpHandlerService interface {
		CreateNewGroup(c echo.Context) error
		NewGroupMember(c echo.Context) error
		GetGroupById(c echo.Context) error
		GetGroupMembersByGroupId(c echo.Context) error
		ListGroups(c echo.Context) error
		EditGroupName(c echo.Context) error
		EditGroupPhoto(c echo.Context) error
		EditMemberRole(c echo.Context) error
		DeleteGroup(c echo.Context) error
		DeleteGroupMember(c echo.Context) error
		CreatePost(c echo.Context) error
		GetPostByPostId(c echo.Context) error
		GetPostsByGroupId(c echo.Context) error
		GetPostsByUserId(c echo.Context) error
		GetPostsByGroupAndMemberId(c echo.Context) error
		EditPost(c echo.Context) error
		DeletePost(c echo.Context) error
		GetPostsForFeedByMemberId(c echo.Context) error
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
		// GroupMockData(c echo.Context) error
		// PostMockData(c echo.Context) error
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
	}

	newGroup, err := h.groupUsecase.CreateNewGroup(ctx, *args)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, newGroup)
}

func (h *groupHttpHandler) NewGroupMember(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.AddGroupMemberParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	newMember, err := h.groupUsecase.NewGroupMember(ctx, *req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, newMember)
}

func (h *groupHttpHandler) GetGroupById(c echo.Context) error {
	ctx := context.Background()
	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	groupInfo, err := h.groupUsecase.GetGroupById(ctx, groupId)
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

	groupMembers, err := h.groupUsecase.GetGroupMembersByGroupId(ctx, groupId)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, groupMembers)
}

func (h *groupHttpHandler) ListGroups(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.ListGroupsParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	groupList, err := h.groupUsecase.ListGroups(ctx, *req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, groupList)
}

func (h *groupHttpHandler) EditGroupName(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.EditGroupNameParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.groupUsecase.EditGroupName(ctx, *req); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: group name edited")
}

func (h *groupHttpHandler) EditGroupPhoto(c echo.Context) error {
	ctx := context.Background()
	fileidStr := c.FormValue("group_id")
	groupId, err := strconv.Atoi(fileidStr)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	file, _ := c.FormFile("photo")
	var url string

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

	req := &groupdb.EditGroupPhotoParams{
		GroupID:  int32(groupId),
		PhotoUrl: utils.ConvertStringToSqlNullString(url),
	}

	if err := h.groupUsecase.EditGroupPhoto(ctx, *req); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: group photo edited")
}

func (h *groupHttpHandler) EditMemberRole(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.EditMemberRoleParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.groupUsecase.EditMemberRole(ctx, *req); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: member role edited")
}

func (h *groupHttpHandler) DeleteGroup(c echo.Context) error {
	ctx := context.Background()
	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.groupUsecase.DeleteGroup(ctx, groupId); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: group deleted")
}

func (h *groupHttpHandler) DeleteGroupMember(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.DeleteMemberParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.groupUsecase.DeleteGroupMember(ctx, *req); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: member deleted from group")
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

	return response.SuccessResponse(c, http.StatusCreated, newPost)
}

func (h *groupHttpHandler) GetPostByPostId(c echo.Context) error {
	ctx := context.Background()
	postId, err := strconv.Atoi(c.QueryParam("post_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	postInfo, err := h.groupUsecase.GetPostByPostId(ctx, postId)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, postInfo)
}

func (h *groupHttpHandler) GetPostsByGroupId(c echo.Context) error {
	ctx := context.Background()
	groupId, err := strconv.Atoi(c.QueryParam("group_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	postsInGroup, err := h.groupUsecase.GetPostsByGroupId(ctx, groupId)
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

	userPosts, err := h.groupUsecase.GetPostsByUserId(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, userPosts)
}

func (h *groupHttpHandler) GetPostsByGroupAndMemberId(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.GetPostsByGroupAndMemberIDParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	posts, err := h.groupUsecase.GetPostsByGroupAndMemberId(ctx, *req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, posts)
}

func (h *groupHttpHandler) EditPost(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(group.EditPostReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	args := &groupdb.EditPostParams{
		PostID:      int32(req.PostID),
		Description: utils.ConvertStringToSqlNullString(req.Description),
	}

	if err := h.groupUsecase.EditPost(ctx, *args); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: post edited")
}

func (h *groupHttpHandler) DeletePost(c echo.Context) error {
	ctx := context.Background()
	postId, err := strconv.Atoi(c.QueryParam("post_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.groupUsecase.DeletePost(ctx, postId); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: post deleted")
}

func (h *groupHttpHandler) GetPostsForFeedByMemberId(c echo.Context) error {
	ctx := context.Background()
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	feedPosts, err := h.groupUsecase.GetPostsForFeedByMemberId(ctx, userId)
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

	reactions, err := h.groupUsecase.GetReactionsByPostId(ctx, postId)
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

	return response.SuccessResponse(c, http.StatusOK, reactionsCount)
}

func (h *groupHttpHandler) EditReaction(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.EditReactionParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.groupUsecase.EditReaction(ctx, *req); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: reaction edited")
}

func (h *groupHttpHandler) DeleteReaction(c echo.Context) error {
	ctx := context.Background()
	reactionId, err := strconv.Atoi(c.QueryParam("reaction_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.groupUsecase.DeleteReaction(ctx, reactionId); err != nil {
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

	newComment, err := h.groupUsecase.CreateComment(ctx, *req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, newComment)
}

func (h *groupHttpHandler) GetCommentsByPostId(c echo.Context) error {
	ctx := context.Background()
	postId, err := strconv.Atoi(c.QueryParam("post_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	comments, err := h.groupUsecase.GetCommentsByPostId(ctx, postId)
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
	wrapper := request.ContextWrapper(c)

	req := new(groupdb.EditCommentParams)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.groupUsecase.EditComment(ctx, *req); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success: comment edited")
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
		TagName: req.TagName,
		IconUrl: utils.ConvertStringToSqlNullString(req.IconUrl),
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

// func (h *groupHttpHandler) GroupMockData(c echo.Context) error {
// 	ctx := context.Background()
// 	wrapper := request.ContextWrapper(c)

// 	req := struct {
// 		Count int `json:"count"`
// 	}{}

// 	if err := wrapper.Bind(&req); err != nil {
// 		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
// 	}

// 	err := h.groupUsecase.GroupMockData(ctx, req.Count)
// 	if err != nil {
// 		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
// 	}

// 	return response.SuccessResponse(c, http.StatusOK, fmt.Sprintf("%d group data created", req.Count))

// }

// func (h *groupHttpHandler) PostMockData(c echo.Context) error {
// 	ctx := context.Background()
// 	wrapper := request.ContextWrapper(c)

// 	req := struct {
// 		Count int `json:"count"`
// 	}{}

// 	if err := wrapper.Bind(&req); err != nil {
// 		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
// 	}

// 	err := h.groupUsecase.PostMockData(ctx, req.Count)
// 	if err != nil {
// 		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
// 	}

// 	return response.SuccessResponse(c, http.StatusOK, fmt.Sprintf("%d Post data created", req.Count))
// }
