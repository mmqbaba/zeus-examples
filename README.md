# zeus开发规范

## 目录

- [**GOLANG**](#golang)
- [**开发工具**](#开发工具)
- [**组件**](#组件)
- [**项目初始化**](#项目初始化)
- [**调试**](#调试)
- [**一些代码规范风格**](#一些代码规范风格)
- [**性能相关**](#性能相关)
- [**git版本分支管理**](#git版本分支管理)
- [**git提交代码**](#git提交代码)
- [**框架相关简介**](#框架相关简介)
- [**环境变量**](#环境变量)
- [**单元测试**](#单元测试)
- [**代码覆盖率**](#代码覆盖率)

## golang
------
- [golang](https://golang.google.cn/)，这里使用go1.13，对gomodule有更完善的支持，详情请参考[安装](https://golang.google.cn/doc/install)
- [淘宝的镜像](https://mirrors.aliyun.com/goproxy/)

```
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
```

## 开发工具
------
goland or vscode

- 安装vscode

- 安装格式化插件 Beautify

- 安装 EditorConfig插件,把[.editorconfig](./.editorconfig)拷贝到项目根目录

- 换行, linux 默认 LF, windows 默认 CRLF, 在windows下也设置为LF, 与linux兼容（可通过EditorConfig插件控制）
```javascript
// 在用户设置默认
// 默认行尾字符。使用 \n 表示 LF，\r\n 表示 CRLF。
{ "files.eol": "\n"}
```

## 组件
------

- etcd
### ETCD安装指南
####下载
- 地址
```
https://github.com/etcd-io/etcd/releases
```
- 根据操作系统下载对应的安装包
- 解压
- 测试
```
./etcd --version
```
```
ETCDCTL_API=3 ./etcdctl version
```
####启动
- 设置环境变量
```
ETCDCTL_API=3
```
- 启动服务
```
./etcd
```
####测试
- 设置和获取值
```
./etcdctl put /mykey "this is awesome"
```
```
./etcdctl get /mykey
```

####Web 管理工具
- 下载
```
https://github.com/evildecay/etcdkeeper/releases
```

- mingw64

windows下需要安装此工具包。安装包 tools/mingw64_x86_64-4.8.2-release-posix-seh-rt_v3-rev2.7z。解压缩后设置目录'mingw64\bin'到path变量

- 日志组件

[logrus](https://github.com/sirupsen/logrus)

- 数据库

- 数据校验

## 项目初始化

([gen-zeus](http://gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/tree/master/tools/gen-zeus))
------

- 安装gen-zeus工具
```bash
git clone ssh://git@gitlab.dg.com:10086/BackEnd/jichuchanpin/tif/zeus.git
cd ./zeus
go build -o tools/bin/ ./tools/gen-zeus
cp tools/bin/gen-zeus $GOPATH/bin/
```

###  安装 protoc-gen-* plugin
```bash
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/protoc-gen-micro
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
go get github.com/favadi/protoc-go-inject-tag
```

- 初始化项目

定义proto `hello.proto`

```proto
syntax="proto3";

package hello;

import "google/api/annotations.proto";
import "github.com/mwitkow/go-proto-validators/validator.proto";

message HelloRequest{
    string name = 1;
    // @inject_tag: validate:"required,gt=0,max=1000"
    int32 age = 2 [(validator.field) = {int_gt: 20, int_lt: 27}];
}

message HelloReply{
    string message=1;
}

message PingRequest{
    string ping = 1;
}

message PongReply{
    string pong = 1;
}

service Hello{
    rpc SayHello(HelloRequest) returns (HelloReply){
        option (google.api.http) = {
            post: "/v1/hello"
            body: "*"
            };
    }

    rpc PingPong(PingRequest) returns (PongReply){
        option (google.api.http) = {
            post: "/v1/pingpong"
            body: "*"
            };
    }
}
```

初始化并启动

设置配置入口文件
```json
// 默认路径linux: /etc/tif/zeus.json windows: C:\\tif\\zeus.json
{
    "engine_type": "etcd",
    "config_path": "/zeus/dzqz", // 服务应用的配置路径
    "config_format": "json",     // 配置格式
	"endpoints": ["127.0.0.1:2379"],
	"username": "root",
	"password": "123456"
}
```
设置应用服务配置
```bash
# 设值
ETCDCTL_API=3 etcdctl put /zeus/dzqz {\"go_micro\":{\"server_name\":\"zeus\",\"registry_plugin_type\":\"etcd\",\"registry_addrs\":[\"127.0.0.1:2379\"],\"registry_authuser\":\"root\",\"registry_authpwd\":\"123456\"},\"mongodb\":{\"host\":\"127.0.0.1:27017\",\"user\":\"root\",\"pwd\":\"123456\"},\"mysql_source\":{\"e_seal\":{\"datasourcename\":\"root:123456@tcp(localhost:3306)/e_seal\",\"maxidleconns\":30,\"maxopenconns\":1024}},\"redis\":{\"host\":\"127.0.0.1:6379\",\"pwd\":\"\",\"enable\":false},\"log_conf\":{\"log\":\"console\",\"level\":\"debug\",\"format\":\"text\",\"rotation_time\":\"hour\",\"log_dir\":\"./\"},\"ext\":{\"httphandler_pathprefix\":\"\",\"grpcgateway_pathprefix\":\"\"}}
```

初始化并启动
```bash
gen-zeus --proto ./hello.proto
cd hello
./build-proto.sh
go run ./cmd/app
```

浏览器打开链接[swagger-ui](http://localhost:8081/swagger-ui/)

编译项目
```bash
make build
```

项目目录结构
```txt
分层请求的流向: {http,rpc} => handler => logic => resource
./hello
|-- global        // 一些全局变量，方法函数，初始化
|   -- global.go
|-- cmd
|   -- app        // 应用启动入口
|      -- main.go
|      -- init.go
|-- rpc          // rpc服务handler注册
|   -- rpc.go
|   -- init.go
|-- http         // http服务handler注册，路由定义
|   -- http.go
|   -- init.go
|-- handler  // rpc/http请求的处理逻辑
|-- logic    // 业务逻辑层，根据业务复杂度决定是否需要分层
|-- proto    // proto
|-- resource // 远程资源访问
|   -- dao
|   -- cache
|   -- http
|   -- rpc
|   -- broker
|-- Dockerfile
|-- Makefile
|-- README.md
```

## 调试
------
vscode开发环境下, f5

## 一些代码规范风格
------

- 

## 性能相关
------

- 

## git版本分支管理
------

- 

## git提交代码
------

- 添加必要的提交描述(功能, bug相关)

- 先拉取, 检查语法, 再推送


## 框架相关简介
------

- 基于[gomicro](https://github.com/micro/go-micro)

## 环境变量 
------
- GOMAXPROCS cpu核数


## 单元测试
------
### 对于一些关键代码最好附加上单元测试
------

- 测试

## 代码覆盖率
------
[recovery]
------
- recovery
