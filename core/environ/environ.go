package environ

import (
	"os"
	"sync"
)

var (
	envOnce sync.Once
	env     Environ
)

type environ struct {
}

// Environ ..
type Environ interface {
	GetAllowedOrigins() string
	IsUseKubernetes() bool
	GetMicroRegistry() string
	GetDBHost() string
	GetDBPort() string
	GetDBUser() string
	GetDBName() string
	GetDBPass() string
	GetLogLevel() string
	GetRepoType() string
}

func createEnviron() Environ {
	return &environ{}
}

// GetEnviron ...
func NewEnviron() Environ {
	envOnce.Do(func() {
		env = createEnviron()
	})
	return env
}

//GetAuthServiceAddress
func (e *environ) GetAuthServiceAddress() string {
	return os.Getenv("AUTH_SERVICE_ADDRESS")
}

//GetApplicationID
func (e *environ) GetApplicationID() string {
	return os.Getenv("APP_ID")
}

//GetApplicationAccessKey
func (e *environ) GetApplicationAccessKey() string {
	return os.Getenv("APP_ACCESS_KEY")
}

//GetAPIListeningPort
func (e *environ) GetAPIListeningPort() string {
	return os.Getenv("PORT")
}

//GetAllowedOrigins
func (e *environ) GetAllowedOrigins() string {
	return os.Getenv("ALLOWED_ORIGINS")
}

//IsUseKubernetes
func (e *environ) IsUseKubernetes() bool {
	return os.Getenv("MICRO_REGISTRY") == "kubernetes"
}

//GetMicroRegistry
func (e *environ) GetMicroRegistry() string {
	return os.Getenv("MICRO_REGISTRY")
}

//GetDBHost
func (e *environ) GetDBHost() string {
	return os.Getenv("MINE_DB_HOST")
}

//GetDBPort
func (e *environ) GetDBPort() string {
	return os.Getenv("MINE_DB_PORT")
}

//GetDBUser
func (e *environ) GetDBUser() string {
	return os.Getenv("MINE_DB_USER")
}

//GetDBName
func (e *environ) GetDBName() string {
	return os.Getenv("MINE_DB_NAME")
}

//GetDBPass
func (e *environ) GetDBPass() string {
	return os.Getenv("MINE_DB_PASS")
}

//GetLogLevel
func (e *environ) GetLogLevel() string {
	return os.Getenv("LOG_LEVEL")
}

//GetRepoType
func (e *environ) GetRepoType() string {
	return os.Getenv("REPO_TYPE")
}
