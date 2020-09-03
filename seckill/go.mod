module github.com/WindNotStop/seckill/seckill

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/go-redis/redis/v7 v7.4.0
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/store/redis/v2 v2.9.1 // indirect
	google.golang.org/protobuf v1.25.0
)
