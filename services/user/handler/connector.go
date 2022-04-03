package handler

import (
	"sync"

	"rpm/microservices/core/config"
	"rpm/microservices/core/environ"
	"rpm/microservices/core/utils"
	"rpm/microservices/services/user/repository"
	"rpm/microservices/services/user/usecase"
)

const (
	// RepoType ...
	RepoType = "REPO_TYPE"
	// ElasticSearch ...
	ElasticSearch = "elastic"
)

var user usecase.UserUseCase
var oneCart sync.Once

// GetUsecaseUser ...
func GetUsecaseUser() usecase.UserUseCase {
	oneCart.Do(func() {
		user = usecase.NewUserUseCase(
			GetUserRepository(),
			config.NewConfig(),
			utils.NewUtils(),
			environ.NewEnviron(),
		)
	})
	return user
}

var repo repository.AbstractRepository
var oneRepo sync.Once

// GetUserRepository ...
func GetUserRepository() repository.AbstractRepository {
	oneRepo.Do(func() {
		repo = repository.NewRepository(utils.NewUtils())
	})
	return repo
}
