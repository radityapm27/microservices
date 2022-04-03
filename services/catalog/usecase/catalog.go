package usecase

import (
	proto "rpm/microservices/core/proto"
)

// GetProductCatalog ..
func (catalogUseCase *catalogUseCase) GetProductCatalog() (*proto.ProductCatalogResponse, error) {

	return catalogUseCase.repository.GetProductCatalog()
}
