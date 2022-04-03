## Prerequisites

* [Go](https://golang.org/doc/install) v1.13
* [Protobuf](https://developers.google.com/protocol-buffers/docs/downloads) v3.x.x (or via `brew install protobuf`)
* [protoc-gen-go](https://pkg.go.dev/mod/github.com/golang/protobuf@v1.3.5) v1.3.5
```
  GIT_TAG="v1.3.5"
  go get -d -u github.com/golang/protobuf/protoc-gen-go
  git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
  go install github.com/golang/protobuf/protoc-gen-go
```

* [protoc-gen-micro](https://pkg.go.dev/github.com/micro/protoc-gen-micro@v1.0.0) v1.0.0
```
  GIT_TAG="v1.0.0"
  go get -d -u github.com/micro/protoc-gen-micro
  git -C "$(go env GOPATH)"/src/github.com/micro/protoc-gen-micro checkout $GIT_TAG
  go install github.com/micro/protoc-gen-micro
```


## Quick Config

* Enter directory core and generate proto file
```
  cd core && ./protoc.sh
```

* Generate graphql schema
```
  cd api \
  && go run github.com/99designs/gqlgen generate
```

## Run server graphql services
```
  cd api \
  && go run server/server.go
```

## Run example services
```
  cd services/<ServiceName> \
  && go run server/server.go 
```

## Run GraphQL on Playground
- Running Script on your localhost Postgresql with DB Name = Microservices based on .env for each services
- Open [localhost](http://localhost:8080/)


## Cart

**Query (read) - Get Shopping Cart by User ID**

```
query Cart($userID: String!) {
  shoppingCart {
    getShoppingCartByUserId(userID: $userID) {
      cart {
        product {
          productName
          productDescription
          price
          stock
        }
        quantity
      }
      userInfo {
        userID
        userName
        userLocation
      }
    }
  }
}
```
**Query Variable Example**

```
{
  "userID": "test1@gmail.com"
}
```

## Catalog

**Query (read) - List of Product**


```
query Catalog {
  productCatalog {
    getProductCatalog {
      catalogs {
        productName
        productDescription
        price
        stock
      }
    }
  }
}

```

## User

**Query (read) - List of User**


```
 query User {
  user {
    listOfUser {
      users {
        userID
        userName
        userLocation
      }
    }
  }
}

```

**Query (read) - Get User Info By User ID**


```
query User($userID: String!) {
   user{
    getUserInfoById(userID: $userID) {
       userID
        userName
        userLocation
    }
  }
}
```
**Query Variable Example**

```
{
  "userID": "test1@gmail.com"
}
```

****


**Unit-Test**

- 
- Running ./mockery.sh to generate mocks

- To run test and Make file coverage.out
```
go test ./... -coverprofile=coverage.out
```