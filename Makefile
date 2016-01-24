## simple makefile to log workflow
.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:)

all: install test

build:
	@echo "==== go build ===="		
	@godep go build $(GOFLAGS) ./...

install:
	@go get $(GOFLAGS) ./...

vet:
	@echo "==== go vet ===="
	@go vet ./...

lint:
	@echo "==== go lint ===="
	@golint ./**/*.go

test: build vet lint	
	@echo "==== go test ===="
	@godep go test $(GOFLAGS) ./...	

test-integration: install
	@echo "==== integration test ===="
	@godep go test -v -tags "integration" ./...

errcheck:	
	@echo "==== errcheck ===="
	@errcheck ./...

cov: install	
	@echo "==== coverage ===="
	$(eval COVER := $(shell mktemp))
	@echo $(COVER)
	@gotestcover -coverprofile=$(COVER) ./...
	@go tool cover -html=$(COVER)

bench: install
	@echo "==== bench ===="
	@godep go test -tags "integration" -run=NONE -bench=. $(GOFLAGS) ./...

clean:
	@godep go clean $(GOFLAGS) -i ./...

doc:
	@echo "==== go doc is running. Can be moved to background ===="	
	@godoc -http=:6060

setup:
	@go get -u golang.org/x/tools/cmd/cover
	@go get -u github.com/tools/godep 
	@go get -u github.com/pierrre/gotestcover
	@go get -u github.com/kisielk/errcheck
	@go get -u golang.org/x/text/encoding/charmap

deploy: test
	@echo "=== deploy ==="
	@git push wints

run-dev:
	@echo "==== run-dev ===="
	@godep go run main.go --fake-mailer