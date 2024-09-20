.PHONY: build, run, test
build:
	@CGO_ENABLED=0 go build -o bin/leetcode -ldflags "-s -w" main.go

run: build
	@./bin/leetcode

test:
	@go test -cover -v ./...
