SHELL := /usr/bin/env bash

build:
	go build -o ./bin/cipher main.go

install:
	cp -f ./bin/cipher $(HOME)/go/bin/

fmt:
	go mod tidy -compat=1.17
	gofmt -l -s -w .

test:
	go test ./...
