package service

import (
	"context"
	"rpm/microservices/api/graph/model"
	serviceModel "rpm/microservices/core/proto"
	"rpm/microservices/core/utils"
)

// Service ...
type service struct {
	shoppingCart serviceModel.ShoppingCartService
	catalog      serviceModel.ProductCatalogService
	user         serviceModel.UserService
	utils        utils.Utils
}

type Service interface {
	// Cart ...
	GetShoppingCartByUserID(ctx context.Context, userID string) (*model.Carts, error)
	// Catalog ...
	GetProductCatalog(ctx context.Context, obj *model.AbstractModel) (*model.Catalogs, error)
	// User ...
	GetUserInfoByID(ctx context.Context, obj *model.AbstractModel, userID string) (*model.UserInfo, error)
	ListOfUser(ctx context.Context, obj *model.AbstractModel) (*model.ListUser, error)
}

// New ...
func New(shoppingCartService serviceModel.ShoppingCartService, productCatalogService serviceModel.ProductCatalogService, userService serviceModel.UserService) Service {
	return &service{
		shoppingCart: shoppingCartService,
		catalog:      productCatalogService,
		user:         userService,
		utils:        utils.NewUtils(),
	}
}
