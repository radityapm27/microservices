package middleware

import (
	"context"
	"net/http"
	"sync"

	coreClient "rpm/microservices/core/client"
	configFile "rpm/microservices/core/config"
	"rpm/microservices/core/environ"
	"rpm/microservices/core/utils"
)

var (
	once sync.Once
	mCl  *middleware
)

type middleware struct {
	mClient   coreClient.MicroClient
	cfg       configFile.Config
	coreUtils utils.Utils
	env       environ.Environ
}
type Middleware interface {
	GetMiddleware(next http.Handler) http.Handler
}

func NewMiddleware() Middleware {
	once.Do(func() {
		mCl = &middleware{
			mClient:   coreClient.NewMicroClient(),
			cfg:       configFile.NewConfig(),
			coreUtils: utils.NewUtils(),
			env:       environ.NewEnviron(),
		}
	})
	return mCl
}

// Middleware ...
func (m *middleware) GetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), utils.ContextHostKey, r.Host)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
