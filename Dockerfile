# Build the Go binary
FROM golang:1.23-alpine AS goapp

WORKDIR /app

RUN apk add --no-cache make curl
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY Makefile go.mod go.sum ./
RUN go mod download

COPY main.go  .
COPY internal/ internal/

RUN make all
RUN go build -o ./goapp

# Build the final image
FROM alpine:latest as release
COPY --from=goapp /app/goapp /goapp

WORKDIR /app
CMD ["/goapp", "serve"]
