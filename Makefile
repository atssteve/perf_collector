build:
	go mod tidy
	go build -o perf_collector main.go 

run: build
	./perf_collector

test-deps:
	go get -v

test-circleci: test-deps
	go test -race -covermode=atomic -coverprofile=coverage.txt ./...

docker-pull:
	docker pull golang

docker-run:
	docker run -v $(CURDIR):/go/perf_collector -it golang /bin/sh -c 'cd perf_collector && make run'