module rpm/microservices/api

go 1.13

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/gorilla/websocket v1.4.2

	github.com/rs/cors v1.7.0
	github.com/vektah/gqlparser v1.1.2 // indirect
)

require (
	github.com/go-chi/chi v3.3.2+incompatible
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/joho/godotenv v1.4.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/mitchellh/mapstructure v1.3.2 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/vektah/gqlparser/v2 v2.0.1
	rpm/microservices/core v0.0.0
)

replace rpm/microservices/core => ../core

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
