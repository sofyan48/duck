# Build And Development
test:
	@go test -v -cover -covermode=atomic ./...

build:
	@go build -o duck src/main.go

run:
	@go run src/main.go

.PHONY: test depend build  run