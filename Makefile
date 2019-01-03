GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=perf_collector

build:
	GOOS=linux  GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_linux -v
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_macos -v

test-deps:
	go get -v -t -d ./...

test-circleci: test-deps
	go test -race -covermode=atomic -coverprofile=coverage.txt ./...
