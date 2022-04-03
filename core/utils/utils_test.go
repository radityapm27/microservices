package utils

import (
	"testing"

	mocksConfig "rpm/microservices/core/mocks/config"
	mocksEnv "rpm/microservices/core/mocks/environ"
	"rpm/microservices/core/model"

	"github.com/stretchr/testify/assert"
)

func resetConfig(u *utils) *mocksConfig.Config {
	conf := new(mocksConfig.Config)
	u.cfg = conf
	return conf
}

func resetEnviron(u *utils) *mocksEnv.Environ {
	env := new(mocksEnv.Environ)
	u.env = env
	return env
}

func TestUtils_NewUtils(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
	)

	t.Run("success", func(t *testing.T) {
		u := NewUtils()
		test.NotNil(u)
	})
}

func TestUtils_GetCacheKey(t *testing.T) {
	t.Parallel()

	var utils = new(utils)

	t.Run("get key", func(t *testing.T) {
		expected := "rpm/microservices/core/utils.TestUtils_GetCacheKey.func1_ab"
		result := utils.GetCacheKey("a", "b")
		assert.Equal(t, expected, result)
	})
}

func TestUtils_ConstructApplicationConfig(t *testing.T) {
	t.Parallel()

	var (
		test    = assert.New(t)
		utils   = new(utils)
		mockCfg = new(model.ApplicationConfig)
	)

	cfg := resetConfig(utils)
	cfg.On("GetApplicationConfig").Return(mockCfg)

	t.Run("success", func(t *testing.T) {
		conn := utils.ConstructApplicationConfig()
		test.NotNil(conn)
	})
}

func TestUtils_GetCacheHandler(t *testing.T) {
	t.Parallel()

	var (
		utils   = new(utils)
		test    = assert.New(t)
		key     = "key"
		value   = "value"
		mockCfg = new(model.ApplicationConfig)
	)

	cfg := resetConfig(utils)
	mockCfg.CacheExpiry = 5
	mockCfg.CacheCleanup = 5
	cfg.On("GetApplicationConfig").Return(mockCfg)
	cacheParam := utils.GetCacheHandler()

	t.Run("set cache", func(t *testing.T) {
		result := utils.SetCache(key, value, cacheParam)
		test.Equal(value, result)
	})

	t.Run("get cache exist", func(t *testing.T) {
		result := utils.GetCache(key, cacheParam)
		test.Equal(value, result)
	})

	t.Run("get cache not exist", func(t *testing.T) {
		result := utils.GetCache("otherkey", cacheParam)
		test.Nil(result)
	})
}

func TestUtils_GetDatasourceInfo(t *testing.T) {
	t.Parallel()

	var (
		test  = assert.New(t)
		utils = new(utils)
	)

	t.Run("default data source", func(t *testing.T) {
		env := resetEnviron(utils)
		env.On("GetRepoType").Return("")
		result := utils.GetDatasourceInfo()
		test.Equal(Postgres, result)
	})

	t.Run("data source from env", func(t *testing.T) {
		repoType := "mysql"
		env := resetEnviron(utils)
		env.On("GetRepoType").Return(repoType)
		result := utils.GetDatasourceInfo()
		test.Equal(repoType, result)
	})
}
