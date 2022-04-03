package main

import (
	"sync"

	"rpm/microservices/core/environ"
	"rpm/microservices/core/utils"

	core "rpm/microservices/core/proto"
	"rpm/microservices/services/cart/handler"

	"github.com/joho/godotenv"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/service"
	grpc "github.com/micro/go-micro/v2/service/grpc"
	kubernetes "github.com/micro/go-plugins/registry/kubernetes/v2"
)

var (
	reportServer *Server
	once         sync.Once
)

type Server struct {
	utils utils.Utils
	env   environ.Environ
}

func newServer() *Server {
	once.Do(func() {
		reportServer = &Server{
			utils: utils.NewUtils(),
			env:   environ.NewEnviron(),
		}
	})
	return reportServer
}

// init is invoked before main()
func init() {
	srv := newServer()
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		logger.Info("No .env file found ..")
	}

	logger.Info("Data Source : " + srv.utils.GetDatasourceInfo())
}

func main() {
	srv := newServer()
	// New Service
	service := grpc.NewService(getServerOptions(srv)...)

	// Initialise service
	service.Init()
	// Register Handler
	core.RegisterShoppingCartHandler(service.Server(), handler.NewHandler())

	// Run service
	if err := service.Run(); err != nil {
		logger.Error("Error Stating Server", err)
	}
}

func getServerOptions(srv *Server) []service.Option {

	options := []service.Option{
		service.Name("go.micro.srv.cart"),
		service.Version("latest"),
	}

	if srv.env.IsUseKubernetes() {
		options = append(options, service.Registry(kubernetes.NewRegistry()))
	}

	return options
}
