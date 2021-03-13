#!/bin/bash

protoc \
		-I proto \
		-I $GOPATH/src/github.com/gogo/googleapis/ \
        -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
        -I=$GOPATH/src/github.com/gogo/protobuf/gogoproto \
		-I $GOPATH/src/ \
		--gogo_out=plugins=grpc,paths=source_relative,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
./proto/ \
		--grpc-gateway_out=allow_patch_feature=false,paths=source_relative,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
./proto/ \
		--swagger_out=third_party/OpenAPI/ \
		--govalidators_out=gogoimport=true,paths=source_relative,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
./proto \
		proto/products.proto