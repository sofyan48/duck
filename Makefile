# Build And Development
test:
	@go test -v -cover -covermode=atomic ./...

build:
	@go build -o bgin main.go

run:
	@go run main.go

.PHONY: test depend build  run