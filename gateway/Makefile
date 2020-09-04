
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config

.PHONY: proto
proto:
	protoc   --go_out=${MODIFY}:.   --go-grpc_out=${MODIFY}:.   --go_opt=paths=source_relative   --go-grpc_opt=paths=source_relative   proto/seckill.proto
	protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:. proto/seckill.proto

.PHONY: build
build: proto

	go build -o gateway *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t taenyang/gateway:latest


