package server

import (
	"log"

	"github.com/Win-TS/gleam-backend.git/modules/user/userHandler"
	userPb "github.com/Win-TS/gleam-backend.git/modules/user/userPb"
	"github.com/Win-TS/gleam-backend.git/modules/user/userUsecase"
	userdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/userdb/sqlc"
	"github.com/Win-TS/gleam-backend.git/pkg/grpcconn"
)

func (s *server) userService() {
	postgresDB, ok := s.db.(userdb.Store)
	if !ok {
		log.Fatal("Unsupported database type")
		return
	}
	usecase := userUsecase.NewUserUsecase(postgresDB, s.storage)
	httpHandler := userHandler.NewUserHttpHandler(s.cfg, usecase)
	grpcHandler := userHandler.NewUserGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(s.cfg, s.cfg.Grpc.UserUrl)
		userPb.RegisterUserGrpcServiceServer(grpcServer, grpcHandler)
		log.Printf("User gRPC server listening on %s", s.cfg.Grpc.UserUrl)
		grpcServer.Serve(lis)
	}()

	_ = grpcHandler

	user := s.app.Group("/user_v1")

	// Health Check
	user.GET("", s.healthCheckService)
	// Fill mock data
	user.POST("/mock", httpHandler.UserMockData)

	user.POST("/createuser", httpHandler.RegisterNewUser)
	user.GET("/userprofile", httpHandler.GetUserProfile)
	user.GET("/userprofilebyemail", httpHandler.GetUserProfileByEmail)
	user.GET("/userprofilebyusername", httpHandler.GetUserProfileByUsername)
	user.POST("/uploaduserphoto", httpHandler.UploadProfilePhoto)
	user.GET("/userinfo", httpHandler.GetUserInfo)
	user.GET("/userinfobyemail", httpHandler.GetUserInfoByEmail)
	user.GET("/userinfobyusername", httpHandler.GetUserInfoByUsername)
	user.PATCH("/editusername", httpHandler.EditUsername)
	user.PATCH("/changephoneno", httpHandler.EditPhoneNumber)
	user.PATCH("/editname", httpHandler.EditName)
	user.DELETE("/deleteuser", httpHandler.DeleteUser)
	user.PATCH("/editphoto", httpHandler.EditUserPhoto)

	friend := s.app.Group("/friend_v1")

	friend.GET("/", httpHandler.FriendInfo)
	friend.GET("/list", httpHandler.FriendListById)
	friend.GET("/count", httpHandler.FriendsCount)
	friend.GET("/requested", httpHandler.FriendsRequestedList)
	friend.GET("/pending", httpHandler.FriendsPendingList)
	friend.POST("/add", httpHandler.AddFriend)
	friend.PATCH("/accept", httpHandler.FriendAccept)
}
