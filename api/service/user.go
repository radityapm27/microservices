package service

import (
	"context"
	"fmt"

	"rpm/microservices/api/graph/model"
	proto "rpm/microservices/core/proto"
)

const (
	// User service
	UserService = "go.micro.srv.user"
)

// GetUserInfoByID ...
func (s *service) GetUserInfoByID(ctx context.Context, obj *model.AbstractModel, userID string) (*model.UserInfo, error) {
	response, err := s.user.GetUserById(ctx, &proto.CoreRequest{
		UserID: userID,
	})
	fmt.Println(response)
	var result *model.UserInfo
	if err == nil {
		result = convertUserInfoResponse(response)
	}

	return result, err
}

// ListOfUser ...
func (s *service) ListOfUser(ctx context.Context, obj *model.AbstractModel) (*model.ListUser, error) {
	response, err := s.user.ListOfUser(ctx, &proto.CoreRequest{})
	var result *model.ListUser
	if err == nil {
		result = convertListUserResponse(response)
	}
	return result, err
}

func convertUserInfoResponse(resp *proto.UserInfoResponse) *model.UserInfo {
	userInfo := &proto.UserInfoResponse{}

	if resp != nil {
		userInfo = resp
	}

	return &model.UserInfo{
		UserID:       userInfo.UserID,
		UserName:     userInfo.UserName,
		UserLocation: userInfo.UserLocation,
	}
}

func convertListUserResponse(users *proto.ListUserResponse) *model.ListUser {
	var userList []*model.UserInfo

	for _, userInfo := range users.Users {
		userList = append(userList, &model.UserInfo{
			UserID:       userInfo.UserID,
			UserName:     userInfo.UserName,
			UserLocation: userInfo.UserLocation,
		})
	}

	return &model.ListUser{
		Users: userList,
	}
}
