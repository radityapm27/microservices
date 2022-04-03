package handler

import (
	"context"

	proto "rpm/microservices/core/proto"
)

// Cart ...
type Cart struct{}

// GetShoppingCartByUserId ..
func (h *handler) GetShoppingCartByUserId(ctx context.Context, request *proto.CoreRequest, response *proto.ShoppingCartResponse) error {
	result, err := h.useCase.GetShoppingCartByUserId(request.UserID)

	response.Cart = result.Cart

	return err
}
