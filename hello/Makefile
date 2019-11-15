

.PHONY:  ALL build init autoinit

service=hello

ALL: build



build: autoinit
	go build -o $(service)_server ./cmd/app


init:
	sh build-proto.sh

autoinit:
	[ -d proto/gw ] || sh build-proto.sh

