package handler

import (
	"context"
	"errors"
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
	_, err := s.Rkv.Get(arg).Int()
	if err != nil {
		log.Error(err.Error())
		return errors.New("该秒杀商品不存在")
	}

	if sold, err := s.Rkv.Get(arg + "_sold").Int(); err != nil {
		log.Error(err.Error())
		return err
	} else if sold == 1 {
		return errors.New("活动已结束，请关注后续通知~")
	}

	res, err := s.Rkv.LPop(GoodsName + "_store").Result()
	if err != nil {
		log.Error(err.Error())
		if len, _ := s.Rkv.LLen(GoodsName + "_store").Result(); len == 0 {
			expiration := time.Duration(EndTime.UnixNano() - time.Now().UnixNano())
			s.Rkv.Set(GoodsName+"_sold", 1, expiration)
			return errors.New("抱歉！手慢了！")
		}
		return err
	}
	log.Info(res)

	rsp.Msg = "恭喜！抢到了！"
	return nil
}
