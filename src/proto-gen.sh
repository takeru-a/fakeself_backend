#!/bin/bash
rm -rf grpc/*.go
protoc --proto_path=api --go_out=grpc --go_opt=paths=source_relative \
  --go-grpc_out=grpc --go-grpc_opt=paths=source_relative \
  api/*.proto
