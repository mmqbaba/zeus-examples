module zeus-examples

go 1.16

replace (
	github.com/golang/lint => golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
	github.com/micro/go-micro => github.com/maidol/go-micro v1.18.1
	github.com/mmqbaba/zeus => github.com/mmqbaba/zeus v0.6.12
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.4
    google.golang.org/grpc => google.golang.org/grpc v1.25.1
)

require (
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.12.0
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/mmqbaba/zeus v0.0.0
	github.com/mwitkow/go-proto-validators v0.2.0
	github.com/sirupsen/logrus v1.8.1
	go.mongodb.org/mongo-driver v1.5.0
	google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a
	google.golang.org/grpc v1.25.1
)
