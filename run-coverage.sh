#!/bin/bash

# Put your modules in here
modules=(
  './api'
  './core'
  './services/cart'
  './services/catalog'
  './services/user'
)

cp core/config/app.config.dev.json core/config/.app.config.json
(cd ./api && go run github.com/99designs/gqlgen --verbose)
(cd ./core/ && ./protoc.sh)

# Create coverage.out file
touch coverage.out

for idx in "${!modules[@]}"; do
  module=${modules[$idx]}
  (cd ${module} && go test -coverprofile coverage.out ./... && go tool cover -func=coverage.out)

  if [[ $idx -eq 0 ]]; then
    cat ${module}/coverage.out >coverage.out
  else
    cat ${module}/coverage.out | grep -v mode >>coverage.out
  fi
done
