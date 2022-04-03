package usecase

import (
	"rpm/microservices/core/config"
	"rpm/microservices/core/environ"
	proto "rpm/microservices/core/proto"
	utils "rpm/microservices/core/utils"
	repository "rpm/microservices/services/user/repository"
)

type userUseCase struct {
	conf       config.Config
	repository repository.AbstractRepository
	utils      utils.Utils
	env        environ.Environ
}

// NewUserUseCase func
func NewUserUseCase(repo repository.AbstractRepository, cnf config.Config, utils utils.Utils, env environ.Environ) UserUseCase {
	return &userUseCase{
		conf:       cnf,
		repository: repo,
		utils:      utils,
		env:        env,
	}
}

// UserUseCase ...
type UserUseCase interface {
	GetUserInfoByID(userID string) (*proto.UserInfoResponse, error)
	ListOfUser() (*proto.ListUserResponse, error)
}
