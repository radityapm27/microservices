package utils

import (
	"runtime"
	"sync"
	"time"

	"rpm/microservices/core/environ"

	configFile "rpm/microservices/core/config"
	connection "rpm/microservices/core/connection"
	"rpm/microservices/core/model"

	cache "github.com/patrickmn/go-cache"
)

type utils struct {
	env environ.Environ
	cfg configFile.Config
}

var (
	coreUtils    *utils
	oneCoreUtils sync.Once

	caches    *cache.Cache
	onceCache sync.Once
)

type Utils interface {
	GetCache(key string, cache *cache.Cache) interface{}
	GetCacheKey(params ...string) string
	GetCacheHandler() *cache.Cache
	SetCache(key string, value interface{}, cacheParam *cache.Cache) interface{}
	GetClientHandler() connection.Connection
	ConstructApplicationConfig() *model.ApplicationConfig
	GetDatasourceInfo() string
	StringToInt(input string) (int, error)
}

func NewUtils() Utils {

	oneCoreUtils.Do(func() {
		coreUtils = &utils{
			env: environ.NewEnviron(),
			cfg: configFile.NewConfig(),
		}
	})

	return coreUtils
}

// GetCache ...
func (u *utils) GetCache(key string, cache *cache.Cache) interface{} {
	result, found := cache.Get(key)
	if found {
		return result
	}
	return nil
}

// SetCache ...
func (u *utils) SetCache(key string, value interface{}, cacheParam *cache.Cache) interface{} {
	cacheParam.Set(key, value, cache.DefaultExpiration)
	return value
}

// GetCacheKey ...
func (u *utils) GetCacheKey(params ...string) string {
	var fnName string
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		fnName = "?"
	} else {
		fn := runtime.FuncForPC(pc)
		fnName = fn.Name()
	}
	var param string
	for _, prm := range params {
		param += prm
	}

	return fnName + "_" + param
}

var conn connection.Connection
var oneCon sync.Once

// GetClientHandler ...
func (u *utils) GetClientHandler() connection.Connection {
	oneCon.Do(func() {
		conn = connection.NewConnection(u.env, u.ConstructApplicationConfig())
	})
	return conn
}

func (u *utils) ConstructApplicationConfig() *model.ApplicationConfig {
	config := u.cfg.GetApplicationConfig()

	cfg := &model.ApplicationConfig{
		MineDBHost: config.MineDBHost,
		MineDBPort: config.MineDBPort,
		MineDBUser: config.MineDBUser,
		MineDBName: config.MineDBName,
	}
	return cfg
}

// GetCacheHandler ...
func (u *utils) GetCacheHandler() *cache.Cache {
	config := u.cfg.GetApplicationConfig()
	onceCache.Do(func() {
		caches = cache.New(config.CacheExpiry*time.Minute, config.CacheCleanup*time.Minute)
	})
	return caches
}

// GetDatasourceInfo ...
func (u *utils) GetDatasourceInfo() string {
	if typ := u.env.GetRepoType(); typ != "" {
		return typ
	}

	return Postgres
}
