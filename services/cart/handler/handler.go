package handler

import (
	"context"

	proto "rpm/microservices/core/proto"
	"rpm/microservices/services/cart/usecase"
)

type handler struct {
	useCase usecase.CartUseCase
}

// Handler ...
type Handler interface {
	GetShoppingCartByUserId(ctx context.Context, req *proto.CoreRequest, rsp *proto.ShoppingCartResponse) error
}

func NewHandler() *handler {
	return &handler{useCase: GetUsecaseCart()}
}
