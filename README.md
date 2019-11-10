# zeus-examples

zeus-examples

- install protoc-gen-* plugin
```bash
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/protoc-gen-micro
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
go get github.com/favadi/protoc-go-inject-tag
```

- gen-grpc
```bash
cd /yourpath/to/yourproject/proto
mkdir gw
protoc -I. \
   -I$GOPATH/src \
   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
   --go_out=plugins=grpc:./gw \
   ./hello.proto
```

- gen-grpc-gateway
```bash
cd /yourpath/to/yourproject/proto
mkdir gw
protoc -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:./gw \
   ./hello.proto
```

- gen-gomicro
```bash
cd /yourpath/to/yourproject/proto
mkdir gomicro
protoc -I. \
   -I$GOPATH/src \
   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
   --micro_out=./gomicro/ \
   --go_out=./gomicro/ \
   ./hello.proto
```

- gen-swagger
```bash
cd /yourpath/to/yourproject/proto
protoc -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:. \
   ./hello.proto
```

- gen-validator
```bash
cd /yourpath/to/yourproject/proto

mkdir gomicro
protoc  \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --proto_path=${GOPATH}/src \
  --proto_path=${GOPATH}/src/github.com/google/protobuf/src \
  --proto_path=. \
  --govalidators_out=./gomicro \
  ./hello.proto

mkdir gw
protoc  \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --proto_path=${GOPATH}/src \
  --proto_path=${GOPATH}/src/github.com/google/protobuf/src \
  --proto_path=. \
  --govalidators_out=./gw \
  ./hello.proto
```