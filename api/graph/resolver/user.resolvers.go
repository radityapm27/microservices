package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"rpm/microservices/api/graph"
	"rpm/microservices/api/graph/model"
)

func (r *userResolver) GetUserInfoByID(ctx context.Context, obj *model.AbstractModel, userID string) (*model.UserInfo, error) {
	return r.service.GetUserInfoByID(ctx, obj, userID)
}

func (r *userResolver) ListOfUser(ctx context.Context, obj *model.AbstractModel) (*model.ListUser, error) {
	return r.service.ListOfUser(ctx, obj)
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
