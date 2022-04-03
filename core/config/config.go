package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	common "rpm/microservices/core/model"

	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/logger"
)

var (
	appConfig  *config
	once       sync.Once
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type config struct {
	ApplicationConfig *common.ApplicationConfig
	FileModified      time.Time
}

type Config interface {
	GetApplicationConfig() *common.ApplicationConfig
}

func NewConfig() Config {
	once.Do(func() {
		appConfig = new(config)
	})
	return appConfig
}

// GetApplicationConfig ..
func (cfg *config) GetApplicationConfig() *common.ApplicationConfig {
	fileName := basepath + "/.app.config.json"

	file, err := os.Stat(fileName)
	if err != nil {
		logger.Error(errors.BadRequest("common.config.GetApplicationConfig", err.Error()), "error on GetApplicationConfig")
		return nil
	}

	if file.ModTime().After(cfg.FileModified) {
		logger.Info("Load application config from files ..")

		jsonFile, err := os.Open(fileName)
		// if we os.Open returns an error then handle it
		if err != nil {
			logger.Error(errors.BadRequest("common.config.GetApplicationConfig", err.Error()), "error on GetApplicationConfig")
			return nil
		}

		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var appConfig *common.ApplicationConfig

		json.Unmarshal([]byte(byteValue), &appConfig)

		cfg.ApplicationConfig = appConfig
		cfg.FileModified = file.ModTime()
	}
	return cfg.ApplicationConfig
}
