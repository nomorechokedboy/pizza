FROM golang:1.19-alpine AS builder

WORKDIR /build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY apps/api/go.mod apps/api/go.sum ./
RUN go mod download

COPY apps/api/src ./src
COPY apps/api/main.go .

RUN go install github.com/swaggo/swag/cmd/swag && \
        $GOPATH/bin/swag init && \
        go build -ldflags="-s -w" -o api .

FROM scratch

COPY --from=builder ["/build/api", "/"]

ENTRYPOINT ["/api"]
