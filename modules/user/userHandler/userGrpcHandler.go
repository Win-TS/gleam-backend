package userHandler

import (
	"context"
	"errors"

	userPb "github.com/Win-TS/gleam-backend.git/modules/user/userPb"
	"github.com/Win-TS/gleam-backend.git/modules/user/userUsecase"
)

type (
	userGrpcHandler struct {
		userPb.UnimplementedUserGrpcServiceServer
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserGrpcHandler(userUsecase userUsecase.UserUsecaseService) *userGrpcHandler {
	return &userGrpcHandler{userUsecase: userUsecase}
}

func (g *userGrpcHandler) SearchUser(ctx context.Context, req *userPb.SearchUserReq) (*userPb.SearchUserRes, error) {
	_, err := g.userUsecase.GetUserInfo(ctx, int(req.UserId))
	if err != nil {
		return &userPb.SearchUserRes{
			UserId: req.UserId,
			Valid:  false,
		}, errors.New("error: userId not found")
	}

	return &userPb.SearchUserRes{
		UserId: req.UserId,
		Valid:  true,
	}, nil
}

func (g *userGrpcHandler) GetUserProfile(ctx context.Context, req *userPb.GetUserProfileReq) (*userPb.GetUserProfileRes, error) {
	user, err := g.userUsecase.GetUserInfo(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}

	return &userPb.GetUserProfileRes{
		UserId:    int32(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Photourl:  user.Photourl.String,
	}, nil
}

func (g *userGrpcHandler) GetBatchUserProfiles(ctx context.Context, req *userPb.GetBatchUserProfileReq) (*userPb.GetBatchUserProfileRes, error) {
	users, err := g.userUsecase.GetBatchUserProfiles(ctx, req.UserIds)
	if err != nil {
		return nil, err
	}

	var userProfileRes []*userPb.GetBatchUserProfileRes_UserProfile
	for _, user := range users {
		userProfileRes = append(userProfileRes, &userPb.GetBatchUserProfileRes_UserProfile{
			UserId:    int32(user.ID),
			Username:  user.Username,
			Email:     user.Email,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Photourl:  user.Photourl.String,
		})
	}

	return &userPb.GetBatchUserProfileRes{
		UserProfiles: userProfileRes,
	}, nil
}

func (g *userGrpcHandler) GetUserFriends(ctx context.Context, req *userPb.GetUserFriendsReq) (*userPb.GetUserFriendsRes, error) {
	friends, err := g.userUsecase.FriendListByIdNoPaginate(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}

	var friendsRes []*userPb.GetUserFriendsRes_Friend
	for _, friend := range friends {
		friendsRes = append(friendsRes, &userPb.GetUserFriendsRes_Friend{
			UserId:    int32(friend.ID),
			Username:  friend.Username,
			Email:     friend.Email,
			Firstname: friend.Firstname,
			Lastname:  friend.Lastname,
			Photourl:  friend.Photourl.String,
		})
	}

	return &userPb.GetUserFriendsRes{
		Friends: friendsRes,
	}, nil
}