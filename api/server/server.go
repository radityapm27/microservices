package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"rpm/microservices/api/graph"
	"rpm/microservices/api/graph/resolver"
	"rpm/microservices/api/middleware"
	"rpm/microservices/api/service"
	coreclient "rpm/microservices/core/client"
	"rpm/microservices/core/environ"
	"rpm/microservices/core/utils"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/micro/go-micro/v2/logger"
	"github.com/rs/cors"
)

const defaultPort = "8080"

// Server as API server
type Server struct {
	instance *http.Server
	port     string
	useTLS   bool
	tlsCert  string
	tlsKey   string
	logger   *log.Logger
	utils    utils.Utils
	env      environ.Environ
	mdl      middleware.Middleware
}

func init() {
	// init environment
	godotenv.Load()
}

func main() {
	srv := createServer()
	logger.Info("Server is starting...")

	ctx := srv.handleShutdown()
	srv.start()

	<-ctx.Done()
	logger.Info("Server stopped...")
}

func createServer() *Server {
	srv := new(Server)
	srv.utils = utils.NewUtils()
	srv.env = environ.NewEnviron()
	srv.mdl = middleware.NewMiddleware()
	srv.port = defaultPort

	srv.logger = log.New(os.Stdout, "[http] ", log.LstdFlags)
	srv.instance = &http.Server{
		Addr:    ":" + srv.port,
		Handler: buildRouter(srv.env, srv.mdl),
	}

	return srv
}

func buildRouter(env environ.Environ, mdl middleware.Middleware) http.Handler {
	router := chi.NewRouter()

	options := []handler.Option{
		handler.WebsocketUpgrader(
			websocket.Upgrader{
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
			}),
		handler.CacheSize(0),
	}

	router.Use(
		getCorsAPI(env).Handler,
		mdl.GetMiddleware,
	)

	router.Handle(
		"/",
		handler.Playground("GraphQL playground", "/query"),
	)

	graphConfig := graph.Config{Resolvers: buildResolver(coreclient.NewMicroClient())}

	router.Handle(
		"/query",
		handler.GraphQL(
			graph.NewExecutableSchema(graphConfig),
			options...,
		),
	)

	return router
}

func buildResolver(mc coreclient.MicroClient) *resolver.Resolver {
	reslv := resolver.New(service.New(
		mc.GetMicroServiceClientCart(), mc.GetMicroServiceClientCatalog(),
		mc.GetMicroServiceClientUser(),
	))
	return reslv
}

func getCorsAPI(env environ.Environ) *cors.Cors {
	allowedOrigins := []string{"*"}
	if origins := env.GetAllowedOrigins(); origins != "" {
		allowedOrigins = strings.Split(origins, "|")
	}

	return cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            false,
	})
}

func (s *Server) start() {
	logger.Info(fmt.Sprintf("Server is ready to handle requests at %s", s.instance.Addr))
	logger.Info(fmt.Sprintf("Connect to http://localhost:%s/ for GraphQL playground", s.port))

	var err error
	if s.useTLS {
		err = s.instance.ListenAndServeTLS(s.tlsCert, s.tlsKey)
	} else {
		err = s.instance.ListenAndServe()
	}

	if err != http.ErrServerClosed {
		logger.Error(fmt.Sprintf("Could not listen on %q: %v", s.instance.Addr, s.port), err)
	}
}

func (s *Server) handleShutdown() context.Context {
	ctx, done := context.WithCancel(context.Background())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer done()

		<-quit
		signal.Stop(quit)
		close(quit)

		logger.Info("Server is shutting down...")

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		s.instance.SetKeepAlivesEnabled(false)
		if err := s.instance.Shutdown(ctx); err != nil {
			logger.Error("Could not gracefully shutdown the server: %v", err)
		}
	}()

	return ctx
}
