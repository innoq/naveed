FROM golang:alpine3.11 as base

WORKDIR /app
COPY . /app

ENV CGO_ENABLED=0
RUN go build -o bin/naveed

FROM busybox

COPY --from=base /app/bin/naveed /usr/bin/naveed
COPY --from=base /app/naveed.ini /

EXPOSE 8465/tcp
USER nobody
CMD ["/usr/bin/naveed"]
