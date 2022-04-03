package service

import (
	"context"
	"testing"

	proto "rpm/microservices/core/proto"

	"github.com/stretchr/testify/assert"
)

func TestService_GetShoppingCartByUserID(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		s    = new(service)
	)

	t.Run("success", func(t *testing.T) {
		var (
			response = &proto.ShoppingCartResponse{

				Cart: &proto.Cart{
					UserInfo: &proto.UserInfoResponse{
						UserID:       "test@gmail.com",
						UserName:     "User Name",
						UserLocation: "User Location",
					},
					ProductCart: []*proto.ProductCart{
						&proto.ProductCart{
							Product: &proto.Product{
								ProductName:        "Product Name 1",
								ProductDescription: "Product Description 1",
								Price:              10000,
								Stock:              100,
							},
							Quantity: 10,
						},
						&proto.ProductCart{
							Product: &proto.Product{
								ProductName:        "Product Name 2",
								ProductDescription: "Product Description 2",
								Price:              20000,
								Stock:              200,
							},
							Quantity: 20,
						},
					},
				},
			}
			userID = "test@gmail.com"
			err    error
			ctx    = context.Background()
		)

		svc := resetShoppingCart(s)
		svc.On("GetShoppingCartByUserId", ctx, &proto.CoreRequest{
			UserID: userID,
		}).Return(response, err)
		resp, respErr := s.GetShoppingCartByUserID(ctx, userID)

		svc.AssertExpectations(t)

		test.NotNil(resp)
		test.Nil(respErr)
	})
}
