package usecase

import (
	proto "rpm/microservices/core/proto"
)

// GetShoppingCartByUserId ..
func (cartUseCase *cartUseCase) GetShoppingCartByUserId(userID string) (*proto.ShoppingCartResponse, error) {

	return cartUseCase.repository.GetShoppingCartByUserId(userID)
}
