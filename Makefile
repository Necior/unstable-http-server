server: server.go
	go build server.go

.PHONY:
format: server.go
	go fmt server.go

.PHONY:
run: server
	./server

.PHONY:
publish-dev-docker-image: server.go Dockerfile
	docker buildx build --platform linux/arm64,linux/amd64 -t necior/unstable-http-server:dev --push .

