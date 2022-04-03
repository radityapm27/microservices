package elastic

import (
	"rpm/microservices/core/connection"
	"rpm/microservices/core/utils"
)

// Repository ...
type Repository struct {
	Connection connection.Connection
	CoreUtils  utils.Utils
}
