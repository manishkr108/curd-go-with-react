# Stage 1: Build the Go application
FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /myapp

# Stage 2: Create a minimal image to run the Go application
FROM alpine:latest

WORKDIR /root/

COPY --from=build /myapp .

CMD ["./myapp"]
