package server

import (
	"log"

	"github.com/Win-TS/gleam-backend.git/modules/user/userHandler"
	"github.com/Win-TS/gleam-backend.git/modules/user/userUsecase"
	userdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/userdb/sqlc"
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

	_ = grpcHandler

	user := s.app.Group("/user_v1")

	// Health Check
	user.GET("", s.healthCheckService)
	// Fill mock data
	user.POST("/mock", httpHandler.UserMockData)

	user.POST("/createuser", httpHandler.RegisterNewUser)
	user.GET("/userprofile", httpHandler.GetUserProfile)
	user.POST("/uploaduserphoto", httpHandler.UploadProfilePhoto)
	user.GET("/userinfo", httpHandler.GetUserInfo)
	user.PATCH("/editusername", httpHandler.EditUsername)
	user.PATCH("/changephoneno", httpHandler.EditPhoneNumber)
	user.DELETE("/deleteuser", httpHandler.DeleteUser)

	friend := s.app.Group("/friend_v1")

	friend.GET("/", httpHandler.FriendInfo)
	friend.GET("/list", httpHandler.FriendListById)
	friend.GET("/count", httpHandler.FriendsCount)
	friend.GET("/pending", httpHandler.FriendsPendingList)
	friend.POST("/add", httpHandler.AddFriend)
	friend.PATCH("/accept", httpHandler.FriendAccept)
}
