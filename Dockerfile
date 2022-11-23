FROM golang:1.19-alpine3.16 AS build

WORKDIR /src
COPY . /src
RUN go mod download

RUN go build -o services server.go
RUN chmod +x services
RUN go clean -modcache

FROM alpine:3.16 as binary
WORKDIR /app
COPY --from=build /src/services /app
CMD ["/app/services"]