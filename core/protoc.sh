#!/bin/bash

# install protoc-gen-micro v2 first
go get github.com/micro/micro/v2/cmd/protoc-gen-micro@v2.9.1

# execute protoc command
protoc --proto_path=${pwd}:. --micro_out=. --go_out=. proto/*.proto