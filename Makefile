.DEFAULT_GOAL := build

.PHONY:fmt vet build clean test
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build cmd/main.go

clean:
	go clean

test:
	go test -v ./...

run:
	go run cmd/main.go