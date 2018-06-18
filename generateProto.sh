#!/bin/sh
protoc -I=/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I=../proto/v1 \
    --go_out=plugins=grpc,import_path=api_v1:./service/v1 \
    ../proto/v1/*.proto

protoc -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I=../proto/v1 \
    --grpc-gateway_out=logtostderr=true,import_path=api_v1:./service/v1 \
    ../proto/v1/*.proto

protoc-go-inject-tag -input=./service/v1/address.pb.go
protoc-go-inject-tag -input=./service/v1/legacy.pb.go

#protoc -I/usr/local/include \
#    -I. \
#    -I$GOPATH/src \
#    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#    -I=../proto/v1 \
#    --swagger_out=logtostderr=true:./service/v1/swagger \
#    ../proto/v1/*.proto
