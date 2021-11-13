FROM golang:1.17 as builder

WORKDIR /go/src/app

COPY server.go server.go
RUN go build server.go

FROM debian:buster-slim

COPY --from=builder /go/src/app/server /usr/local/bin/unstable-http-server
ENTRYPOINT ["unstable-http-server"]

