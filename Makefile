SHELL := /bin/bash
CI_DOCKER_IMAGE=<your docker repo>

# build images from all services
.PHONY: build-image-local
build-image-local: build-linux
	docker build --platform linux/amd64 -f build/local.Dockerfile -t ${CI_DOCKER_IMAGE}/${service}:${tag} . --build-arg service=${service} --build-arg BOT_USER=${BOT_USER} --build-arg BOT_PRIVATE_TOKEN=${BOT_PRIVATE_TOKEN}

.PHONY: build-linux
build-linux: download-go-mod
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -o out/cmd/main cmd/main.go

test:
	echo "Test success"

.PHONY: download-go-mod
download-go-mod: go.mod
	go mod download all

