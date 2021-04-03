#!/bin/bash

# https://stackoverflow.com/questions/592620/how-can-i-check-if-a-program-exists-from-a-bash-script
if ! command -v protoc &> /dev/null
then
    wget -O https://github.com/protocolbuffers/protobuf/releases/download/v3.15.6/protoc-3.15.6-linux-x86_64.zip
    unzip protoc-3.15.6-linux-x86_64.zip -d $HOME/protobuf
    cp $HOME/protobuf/bin/protoc $HOME/go/bin
fi

# export GO111MODULE=on  # Enable module mode
# go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
# go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
# go get github.com/gogo/protobuf/protobuf
# go get github.com/gogo/googleapis

# add to ~/bachrc -> export GOPATH=$HOME/gopath
# export GO111MODULE=off
# go get github.com/gogo/googleapis github.com/gogo/protobuf/protobuf github.com/gogo/protobuf/gogoproto github.com/gogo/protobuf/proto github.com/gogo/protobuf/jsonpb github.com/gogo/protobuf/protoc-gen-gogo github.com/gogo/protobuf/gogoproto github.com/mwitkow/go-proto-validators/protoc-gen-govalidators


protoc \
    -I=. \
    -I=$GOPATH/src \
    -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
    -I=$GOPATH/src/github.com/gogo/protobuf/gogoproto \
    -I $GOPATH/src/github.com/gogo/googleapis/ \
    --gogoslick_out=. \
	--swagger_out=. \
    ./proto/*.proto