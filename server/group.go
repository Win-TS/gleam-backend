package server

import (
	"log"

	"github.com/Win-TS/gleam-backend.git/modules/group/groupHandler"
	"github.com/Win-TS/gleam-backend.git/modules/group/groupUsecase"
	groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
)

func (s *server) groupService() {
	postgresDB, ok := s.db.(groupdb.Store)
	if !ok {
		log.Fatal("Unsupported database type")
		return
	}
	usecase := groupUsecase.NewGroupUsecase(postgresDB, s.storage)
	httpHandler := groupHandler.NewGroupHttpHandler(s.cfg, usecase)
	grpcHandler := groupHandler.NewGroupGrpcHandler(usecase)

	_ = grpcHandler

	group := s.app.Group("/group_v1")
	post := s.app.Group("/post_v1")
	reaction := s.app.Group("/reaction_v1")
	comment := s.app.Group("/comment_v1")

	// Health Check
	group.GET("", s.healthCheckService)

	// Group Endpoints
	group.POST("/group", httpHandler.CreateNewGroup)
	group.POST("/newmember", httpHandler.NewGroupMember)
	group.GET("/group", httpHandler.GetGroupById)
	group.GET("/groupmembers", httpHandler.GetGroupMembersByGroupId)
	group.GET("/listgroups", httpHandler.ListGroups)
	group.PATCH("/editgroupname", httpHandler.EditGroupName)
	group.PATCH("/editgroupphoto", httpHandler.EditGroupPhoto)
	group.PATCH("/editmemberrole", httpHandler.EditMemberRole)
	group.DELETE("/group", httpHandler.DeleteGroup)
	group.DELETE("/groupmember", httpHandler.DeleteGroupMember)

	// Post Endpoints
	post.POST("/post", httpHandler.CreatePost)
	post.GET("/post", httpHandler.GetPostByPostId)
	post.GET("/groupposts", httpHandler.GetPostsByGroupId)
	post.GET("/userposts", httpHandler.GetPostsByUserId)
	post.GET("/groupuserposts", httpHandler.GetPostsByGroupAndMemberId)
	post.PATCH("/post", httpHandler.EditPost)
	post.DELETE("/post", httpHandler.DeletePost)
	post.GET("/feedposts", httpHandler.GetPostsForFeedByMemberId)

	// Reaction Endpoints
	reaction.POST("/reaction", httpHandler.CreateReaction)
	reaction.GET("/postreactions", httpHandler.GetReactionsByPostId)
	reaction.GET("/postreactioncount", httpHandler.GetReactionsCountByPostId)
	reaction.PATCH("/reaction", httpHandler.EditReaction)
	reaction.DELETE("/reaction", httpHandler.DeleteReaction)

	// Comment Endpoints
	comment.POST("/comment", httpHandler.CreateComment)
	comment.GET("/postcomments", httpHandler.GetCommentsByPostId)
	comment.GET("/postcommentcount", httpHandler.GetCommentCountByPostId)
	comment.PATCH("/comment", httpHandler.EditComment)
	comment.DELETE("/comment", httpHandler.DeleteComment)
}
