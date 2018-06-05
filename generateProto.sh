#!/bin/sh
protoc -I=../proto/v1 -I=/usr/local/include -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,import_path=api_v1:./service/v1 ../proto/v1/*.proto
