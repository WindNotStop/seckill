
GOPATH:=$(shell go env GOPATH)

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t taenyang/user:latest

.PHONY: push
push:
	docker push taenyang/user:latest


