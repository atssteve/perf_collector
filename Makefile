build:
	go mod tidy
	go build -o perf_collector main.go 

run: build
	./perf_collector

test-deps:
	go get -v

test-circleci: test-deps
	go test -race -covermode=atomic -coverprofile=coverage.txt ./...