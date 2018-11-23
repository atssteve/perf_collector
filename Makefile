build:
	go build -o perf_collector main.go 

test-deps:
	go get -v -t -d ./...

test-circleci: test-deps
	go test -race -covermode=atomic -coverprofile=coverage.txt ./...