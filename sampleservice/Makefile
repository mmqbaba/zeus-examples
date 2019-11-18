

.PHONY:  ALL build init autoinit

service=hello

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
	[ -f proto/${service}/${service}.micro.go ] || sh build-proto.sh


start:
	./$(service)_server

stop:
	killall $(service)_server

restart: stop start

