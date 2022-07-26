#!/bin/bash

# https://stackoverflow.com/questions/592620/how-can-i-check-if-a-program-exists-from-a-bash-script
# ALTERNATIVE go install google.golang.org/protobuf/cmd/protoc-gen-go
if ! command -v protoc &> /dev/null
then
    wget https://github.com/protocolbuffers/protobuf/releases/download/v21.4/protoc-21.4-linux-x86_64.zip
    unzip protoc-21.4-linux-x86_64.zip -d $HOME/protobuf
    cp $HOME/protobuf/bin/protoc $HOME/go/bin
fi

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go get github.com/envoyproxy/protoc-gen-validate
go get google.golang.org/protobuf


cd ~/code/projects
git clone https://github.com/protocolbuffers/protobuf.git

go get -d github.com/envoyproxy/protoc-gen-validate
cd $GOPATH/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.7
ln -s /home/hantone/code/projects/protobuf/src/google google
chmod -R u=rwx .
protoc-gen-validate@v0.6.7: make build

	# github.com/go-chi/chi v3.3.2+incompatible
	# github.com/gogo/googleapis v1.4.1
	# github.com/gogo/protobuf v1.3.2
	# github.com/grpc-ecosystem/grpc-gateway/v2 v2.3.0
	# github.com/mattn/go-sqlite3 v1.14.6
	# github.com/mwitkow/go-proto-validators v0.3.2
	# github.com/vektah/gqlparser/v2 v2.1.0

# export GO111MODULE=on  # Enable module mode
# go get google.golang.org/protobuf/cmd/protoc-gen-go \
#     google.golang.org/grpc/cmd/protoc-gen-go-grpc \
#     github.com/gogo/protobuf/protoc-gen-gogoslick

# go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
# go get github.com/gogo/protobuf/protobuf
# go get github.com/gogo/googleapis

# add to ~/bachrc -> export GOPATH=$HOME/gopath
# export GO111MODULE=off
# go get github.com/gogo/googleapis github.com/gogo/protobuf/protobuf github.com/gogo/protobuf/gogoproto github.com/gogo/protobuf/proto github.com/gogo/protobuf/jsonpb github.com/gogo/protobuf/protoc-gen-gogo github.com/gogo/protobuf/gogoproto github.com/mwitkow/go-proto-validators/protoc-gen-govalidators


# OTHER WAY TO IMPORT - https://stepan.wtf/importing-protobuf-with-go-modules/
# go mod vendor

# mkdir -p ~/gopath/src/github.com/gogo/googleapis/
# mkdir -p ~/gopath/src/github.com/gogo/protobuf/
# cp -r ~/go/pkg/mod/github.com/gogo/googleapis@v1.4.1 ~/gopath/src/github.com/gogo/googleapis
# cp -r ~/go/pkg/mod/github.com/gogo/protobuf@v1.3.2 ~/gopath/src/github.com/gogo/protobuf

# protoc \
# 		-I proto \
# 		# -I $GOPATH/src/github.com/gogo/googleapis/ \
#         # -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
#         # -I=$GOPATH/src/github.com/gogo/protobuf/gogoproto \
# 		# -I $GOPATH/src/ \
# 		--proto_path=. \
# 		--go_out=:\
# . \
# 		proto/product/products.proto

protoc  \
  -I . \
  -I ${GOPATH}/src \
  -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.7 \
    --go_out=.  \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
    --validate_out="lang=go:." \
    --validate_opt=paths=source_relative \
    proto/product/products.proto