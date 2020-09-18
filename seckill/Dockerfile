FROM debian
ADD seckill-service /seckill-service
ENTRYPOINT [ "/seckill-service","--server_address=:9090","--mode=k8s"]
