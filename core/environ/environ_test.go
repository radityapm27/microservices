package environ

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ALLOWED_ORIGINS = "*"
	MICRO_REGISTRY  = "kubernetes"
	MINE_DB_HOST    = "192.168.0.1"
	MINE_DB_PORT    = "5432"
	MINE_DB_USER    = "postgres"
	MINE_DB_NAME    = "databasename"
	MINE_DB_PASS    = "postgres"
	LOG_LEVEL       = "info"
	REPO_TYPE       = "postgres"
)

func TestNewEnviron(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
	)

	t.Run("success", func(t *testing.T) {
		env := NewEnviron()
		test.NotNil(env)
	})
}

func TestEnviron_GetAllowedOrigins(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		env  = new(environ)
	)

	t.Run("success", func(*testing.T) {
		res := env.GetAllowedOrigins()
		test.NotNil(res)
		test.Equal(res, "")

	})

	t.Run("not blank", func(t *testing.T) {
		os.Setenv("ALLOWED_ORIGINS", ALLOWED_ORIGINS)
		res := env.GetAllowedOrigins()
		test.NotNil(res)
		test.Equal(res, ALLOWED_ORIGINS)
	})
}

func TestEnviron_IsUseKubernetes(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		env  = new(environ)
	)

	t.Run("success", func(*testing.T) {
		res := env.IsUseKubernetes()
		test.NotNil(res)
		test.Equal(res, false)
	})

	t.Run("true - kubernetes", func(t *testing.T) {
		os.Setenv("MICRO_REGISTRY", MICRO_REGISTRY)
		res := env.IsUseKubernetes()
		test.NotNil(res)
		test.Equal(res, true)
	})
}

func TestEnviron_GetMicroRegistry(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		env  = new(environ)
	)

	t.Run("success", func(*testing.T) {
		res := env.GetMicroRegistry()
		test.NotNil(res)
		test.Equal(res, "")
	})

	t.Run("true - kubernetes", func(t *testing.T) {
		os.Setenv("MICRO_REGISTRY", MICRO_REGISTRY)
		res := env.GetMicroRegistry()
		test.NotNil(res)
		test.Equal(res, MICRO_REGISTRY)
	})
}

func TestEnviron_GetDBHost(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		env  = new(environ)
	)

	t.Run("success", func(*testing.T) {
		res := env.GetDBHost()
		test.NotNil(res)
		test.Equal(res, "")
	})

	t.Run("not blank", func(t *testing.T) {
		os.Setenv("MINE_DB_HOST", MINE_DB_HOST)
		res := env.GetDBHost()
		test.NotNil(res)
		test.Equal(res, MINE_DB_HOST)
	})
}

func TestEnviron_GetDBPort(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		env  = new(environ)
	)

	t.Run("success", func(*testing.T) {
		res := env.GetDBPort()
		test.NotNil(res)
		test.Equal(res, "")
	})

	t.Run("not blank", func(t *testing.T) {
		os.Setenv("MINE_DB_PORT", MINE_DB_PORT)
		res := env.GetDBPort()
		test.NotNil(res)
		test.Equal(res, MINE_DB_PORT)
	})
}

func TestEnviron_GetDBUser(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		env  = new(environ)
	)

	t.Run("success", func(*testing.T) {
		res := env.GetDBUser()
		test.NotNil(res)
		test.Equal(res, "")
	})

	t.Run("not blank", func(t *testing.T) {
		os.Setenv("MINE_DB_USER", MINE_DB_USER)
		res := env.GetDBUser()
		test.NotNil(res)
		test.Equal(res, MINE_DB_USER)
	})
}

func TestEnviron_GetDBName(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		env  = new(environ)
	)

	t.Run("success", func(*testing.T) {
		res := env.GetDBName()
		test.NotNil(res)
		test.Equal(res, "")
	})

	t.Run("not blank", func(t *testing.T) {
		os.Setenv("MINE_DB_NAME", MINE_DB_NAME)
		res := env.GetDBName()
		test.NotNil(res)
		test.Equal(res, MINE_DB_NAME)
	})
}

func TestEnviron_GetDBPass(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		env  = new(environ)
	)

	t.Run("success", func(*testing.T) {
		res := env.GetDBPass()
		test.NotNil(res)
		test.Equal(res, "")
	})

	t.Run("not blank", func(t *testing.T) {
		os.Setenv("MINE_DB_PASS", MINE_DB_PASS)
		res := env.GetDBPass()
		test.NotNil(res)
		test.Equal(res, MINE_DB_PASS)
	})
}

func TestEnviron_GetLogLevel(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		env  = new(environ)
	)

	t.Run("success", func(*testing.T) {
		res := env.GetLogLevel()
		test.NotNil(res)
		test.Equal(res, "")
	})

	t.Run("not blank", func(t *testing.T) {
		os.Setenv("LOG_LEVEL", LOG_LEVEL)
		res := env.GetLogLevel()
		test.NotNil(res)
		test.Equal(res, LOG_LEVEL)
	})
}

func TestEnviron_GetRepoType(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		env  = new(environ)
	)

	t.Run("success", func(*testing.T) {
		res := env.GetRepoType()
		test.NotNil(res)
		test.Equal(res, "")
	})

	t.Run("not blank", func(t *testing.T) {
		os.Setenv("REPO_TYPE", REPO_TYPE)
		res := env.GetRepoType()
		test.NotNil(res)
		test.Equal(res, REPO_TYPE)
	})
}
