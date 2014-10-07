.PHONY: server lint

server:
	go run main.go

lint:
	gofmt -s -l -w *.go naveed/**.go
