package usecase

import (
	"rpm/microservices/core/config"
	"rpm/microservices/core/environ"
	proto "rpm/microservices/core/proto"
	utils "rpm/microservices/core/utils"
	repository "rpm/microservices/services/cart/repository"
)

type cartUseCase struct {
	conf       config.Config
	repository repository.AbstractRepository
	utils      utils.Utils
	env        environ.Environ
}

// NewCartUseCase func
func NewCartUseCase(repo repository.AbstractRepository, cnf config.Config, utils utils.Utils, env environ.Environ) CartUseCase {
	return &cartUseCase{
		conf:       cnf,
		repository: repo,
		utils:      utils,
		env:        env,
	}
}

// CartUseCase ...
type CartUseCase interface {
	GetShoppingCartByUserId(userID string) (*proto.ShoppingCartResponse, error)
}
