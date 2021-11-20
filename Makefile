version := $(shell cat VERSION)

unstable-http-server: unstable-http-server.go VERSION
	go build unstable-http-server.go

.PHONY:
format: unstable-http-server.go
	go fmt unstable-http-server.go

.PHONY:
run: unstable-http-server
	./unstable-http-server

.PHONY:
publish-dev-docker-image: unstable-http-server.go Dockerfile VERSION
	docker buildx build --platform linux/arm64,linux/amd64 -t necior/unstable-http-server:$(version) --push .

