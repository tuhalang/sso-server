# Build stage
FROM golang:1.19-alpine3.16 AS builder
ENV GO111MODULE on
RUN go version
COPY . /src
WORKDIR /src
RUN go mod download
RUN go build -o main cmd/main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /src/main .

EXPOSE 8080
ENTRYPOINT [ "/app/main" ]