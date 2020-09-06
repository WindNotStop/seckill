package handler

import (

	"context"
	"errors"
	"strconv"
	"time"

	. "github.com/WindNotStop/seckill/seckill/config"
	pb "github.com/WindNotStop/seckill/seckill/seckill/proto"

	"github.com/go-redis/redis/v7"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"

)

type Seckill struct {
	Client client.Client
	Rkv    *redis.Client
}


func (s *Seckill) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Info("call")
	if time.Now().Before(StartTime) {
		return errors.New("活动尚未开始，请耐心等待~")
	}
	if time.Now().After(EndTime) {
		return errors.New("活动已结束，请关注后续通知~")
	}

	arg := req.Name
	num, err := s.Rkv.Get(arg).Int()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if sold, err := s.Rkv.Get(arg+"sold").Int();err != nil {
		log.Error(err.Error())
		return err
	}else if sold == 1{
		return errors.New("活动已结束，请关注后续通知~")
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
	expiration := time.Duration(EndTime.UnixNano()-time.Now().UnixNano())
	s.Rkv.Set(GoodsName + "sold", 1, expiration)
	rsp.Msg = "抱歉！手慢了！"

	return nil
}
