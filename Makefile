#!/usr/bin/make -f

VERSION := $(shell git describe)

test:
	go test --race ./...

install: test
	go install -ldflags="-X 'main.Version=$(VERSION)'" github.com/mdw-tools/md2html/cmd/md2html
