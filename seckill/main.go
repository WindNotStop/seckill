package main

import (
	"flag"
	"strconv"
	"time"

	. "github.com/WindNotStop/seckill/seckill/config"
	"github.com/WindNotStop/seckill/seckill/handler"
	"github.com/WindNotStop/seckill/seckill/ratelimiter"
	pb "github.com/WindNotStop/seckill/seckill/seckill/proto"

	"github.com/go-redis/redis/v7"
	"github.com/juju/ratelimit"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

var serverAddress = flag.String("server_address",  ":9090", "server_address")
var mode = flag.String("config",  "local", "mode")

func main() {
	flag.Parse()

	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Version(ServiceVersion),
		micro.WrapHandler(ratelimiter.NewHandlerWrapper(ratelimit.NewBucket(FillInterval, Capacity), false)),
	)
	service.Init()
	
	var rkv *redis.Client
	switch *mode {
	case "local":
		nodes := []string{RedisURL}
		redisOptions, err := redis.ParseURL(nodes[0])
		if err != nil {
			log.Error(err.Error())
		}
		rkv = redis.NewClient(redisOptions)
	case "k8s":
		rkv = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    MasterName,
			SentinelAddrs: []string{RedisSentinelAddress},
		})
	default:
		return
	}

	expiration := time.Duration(EndTime.UnixNano() - time.Now().UnixNano())
	rkv.Set(GoodsName, GoodsNum, expiration)
	rkv.Set(GoodsName+"_sold", 0, expiration)
	for i := 0; i < GoodsNum; i++ {
		rkv.RPush(GoodsName+"_store", GoodsName+strconv.Itoa(i))
	}

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
