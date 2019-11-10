#!/bin/bash

projectpath=/e/works/projects/zeus-examples/sampleservice
service=hello
cd $projectpath/proto

# gen-gomicro gen-validator
mkdir gomicro
protoc -I. \
   -I$GOPATH/src \
   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
   --proto_path=${GOPATH}/src/github.com/google/protobuf/src \
   --go_out=./gomicro/ \
   --micro_out=./gomicro/ \
   --govalidators_out=./gomicro \
   ./$service.proto
protoc-go-inject-tag -input=./gomicro/$service.pb.go # inject tag

# gen-grpc gen-grpc-gateway gen-validator
mkdir gw
protoc -I. \
   -I$GOPATH/src \
   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
   --proto_path=${GOPATH}/src/github.com/google/protobuf/src \
   --go_out=plugins=grpc:./gw \
   --grpc-gateway_out=logtostderr=true:./gw \
   --govalidators_out=./gw \
   ./$service.proto
protoc-go-inject-tag -input=./gw/$service.pb.go # inject tag

# gen-swagger
protoc -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:. \
   ./$service.proto