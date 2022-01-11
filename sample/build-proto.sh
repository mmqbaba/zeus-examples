#!/bin/bash


projectpath=. # 具体的项目路径
service=sample # 服务名
pbout=${service}pb

test -f ../proto/${service}.proto || exit 1
# gen-zeus
gen-zeus --proto ../proto/${service}.proto --errdef ../proto/errdef.proto --dest ../
if [ $? -eq 1 ]; then
    echo "gen-zeus failed"
    exit 1
fi

# 生成handler单元测试模板
gotests -w -all ./handler/

mkdir -p $projectpath/proto/${service}pb
cd $projectpath/proto

# gen-gomicro gen-grpc-gateway gen-validator swagger

protoc -I../../proto \
   -I$GOPATH/src \
   -I$GOPATH/src/github.com/google/protobuf/src \
   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
   -I../../../zeus/proto/third_party \
   --go_out=plugins=grpc:./${service}pb \
   --grpc-gateway_out=logtostderr=true:./${service}pb \
   --micro_out=./${service}pb \
   --govalidators_out=./${service}pb \
   --swagger_out=logtostderr=true:. \
   $service.proto

if [ $? -eq 1 ]; then
    echo "protoc failed"
    exit 1
fi
protoc-go-inject-tag -input=./${service}pb/$service.pb.go # inject tag

if [ "$(uname)" == "Darwin" ]; then
    # Mac OS X 操作系统
    sed -i '' -e 's/RegisterExampleHandler(/RegisterExampleHandlerGW(/g' ./${service}pb/$service.pb.gw.go
    sed -i '' -e 's/ RegisterExampleHandler / RegisterExampleHandlerGW /g' ./${service}pb/$service.pb.gw.go
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    # GNU/Linux操作系统
    sed -i 's/RegisterExampleHandler(/RegisterExampleHandlerGW(/g' ./${service}pb/$service.pb.gw.go
    sed -i 's/ RegisterExampleHandler / RegisterExampleHandlerGW /g' ./${service}pb/$service.pb.gw.go
elif [ "$(expr substr $(uname -s) 1 5)" == "MINGW" ]; then
    # Windows NT操作系统
    sed -i 's/RegisterExampleHandler(/RegisterExampleHandlerGW(/g' ./${service}pb/$service.pb.gw.go
    sed -i 's/ RegisterExampleHandler / RegisterExampleHandlerGW /g' ./${service}pb/$service.pb.gw.go
fi

# gen-gomicro gen-grpc-gateway gen-validator swagger
# protoc -I. \
#    -I$GOPATH/src \
#    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
#    --proto_path=${GOPATH}/src/github.com/google/protobuf/src \
#    -I../../../zeus/proto/third_party \
#    --go_out=plugins=grpc:./$service \
#    --grpc-gateway_out=logtostderr=true:./$service \
#    --micro_out=./$service \
#    --govalidators_out=./$service \
#    --swagger_out=logtostderr=true:. \
#    ./$service.proto
# protoc-go-inject-tag -input=./$service/$service.pb.go # inject tag

cd -
