#!/bin/sh
protoc -I=../proto/v1 --go_out=plugins=grpc,import_path=api_v1:./service/v1 ../proto/v1/*.proto
