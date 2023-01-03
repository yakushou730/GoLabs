SHELL := /bin/bash

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

brew_install:
	brew install bufbuild/buf/buf

buf:
	buf lint
	buf generate
