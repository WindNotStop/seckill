package handler

import (
	"context"
	pb "github.com/WindNotStop/seckill/seckill/proto"
	"github.com/go-redis/redis/v7"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"
	"strconv"
	"time"
)

type Seckill struct {
	Client client.Client
	Rkv    *redis.Client
}

// Call is a single request handler called via client.Call or the generated client code
func (s *Seckill) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	arg := req.Name
	num, err := s.Rkv.Get("num").Int()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	for i := 0; i < num; i++ {
		res, err := s.Rkv.SetNX(arg+strconv.Itoa(i), "sold", time.Minute).Result()
		if err != nil {
			log.Error(err.Error())
			return err
		}
		if res {
			rsp.Msg = "恭喜！抢到了！"
			return nil
		}
	}
	rsp.Msg = "抱歉！手慢了！"
	return nil
}
