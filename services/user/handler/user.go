package handler

import (
	"context"

	proto "rpm/microservices/core/proto"
)

// Cart ...
type Cart struct{}

// GetUserById ..
func (h *handler) GetUserById(ctx context.Context, request *proto.CoreRequest, response *proto.UserInfoResponse) error {
	result, err := h.useCase.GetUserInfoByID(request.UserID)

	response.UserID = result.UserID
	response.UserLocation = result.UserLocation
	response.UserName = result.UserName

	return err
}

func (h *handler) ListOfUser(ctx context.Context, request *proto.CoreRequest, response *proto.ListUserResponse) error {
	result, err := h.useCase.ListOfUser()
	response.Users = result.Users

	return err
}
