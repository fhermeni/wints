## simple makefile to log workflow
.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:)

all: install test

build:
	@godep go build
	@go build $(GOFLAGS) ./...

install:
	@go get $(GOFLAGS) ./...

test: install	
	@go test $(GOFLAGS) ./...	

cov: install	
	$(eval COVER := $(shell mktemp))
	@echo $(COVER)
	@gotestcover -coverprofile=$(COVER) ./...
	@go tool cover -html=$(COVER)

bench: install
	@go test -run=NONE -bench=. $(GOFLAGS) ./...

clean:
	@go clean $(GOFLAGS) -i ./...

doc:
	@godoc -http=:6060