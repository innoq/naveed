.PHONY: server test lint

server:
	go run main.go

test:
	go test -v ./...

lint:
	gofmt -s -l -w *.go naveed/**.go
