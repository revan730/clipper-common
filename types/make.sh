#/bin/sh

protoc --gofast_out=plugins=grpc:. CI.proto