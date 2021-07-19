# go build
FROM golang:alpine3.11
RUN apk --update add \
    # git \
    # glide \
    # upx \
    libc-dev \
    gcc \
    && rm -rf /var/lib/apt/lists/* \
    && rm /var/cache/apk/*

WORKDIR /app
COPY . /app

RUN go test ./... -v
# RUN go build -o bin/naveed && /usr/bin/upx /app/bin/naveed
RUN go build -o bin/naveed

# RUN chmod a+x /app/bin

# COPY patch/latexmkrc /root/.latexmkrc
# COPY test /test
# COPY validations /validations
# COPY templates /templates
# COPY credentials.conf /credentials.conf

# ENTRYPOINT ["/app/bin/naveed"]
# RUN /usr/local/bin/box version
# COPY --from=build /app/bin/naveed /app/naveed

# RUN /app/bin/naveed -version

CMD ["/app/bin/naveed"]


