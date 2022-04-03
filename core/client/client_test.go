package client

import (
	"testing"

	mocksConfig "rpm/microservices/core/mocks/config"
	mocksEnv "rpm/microservices/core/mocks/environ"
	mocksUtils "rpm/microservices/core/mocks/utils"
	"rpm/microservices/core/model"

	"github.com/stretchr/testify/assert"
)

func resetUtils(c *microClient) *mocksUtils.Utils {
	utils := new(mocksUtils.Utils)
	c.coreUtils = utils
	return utils
}

func resetEnviron(c *microClient) *mocksEnv.Environ {
	env := new(mocksEnv.Environ)
	c.env = env
	return env
}

func resetConfig(c *microClient) *mocksConfig.Config {
	conf := new(mocksConfig.Config)
	c.cfg = conf
	return conf
}

func TestNewMicroClient(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
	)

	t.Run("success", func(t *testing.T) {
		cl := NewMicroClient()
		test.NotNil(cl)
	})
}

func TestMicroClient_GetMicroServiceClientCart(t *testing.T) {
	t.Parallel()

	var (
		test    = assert.New(t)
		cl      = new(microClient)
		mockCfg = new(model.ApplicationConfig)
	)

	t.Run("success", func(t *testing.T) {
		cfg := resetConfig(cl)
		cfg.On("GetApplicationConfig").Return(mockCfg)

		env := resetEnviron(cl)
		env.On("GetMicroRegistry").Return("")

		auth := cl.GetMicroServiceClientCart()

		cfg.AssertExpectations(t)

		test.NotNil(auth)
	})
}

func TestMicroClient_GetMicroServiceClientCatalog(t *testing.T) {
	t.Parallel()

	var (
		test    = assert.New(t)
		cl      = new(microClient)
		mockCfg = new(model.ApplicationConfig)
	)

	t.Run("success", func(t *testing.T) {
		cfg := resetConfig(cl)
		cfg.On("GetApplicationConfig").Return(mockCfg)

		env := resetEnviron(cl)
		env.On("GetMicroRegistry").Return("")

		auth := cl.GetMicroServiceClientCatalog()

		cfg.AssertExpectations(t)

		test.NotNil(auth)
	})
}

func TestMicroClient_GetMicroServiceClientUser(t *testing.T) {
	t.Parallel()

	var (
		test    = assert.New(t)
		cl      = new(microClient)
		mockCfg = new(model.ApplicationConfig)
	)

	t.Run("success", func(t *testing.T) {
		cfg := resetConfig(cl)
		cfg.On("GetApplicationConfig").Return(mockCfg)

		env := resetEnviron(cl)
		env.On("GetMicroRegistry").Return("")

		auth := cl.GetMicroServiceClientUser()

		cfg.AssertExpectations(t)

		test.NotNil(auth)
	})
}
