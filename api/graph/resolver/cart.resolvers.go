package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"rpm/microservices/api/graph"
	"rpm/microservices/api/graph/model"
)

func (r *shoppingCartResolver) GetShoppingCartByUserID(ctx context.Context, obj *model.AbstractModel, userID string) (*model.Carts, error) {
	return r.service.GetShoppingCartByUserID(ctx, userID)
}

// ShoppingCart returns graph.ShoppingCartResolver implementation.
func (r *Resolver) ShoppingCart() graph.ShoppingCartResolver { return &shoppingCartResolver{r} }

type shoppingCartResolver struct{ *Resolver }
