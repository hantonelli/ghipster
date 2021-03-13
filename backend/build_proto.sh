#!/bin/bash

# go get github.com/gogo/protobuf/protobuf
# go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

protoc \
    -I=. \
    -I=$GOPATH/src \
    -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
    -I=$GOPATH/src/github.com/gogo/protobuf/gogoproto \
    -I $GOPATH/src/github.com/gogo/googleapis/ \
    --gogoslick_out=. \
	--swagger_out=. \
    ./proto/*.proto