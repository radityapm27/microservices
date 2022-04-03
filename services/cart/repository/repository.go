package repository

import (
	proto "rpm/microservices/core/proto"
	utils "rpm/microservices/core/utils"
	"rpm/microservices/services/cart/repository/postgres"

	"github.com/patrickmn/go-cache"
)

const (
	here = "repo.abstractreport."
)

// Repository interface
type Repository interface {
	GetShoppingCartByUserIdQuery(userID string) ([][]string, error)
}

// AbstractRepository interface
type AbstractRepository interface {
	GetShoppingCartByUserId(userID string) (*proto.ShoppingCartResponse, error)
}

type abstractRepository struct {
	Repository
	Cache *cache.Cache
	utils utils.Utils
}

// NewRepository ...
func NewRepository(coreUtils utils.Utils) AbstractRepository {
	newRepo := new(abstractRepository)
	newRepo.Repository = nil

	if coreUtils.GetDatasourceInfo() == utils.Postgres {
		newRepo.Repository = &postgres.Repository{
			Connection: coreUtils.GetClientHandler(),
			CoreUtils:  coreUtils,
		}
	}

	newRepo.Cache = coreUtils.GetCacheHandler()
	newRepo.utils = coreUtils

	return newRepo
}
