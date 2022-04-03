package client

import (
	"sync"
	"time"

	"rpm/microservices/core/environ"
	"rpm/microservices/core/utils"

	"github.com/micro/go-micro/v2/client"
	grpcClient "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	configFile "rpm/microservices/core/config"
	service "rpm/microservices/core/proto"
)

const (
	// Cart Service ...
	CartService = "go.micro.srv.cart"
	// Catalog Service ...
	CatalogService = "go.micro.srv.catalog"
	// User Service ...
	UserService = "go.micro.srv.user"
)

var (
	//Cart
	cartService     service.ShoppingCartService
	onceCartService sync.Once
	//Catalog
	catalogService     service.ProductCatalogService
	onceCatalogService sync.Once
	//User
	userService     service.UserService
	onceUserService sync.Once

	oneMclient sync.Once
	mClient    *microClient
)

type microClient struct {
	cfg       configFile.Config
	coreUtils utils.Utils
	env       environ.Environ
}

type MicroClient interface {
	GetMicroServiceClientCart() service.ShoppingCartService
	GetMicroServiceClientCatalog() service.ProductCatalogService
	GetMicroServiceClientUser() service.UserService
}

func NewMicroClient() MicroClient {
	oneMclient.Do(func() {
		mClient = &microClient{
			cfg:       configFile.NewConfig(),
			coreUtils: utils.NewUtils(),
			env:       environ.NewEnviron(),
		}
	})
	return mClient
}

// GetMicroServiceClientCart ...
func (mc *microClient) GetMicroServiceClientCart() service.ShoppingCartService {
	onceCartService.Do(func() {
		options := getClientOptions(mc)
		cartService = service.NewShoppingCartService(CartService, grpcClient.NewClient(options...))
	})
	return cartService
}

// GetMicroServiceClientCatalog ...
func (mc *microClient) GetMicroServiceClientCatalog() service.ProductCatalogService {

	onceCatalogService.Do(func() {
		options := getClientOptions(mc)
		catalogService = service.NewProductCatalogService(CatalogService, grpcClient.NewClient(options...))
	})
	return catalogService
}

// GetMicroServiceClientUser ...
func (mc *microClient) GetMicroServiceClientUser() service.UserService {

	onceUserService.Do(func() {
		options := getClientOptions(mc)
		userService = service.NewUserService(UserService, grpcClient.NewClient(options...))
	})
	return userService
}

func getClientOptions(mc *microClient) []client.Option {
	config := mc.cfg.GetApplicationConfig()
	options := []client.Option{
		client.RequestTimeout(config.GRPCTimeout * time.Second),
		client.Retries(0),
	}

	if mc.env.GetMicroRegistry() == "kubernetes" {
		options = append(options, client.Registry(kubernetes.NewRegistry()))
	}

	return options
}
