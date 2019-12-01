# Build And Development
test:
	@go test -v -cover -covermode=atomic ./...

depend:
	@go get github.com/meongbego/bgin

build:
	@go build -o bgin main.go

run:
	@go run main.go

.PHONY: test depend build  run