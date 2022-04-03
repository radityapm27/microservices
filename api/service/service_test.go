package service

import (
	"testing"

	mocksProto "rpm/microservices/core/mocks/proto"
	mocksUtils "rpm/microservices/core/mocks/utils"

	"github.com/stretchr/testify/assert"
)

func TestService_New(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
	)

	t.Run("success", func(t *testing.T) {
		svc := New(new(mocksProto.ShoppingCartService), new(mocksProto.ProductCatalogService), new(mocksProto.UserService))
		objSvc := svc.(*service)
		test.NotNil(objSvc)
		test.NotNil(objSvc.shoppingCart)
		test.NotNil(objSvc.catalog)
		test.NotNil(objSvc.user)
		test.NotNil(objSvc.utils)
	})
}

func resetShoppingCart(s *service) *mocksProto.ShoppingCartService {
	svc := new(mocksProto.ShoppingCartService)
	s.shoppingCart = svc
	return svc
}

func resetCatalog(s *service) *mocksProto.ProductCatalogService {
	svc := new(mocksProto.ProductCatalogService)
	s.catalog = svc
	return svc
}

func resetUser(s *service) *mocksProto.UserService {
	svc := new(mocksProto.UserService)
	s.user = svc
	return svc
}

func resetUtils(s *service) *mocksUtils.Utils {
	utils := new(mocksUtils.Utils)
	s.utils = utils
	return utils
}
