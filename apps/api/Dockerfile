FROM golang:1.19-alpine AS builder

WORKDIR /build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY apps/api/go.mod apps/api/go.sum ./
RUN go mod download && \
    go install github.com/cosmtrek/air@latest && \
    go install github.com/swaggo/swag/cmd/swag

COPY apps/api/src ./src
COPY apps/api/main.go apps/api/.air.toml ./

RUN $GOPATH/bin/swag init && \
    go build -ldflags="-s -w" -o api .

FROM busybox AS deps

ARG BUSYBOX_VERSION=1.31.0-i686-uclibc
ADD https://busybox.net/downloads/binaries/$BUSYBOX_VERSION/busybox_WGET /wget
RUN chmod a+x /wget

FROM gcr.io/distroless/static-debian11:nonroot AS runtime

COPY --from=deps /wget /usr/bin/wget
COPY --from=builder ["/build/api", "/"]

ENTRYPOINT ["/api"]

