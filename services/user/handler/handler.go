package handler

import (
	"context"

	proto "rpm/microservices/core/proto"
	"rpm/microservices/services/user/usecase"
)

type handler struct {
	useCase usecase.UserUseCase
}

// Handler ...
type Handler interface {
	GetUserById(ctx context.Context, req *proto.CoreRequest, rsp *proto.UserInfoResponse) error
	ListOfUser(ctx context.Context, request *proto.CoreRequest, response *proto.ListUserResponse) error
}

func NewHandler() *handler {
	return &handler{useCase: GetUsecaseUser()}
}
