package repository

import (
	proto "rpm/microservices/core/proto"
)

// GetUserInfoByID ..
func (ar *abstractRepository) GetUserInfoByID(userId string) (*proto.UserInfoResponse, error) {
	var (
		result = new(proto.UserInfoResponse)
	)

	records, err := ar.GetUserInfoByIDQuery(userId)
	if err != nil {
		return nil, err
	}

	if len(records) == 2 {
		result = &proto.UserInfoResponse{
			UserID:       records[1][0],
			UserName:     records[1][1],
			UserLocation: records[1][2],
		}

	}
	return result, nil
}

// ListOfUser ..
func (ar *abstractRepository) ListOfUser() (*proto.ListUserResponse, error) {
	var (
		result = new(proto.ListUserResponse)
	)

	records, err := ar.ListOfUserQuery()
	if err != nil {
		return nil, err
	}
	var users []*proto.UserInfoResponse
	for i := range records {
		if i > 0 {

			userData := &proto.UserInfoResponse{
				UserID:       records[i][0],
				UserName:     records[i][1],
				UserLocation: records[i][2],
			}
			users = append(users, userData)
		}
	}
	result.Users = users
	return result, nil
}
