package main

import (

	"time"

	"github.com/WindNotStop/seckill/seckill/handler"
	"github.com/WindNotStop/seckill/seckill/ratelimiter"
	pb "github.com/WindNotStop/seckill/seckill/seckill/proto"

	"github.com/go-redis/redis/v7"
	"github.com/juju/ratelimit"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

)

var(
	serviceName = "seckill"
	serviceVersion = "v1"

	redisURL = "redis://redis:6379"

	fillInterval = time.Millisecond
	capacity int64 = 1000
)

func main() {

	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(serviceVersion),
		micro.WrapHandler(ratelimiter.NewHandlerWrapper(ratelimit.NewBucket(fillInterval,capacity),false)),
	)
	service.Init()

	nodes := []string{redisURL}
	redisOptions, err := redis.ParseURL(nodes[0])
	if err != nil {
		log.Error(err.Error())
	}
	rkv := redis.NewClient(redisOptions)
	rkv.Set("num", 10, 24*time.Hour)

	pb.RegisterSeckillHandler(
		service.Server(),
		&handler.Seckill{
			Client: service.Client(),
			Rkv:    rkv,
		},
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
