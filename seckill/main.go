package main

import (
	"github.com/WindNotStop/seckill/seckill/handler"
	pb "github.com/WindNotStop/seckill/seckill/seckill/proto"
	"github.com/go-redis/redis/v7"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	"time"
)

func main() {
	service := micro.NewService(
		micro.Name("seckill"),
		micro.Version("v1"),
	)
	service.Init()

	nodes := []string{"redis://redis:6379"}
	//nodes := []string{"redis://localhost:6379"}

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
