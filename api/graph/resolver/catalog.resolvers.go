package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"rpm/microservices/api/graph"
	"rpm/microservices/api/graph/model"
)

func (r *productCatalogResolver) GetProductCatalog(ctx context.Context, obj *model.AbstractModel) (*model.Catalogs, error) {
	return r.service.GetProductCatalog(ctx, obj)
}

// ProductCatalog returns graph.ProductCatalogResolver implementation.
func (r *Resolver) ProductCatalog() graph.ProductCatalogResolver { return &productCatalogResolver{r} }

type productCatalogResolver struct{ *Resolver }
