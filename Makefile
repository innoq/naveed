.PHONY: server test lint

server:
	go run main.go

test:
	cd naveed && go test

lint:
	gofmt -s -l -w *.go naveed/**.go
