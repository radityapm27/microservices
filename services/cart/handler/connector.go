package handler

import (
	"sync"

	"rpm/microservices/core/config"
	"rpm/microservices/core/environ"
	"rpm/microservices/core/utils"
	"rpm/microservices/services/cart/repository"
	"rpm/microservices/services/cart/usecase"
)

const (
	// RepoType ...
	RepoType = "REPO_TYPE"
	// ElasticSearch ...
	ElasticSearch = "elastic"
)

var cart usecase.CartUseCase
var oneCart sync.Once

// GetUsecaseCart ...
func GetUsecaseCart() usecase.CartUseCase {
	oneCart.Do(func() {
		cart = usecase.NewCartUseCase(
			GetCartRepository(),
			config.NewConfig(),
			utils.NewUtils(),
			environ.NewEnviron(),
		)
	})
	return cart
}

var repo repository.AbstractRepository
var oneRepo sync.Once

// GetCartRepository ...
func GetCartRepository() repository.AbstractRepository {
	oneRepo.Do(func() {
		repo = repository.NewRepository(utils.NewUtils())
	})
	return repo
}
