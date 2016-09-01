## simple makefile to log workflow
.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:)

all: install test


build:
	@echo "==== go build ===="
	@go build $(GOFLAGS) ./...

install:
	@echo "==== install ===="
	@go install -ldflags "-X main.Version=${VERSION}" $(GOFLAGS) .

vet:
	@echo "==== go vet ===="
	@go vet ./...

lint:
	@echo "==== go lint ===="
	@golint ./**/*.go

test: build vet lint
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

assets:
	@echo "=== production level assets ==="
	@gulp assets --production

deploy: install assets
	@git add assets
	@git commit -m "production level assets"
	@echo "=== deploy ==="
	@git push wints