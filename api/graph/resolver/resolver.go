package resolver

import (
	"rpm/microservices/api/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service service.Service
}

func New(apiService service.Service) *Resolver {

	return &Resolver{
		service: apiService,
	}
}
