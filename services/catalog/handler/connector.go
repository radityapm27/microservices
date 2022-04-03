package handler

import (
	"sync"

	"rpm/microservices/core/config"
	"rpm/microservices/core/environ"
	"rpm/microservices/core/utils"
	"rpm/microservices/services/catalog/repository"
	"rpm/microservices/services/catalog/usecase"
)

const (
	// RepoType ...
	RepoType = "REPO_TYPE"
	// ElasticSearch ...
	ElasticSearch = "elastic"
)

var catalog usecase.CatalogUseCase
var oneCatalog sync.Once

// GetUsecaseCatalog ...
func GetUsecaseCatalog() usecase.CatalogUseCase {
	oneCatalog.Do(func() {
		catalog = usecase.NewCatalogUseCase(
			GetCatalogRepository(),
			config.NewConfig(),
			utils.NewUtils(),
			environ.NewEnviron(),
		)
	})
	return catalog
}

var repo repository.AbstractRepository
var oneRepo sync.Once

// GetCatalogRepository ...
func GetCatalogRepository() repository.AbstractRepository {
	oneRepo.Do(func() {
		repo = repository.NewRepository(utils.NewUtils())
	})
	return repo
}
