package main

import (
	"context"
	"fmt"
	service "geektime/week4/api/demo2/service/v1"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *service.HelloRequest) (*service.HelloReply, error) {
	return &service.HelloReply{Message: "hello" + in.Name}, nil
}

func main() {
	// 监听本地端口
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("端口监听失败：%s", err)
		return
	}

	// 创建gRPC服务
	s := grpc.NewServer()
	// 注册服务
	service.RegisterGreeterServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败：%s", err)
		return
	}

}
