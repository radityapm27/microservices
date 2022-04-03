package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"rpm/microservices/api/graph"
	"rpm/microservices/api/graph/model"
)

func (r *queryResolver) ShoppingCart(ctx context.Context) (*model.AbstractModel, error) {
	return &model.AbstractModel{}, nil
}

func (r *queryResolver) ProductCatalog(ctx context.Context) (*model.AbstractModel, error) {
	return &model.AbstractModel{}, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.AbstractModel, error) {
	return &model.AbstractModel{}, nil
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
