package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
	)

	t.Run("test success", func(t *testing.T) {
		cfg := NewConfig()
		test.NotNil(cfg)
	})
}

func TestConfig_GetApplicationConfig(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		c    = new(config)
	)

	t.Run("success", func(t *testing.T) {

		cc := c.GetApplicationConfig()
		test.NotNil(cc)
	})
}
