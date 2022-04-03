package handler

import (
	"context"

	proto "rpm/microservices/core/proto"
)

// Cart ...
type Cart struct{}

// GetShoppingCartByUserId ..
func (h *handler) GetProductCatalog(ctx context.Context, req *proto.CoreRequest, rsp *proto.ProductCatalogResponse) error {
	result, err := h.useCase.GetProductCatalog()

	rsp.Catalogs = result.Catalogs

	return err
}
