package server

import (
	"log"

	"github.com/Win-TS/gleam-backend.git/modules/group/groupHandler"
	"github.com/Win-TS/gleam-backend.git/modules/group/groupUsecase"
	groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
	groupPb "github.com/Win-TS/gleam-backend.git/modules/group/groupPb"
	"github.com/Win-TS/gleam-backend.git/pkg/grpcconn"
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

	// gRPC
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(s.cfg, s.cfg.Grpc.GroupUrl)
		groupPb.RegisterGroupGrpcServiceServer(grpcServer, grpcHandler)
		log.Printf("User gRPC server listening on %s", s.cfg.Grpc.GroupUrl)
		grpcServer.Serve(lis)
	}()

	group := s.app.Group("/group_v1")
	post := s.app.Group("/post_v1")
	reaction := s.app.Group("/reaction_v1")
	comment := s.app.Group("/comment_v1")
	tag := s.app.Group("/tag_v1")

	// Health Check
	group.GET("", s.healthCheckService)

	//Mock data
	group.POST("/mock", httpHandler.GroupMockData)

	// Group Endpoints
	group.POST("/group", httpHandler.CreateNewGroup)
	group.POST("/requesttojoin", httpHandler.SendRequestToJoinGroup)
	group.POST("/acceptrequest", httpHandler.AcceptGroupRequest)
	group.DELETE("/declinerequest", httpHandler.DeclineGroupRequest)
	group.GET("/grouprequests", httpHandler.GetGroupJoinRequests)
	group.GET("/userrequests", httpHandler.GetUserJoinRequests)
	group.GET("/group", httpHandler.GetGroupById)
	group.GET("/groupmembers", httpHandler.GetGroupMembersByGroupId)
	group.GET("/listgroups", httpHandler.ListGroups)
	group.PATCH("/editgroupname", httpHandler.EditGroupName)
	group.PATCH("/editgroupphoto", httpHandler.EditGroupPhoto)
	group.PATCH("/editmemberrole", httpHandler.EditMemberRole)
	group.PATCH("/editgroupvisibility", httpHandler.EditGroupVisibility)
	group.PATCH("/editgroupdescription", httpHandler.EditGroupDescription)
	group.DELETE("/group", httpHandler.DeleteGroup)
	group.DELETE("/groupmember", httpHandler.DeleteGroupMember)
	group.GET("/search", httpHandler.SearchGroupByGroupName)
	group.GET("/acceptorrequests", httpHandler.GetAcceptorGroupRequests)
	group.GET("/acceptorrequestscount", httpHandler.GetAcceptorGroupRequestsCount)

	// Post Mock data
	post.POST("/mock", httpHandler.PostMockData)

	// Post Endpoints
	post.POST("/post", httpHandler.CreatePost)
	post.GET("/post", httpHandler.GetPostByPostId)
	post.GET("/groupposts", httpHandler.GetPostsByGroupId)
	post.GET("/userposts", httpHandler.GetPostsByUserId)
	post.GET("/groupuserposts", httpHandler.GetPostsByGroupAndMemberId)
	post.PATCH("/post", httpHandler.EditPost)
	post.DELETE("/post", httpHandler.DeletePost)
	post.GET("/ongoingfeed", httpHandler.GetPostsForOngoingFeedByMemberId)
	post.GET("/followingfeed", httpHandler.GetPostsForFollowingFeedByMemberId)

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

	// Tag Endpoints
	tag.POST("/tag", httpHandler.CreateTag)
	tag.GET("/alltags", httpHandler.GetAvailableTags)
	tag.GET("/groupswithtag", httpHandler.GetGroupsByTagID)
	tag.GET("/tagbycategory", httpHandler.GetTagByCategory)
	tag.GET("/tagofgroup", httpHandler.GetTagByGroupId)
	tag.GET("/groupswithcategory", httpHandler.GetGroupsByCategoryID)
	tag.PATCH("/edittagname", httpHandler.EditTagName)
	tag.PATCH("/edittagcategory", httpHandler.EditTagCategory)
	tag.PATCH("/edittagicon", httpHandler.EditTagIcon)
	tag.DELETE("/tag", httpHandler.DeleteTag)
}
