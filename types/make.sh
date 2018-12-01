#/bin/sh

protoc --go_out=plugins=grpc:. CI.proto