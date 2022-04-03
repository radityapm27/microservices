package repository

import (
	proto "rpm/microservices/core/proto"
	utils "rpm/microservices/core/utils"
	"rpm/microservices/services/user/repository/postgres"

	"github.com/patrickmn/go-cache"
)

const (
	here = "repo.abstracuser."
)

// Repository interface
type Repository interface {
	GetUserInfoByIDQuery(userID string) ([][]string, error)
	ListOfUserQuery() ([][]string, error)
}

// AbstractRepository interface
type AbstractRepository interface {
	GetUserInfoByID(userID string) (*proto.UserInfoResponse, error)
	ListOfUser() (*proto.ListUserResponse, error)
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
