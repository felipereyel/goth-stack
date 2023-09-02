# Build the Go binary
FROM golang:1.20-alpine AS goapp
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go  .
COPY pkgs/ pkgs/
RUN go build -o ./goapp

# Build the final image
FROM alpine:latest as release
COPY --from=goapp /app/goapp /goapp

WORKDIR /app
EXPOSE 8080
CMD ["/goapp"]
