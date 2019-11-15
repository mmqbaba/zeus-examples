# zeus-examples

zeus-examples

###  install protoc-gen-* plugin
```bash
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/protoc-gen-micro
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
go get github.com/favadi/protoc-go-inject-tag
```

### 实例
* 生成项目
```bash
gen-zeus -proto proto/hello.proto
```
* 编译项目
```bash
cd hello && make build
```
### install [gen-zeus](http://gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/blob/master/README.md)
- [项目生成器](http://gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/blob/master/README.md)
- [使用说明](http://gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/blob/master/tools/gen-zeus/README.md)