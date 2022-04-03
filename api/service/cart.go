package service

import (
	"context"

	"rpm/microservices/api/graph/model"
	proto "rpm/microservices/core/proto"
)

const (
	// Cart service
	CartService = "go.micro.srv.cart"
)

// GetShoppingCartByUserID ...
func (s *service) GetShoppingCartByUserID(ctx context.Context, userID string) (*model.Carts, error) {
	response, err := s.shoppingCart.GetShoppingCartByUserId(ctx, &proto.CoreRequest{
		UserID: userID,
	})

	var result *model.Carts
	if err == nil {
		result = convertCartsResponse(response)
	}
	return result, err
}

func convertCartsResponse(shoppingCart *proto.ShoppingCartResponse) *model.Carts {
	var productCart []*model.ProductCart
	userInfo := &proto.UserInfoResponse{}

	if shoppingCart != nil {
		userInfo = shoppingCart.Cart.UserInfo
	}

	for _, cart := range shoppingCart.Cart.ProductCart {
		productCart = append(productCart, &model.ProductCart{
			Product: &model.Product{
				ProductName:        cart.Product.ProductName,
				ProductDescription: cart.Product.ProductDescription,
				Price:              int(cart.Product.Price),
				Stock:              int(cart.Product.Stock),
			},
			Quantity: int(cart.Quantity),
		})
	}

	return &model.Carts{
		UserInfo: &model.UserInfo{
			UserID:       userInfo.UserID,
			UserName:     userInfo.UserName,
			UserLocation: userInfo.UserLocation,
		},
		Cart: productCart,
	}
}
