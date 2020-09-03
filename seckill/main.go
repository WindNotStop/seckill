package main

import (
	"github.com/WindNotStop/seckill/seckill/handler"
	pb "github.com/WindNotStop/seckill/seckill/proto"
	"github.com/go-redis/redis/v7"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main(){
	service := micro.NewService(
		micro.Name("seckill"),
		micro.Version("v1"),
	)
	service.Init()

	nodes := []string{"redis://127.0.0.1:6379"}

	redisOptions, err := redis.ParseURL(nodes[0])
	if err != nil {
		//Backwards compatibility
		redisOptions = &redis.Options{
			Addr:     nodes[0],
			Password: "", // no password set
			DB:       0,  // use default DB
		}
	}

	rkv := redis.NewClient(redisOptions)

	pb.RegisterSeckillHandler(
		service.Server(),
		&handler.Seckill{
			Client:service.Client(),
			Rkv:rkv,
		},
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

