package service

import (
	"context"

	"rpm/microservices/api/graph/model"
	proto "rpm/microservices/core/proto"
)

const (
	// Catalog service
	CatalogService = "go.micro.srv.catalog"
)

// GetProductCatalog ...
func (s *service) GetProductCatalog(ctx context.Context, obj *model.AbstractModel) (*model.Catalogs, error) {
	response, err := s.catalog.GetProductCatalog(ctx, &proto.CoreRequest{})
	var result *model.Catalogs
	if err == nil {
		result = convertCatalogResponse(response)
	}
	return result, err
}

func convertCatalogResponse(productCatalog *proto.ProductCatalogResponse) *model.Catalogs {
	var products []*model.Product

	for _, product := range productCatalog.Catalogs.Products {
		products = append(products, &model.Product{
			ProductName:        product.ProductName,
			ProductDescription: product.ProductDescription,
			Price:              int(product.Price),
			Stock:              int(product.Stock),
		})
	}

	return &model.Catalogs{
		Catalogs: products,
	}
}
