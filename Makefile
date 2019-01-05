build:
	go mod tidy
	go build -o perf_collector main.go 

test-deps:
	get get -v

test-circleci: test-deps
	go test -race -covermode=atomic -coverprofile=coverage.txt ./...