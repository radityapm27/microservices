package handler

import (
	"context"

	proto "rpm/microservices/core/proto"
	"rpm/microservices/services/catalog/usecase"
)

type handler struct {
	useCase usecase.CatalogUseCase
}

// Handler ...
type Handler interface {
	GetProductCatalog(ctx context.Context, req *proto.CoreRequest, rsp *proto.ProductCatalogResponse) error
}

func NewHandler() *handler {
	return &handler{useCase: GetUsecaseCatalog()}
}
