

.PHONY:  ALL build init autoinit

service=hellodemo

build_date = $(shell date '+%Y-%m-%d %H:%M:%S')
version = $(shell git describe --tags --always | sed 's/-/+/' | sed 's/^v//')
goversion = $(shell go version)
ldflags = -ldflags "-X 'main.Version=$(version)' -X 'main.BuildDate=$(build_date)' -X 'main.GoVersion=$(goversion)'"

ALL: build



build: autoinit
	go build $(ldflags) -o $(service)_server ./cmd/app


init:
	sh build-proto.sh

autoinit:
	[ -f proto/hellodemopb/hellodemo.micro.go ] || sh build-proto.sh


start:
	./$(service)_server --port 9090 --apiPort 8081 --apiInterface 127.0.0.1 --confEntryPath /zeus/$(service)

stop:
	killall $(service)_server

restart: stop start

docker: build docker-build


docker-build:
	docker build .

