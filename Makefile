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
	@godep go clean $(GOFLAGS) -i ./...

doc:
	@echo "==== go doc is running. Can be moved to background ===="	
	@godoc -http=:6060

setup:
	@npm install --save-dev gulp-wrap\
							gulp-declare\
							gulp-concat\
							gulp-uglify\
							gulp-rename\
							gulp-clean-css\
							gulp-htmlmin\
							gulp-handlebars\
							gulp-util
	@go get -u golang.org/x/tools/cmd/cover
	@go get -u github.com/tools/godep 
	@go get -u github.com/pierrre/gotestcover
	@go get -u github.com/kisielk/errcheck
	@go get -u golang.org/x/text/encoding/charmap
	@go get -u github.com/maruel/panicparse/cmd/pp

assets:
	@gulp assets --production
	
deploy: install
	@gulp assets --production
	@git add assets
	@git commit -m "production level assets"	
	@echo "=== deploy ==="
	@git push wints

run-dev:
	@echo "==== run-dev ===="
	@go run -ldflags "-X main.Version=`git rev-parse HEAD`" main.go --fake-mailer
