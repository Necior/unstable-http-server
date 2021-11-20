FROM golang:1.17 as builder

WORKDIR /go/src/app

COPY unstable-http-server.go unstable-http-server.go
COPY VERSION VERSION
RUN go build unstable-http-server.go

FROM debian:buster-slim

COPY --from=builder /go/src/app/unstable-http-server /usr/local/bin/unstable-http-server
ENTRYPOINT ["unstable-http-server"]

