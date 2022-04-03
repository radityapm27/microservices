package usecase

import (
	proto "rpm/microservices/core/proto"
)

// GetUserInfoByID ..
func (userUseCase *userUseCase) GetUserInfoByID(userID string) (*proto.UserInfoResponse, error) {
	return userUseCase.repository.GetUserInfoByID(userID)
}

// GetUserInfoByID ..
func (userUseCase *userUseCase) ListOfUser() (*proto.ListUserResponse, error) {
	return userUseCase.repository.ListOfUser()
}
