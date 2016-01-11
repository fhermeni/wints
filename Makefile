## simple makefile to log workflow
.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:)

all: install test

build:
	@echo "==== go build ===="		
	@go build $(GOFLAGS) ./...

install:
	@go get $(GOFLAGS) ./...

vet:
	@echo "==== go vet ===="
	@go vet ./...

lint:
	@echo "==== go lint ===="
	@golint ./**/*.go

test: install vet lint	
	@echo "==== go test ===="
	@go test $(GOFLAGS) ./...	

test-integration: install
	@echo "==== integration test ===="
	@go test -v -tags "integration" ./...

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
	@go test -tags "integration" -run=NONE -bench=. $(GOFLAGS) ./...

clean:
	@go clean $(GOFLAGS) -i ./...

doc:
	@echo "==== go doc is running. Can be moved to background ===="	
	@godoc -http=:6060

setup:
	@go get -u golang.org/x/tools/cmd/cover
	@go get -u github.com/tools/godep 
	@go get -u github.com/pierrre/gotestcover
	@go get -u github.com/shurcooL/vfsgen
	@go get -u github.com/shurcooL/vfsgen/cmd/vfsgendev
	@go get -u github.com/jteeuwen/go-bindata/...
	@go get -u github.com/elazarl/go-bindata-assetfs/...
	@go get -u github.com/kisielk/errcheck

run-dev: 
	@echo "==== dev-run ===="
	@go run wints/wints.go --fake-mailer

assets:
	go run assets/assets_generate.go		