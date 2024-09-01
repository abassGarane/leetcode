.PHONY: build, run, test
build:
	go build .

run: build
	go run .

test:
	@go test -cover -v ./...
