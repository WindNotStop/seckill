package main

import (
	"context"
	"flag"
	"net/http"

	gw "github.com/WindNotStop/seckill/gateway/proto"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "seckill:9090", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterSeckillHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
