FROM golang:1.19-alpine
WORKDIR /app

COPY $PWD ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
CMD swag init; go run main.go

