package usecase

import (
	"rpm/microservices/core/config"
	"rpm/microservices/core/environ"
	proto "rpm/microservices/core/proto"
	utils "rpm/microservices/core/utils"
	repository "rpm/microservices/services/catalog/repository"
)

type catalogUseCase struct {
	conf       config.Config
	repository repository.AbstractRepository
	utils      utils.Utils
	env        environ.Environ
}

// NewCatalogUseCase func
func NewCatalogUseCase(repo repository.AbstractRepository, cnf config.Config, utils utils.Utils, env environ.Environ) CatalogUseCase {
	return &catalogUseCase{
		conf:       cnf,
		repository: repo,
		utils:      utils,
		env:        env,
	}
}

// CatalogUseCase ...
type CatalogUseCase interface {
	GetProductCatalog() (*proto.ProductCatalogResponse, error)
}
