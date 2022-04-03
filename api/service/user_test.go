package service

import (
	"context"
	"testing"

	"rpm/microservices/api/graph/model"
	proto "rpm/microservices/core/proto"

	"github.com/stretchr/testify/assert"
)

func TestService_GetUserInfoByID(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		s    = new(service)
	)

	t.Run("success", func(t *testing.T) {
		var (
			response = &proto.UserInfoResponse{
				UserID:       "test@gmail.com",
				UserName:     "User Name",
				UserLocation: "User Location",
			}
			err error
			ctx = context.Background()
			obj = &model.AbstractModel{}
		)

		svc := resetUser(s)
		svc.On("GetUserById", ctx, &proto.CoreRequest{
			UserID: response.UserID,
		}).Return(response, err)
		resp, respErr := s.GetUserInfoByID(ctx, obj, response.UserID)

		svc.AssertExpectations(t)

		test.NotNil(resp)
		test.Nil(respErr)
	})
}

func TestService_ListOfUser(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		s    = new(service)
	)

	t.Run("success", func(t *testing.T) {
		var (
			response = &proto.ListUserResponse{
				Users: []*proto.UserInfoResponse{
					&proto.UserInfoResponse{
						UserID:       "test2@gmail.com",
						UserName:     "User Name 1",
						UserLocation: "User Location 1",
					},
					&proto.UserInfoResponse{
						UserID:       "test2@gmail.com",
						UserName:     "User Name 2",
						UserLocation: "User Location 2",
					},
				},
			}
			err error
			ctx = context.Background()
			obj = &model.AbstractModel{}
		)

		svc := resetUser(s)
		svc.On("ListOfUser", ctx, &proto.CoreRequest{}).Return(response, err)
		resp, respErr := s.ListOfUser(ctx, obj)

		svc.AssertExpectations(t)

		test.NotNil(resp)
		test.Nil(respErr)
	})
}
