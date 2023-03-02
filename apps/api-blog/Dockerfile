FROM golang:1.19-alpine AS builder

WORKDIR /build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY apps/api-blog/go.mod apps/api-blog/go.sum ./
RUN go mod download && \
    go install github.com/cosmtrek/air@latest && \
    go install github.com/swaggo/swag/cmd/swag

COPY apps/api-blog/api ./src 
COPY apps/api-blogs/pkg ./src
COPY apps/api-blog/main.go apps/api/.air.toml ./

RUN $GOPATH/bin/swag init && \
    go build -ldflags="-s -w" -o api .

FROM scratch

COPY --from=builder ["/build/api", "/"]

ENTRYPOINT ["/api"]

