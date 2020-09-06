package config

import "time"

var(
	ServiceName = "seckill"
	ServiceVersion = "v1"

	RedisURL = "redis://redis:6379"

	GoodsName = "iphone"
	GoodsNum = 10
	loc, _ = time.LoadLocation("Asia/Shanghai")
	StartTime = time.Date(2020, 9, 6, 14, 0, 0 ,0, loc)
	EndTime = time.Date(2020, 9, 6, 15, 0, 0 ,0, loc)

	FillInterval = time.Millisecond
	Capacity int64 = 1000
)
