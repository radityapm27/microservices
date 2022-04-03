module rpm/microservices/core

go 1.13

require (
	github.com/denisenkom/go-mssqldb v0.0.0-20191128021309-1d7a30a10f73
	github.com/elastic/go-elasticsearch/v7 v7.4.1
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.3.0

	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/rs/zerolog v1.21.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
