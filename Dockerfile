FROM golang:alpine3.11

WORKDIR /app
COPY . /app

RUN go build -o bin/naveed
CMD ["/app/bin/naveed"]
