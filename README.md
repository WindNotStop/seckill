# 基于云原生的秒杀系统
## Introduction
- 容器化应用，微服务架构
- 无缝融入Kubernetes、 Istio，实现服务高可用、弹性扩展、易于观测等

## Architecture
![overview](https://github.com/WindNotStop/seckill/blob/master/arch.png)

## Roadmap
- [x] Grpc-gateway
- [x] 秒杀服务
- [x] Redis集群
- [x] 服务限流
- [ ] Traefik
- [ ] 动态配置
- [ ] Mysql集群
- [ ] 用户管理服务
- [ ] 支付服务
- [ ] URL动态化
- [ ] 前端页面


## Usage
### local
```
#vim /etc/hosts
127.0.0.1       localhost redis seckill

# run redis
docker run --name redis -d -p 6379:6379 redis redis-server --appendonly yes

# run seckill
cd seckill/
go run main.go --server_address=:9090

# run gateway
cd gateway/
go run main.go

#test
curl localhost:8081/v1/seckill?name=iphone
```
### Kubernates
```
# run redis
helm install redis seckill/charts/redis/

# run seckill
helm install seckill seckill/charts/seckill/

# run gateway
helm install gateway seckill/charts/gateway/

# test
curl ip:30080/v1/seckill?name=iphone
```
