
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:

	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:.  seckill/proto/seckill.proto


.PHONY: build
build:

	go build -o seckill *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t taenyang/seckill:latest


