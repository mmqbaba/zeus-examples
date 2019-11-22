#!/bin/bash

projectpath=. # 具体的项目路径
service=hello # 服务名

# gen-zeus --proto ./proto/$service.proto # 生成或更新项目目录结构

cd $projectpath/proto

# gen-gomicro gen-grpc-gateway gen-validator swagger
mkdir $service
protoc -I. \
   -I../../../zeus/proto/third_party \
   -I$GOPATH/src \
   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
   --go_out=plugins=grpc:./$service \
   --grpc-gateway_out=logtostderr=true:./$service \
   --micro_out=./$service \
   --govalidators_out=./$service \
   --swagger_out=logtostderr=true:. \
   $service.proto
protoc-go-inject-tag -input=./$service/$service.pb.go # inject tag

sed -i 's/RegisterHelloHandler(/RegisterHelloHandlerGW(/g' ./$service/$service.pb.gw.go
sed -i 's/ RegisterHelloHandler / RegisterHelloHandlerGW /g' ./$service/$service.pb.gw.go

# gen-gomicro gen-grpc-gateway gen-validator swagger
# protoc -I. \
#    -I$GOPATH/src \
#    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
#    --go_out=plugins=grpc:./$service \
#    --grpc-gateway_out=logtostderr=true:./$service \
#    --micro_out=./$service \
#    --govalidators_out=./$service \
#    --swagger_out=logtostderr=true:. \
#    ./$service.proto
# protoc-go-inject-tag -input=./$service/$service.pb.go # inject tag