package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Service struct {
	UnimplementedSayHelloServer
}

func (s *Service) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	fmt.Printf("hello" + req.RequestName)
	return &HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func Gservice() {
	//开启端口
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("开启端口监听失败:", err)
		return
	}

	//创建grpc服务
	grpcServer := grpc.NewServer() //不加密

	//在grpc服务端中去注册我们自己编写的服务
	RegisterSayHelloServer(grpcServer, &Service{})

	//启动服务
	err1 := grpcServer.Serve(listen)
	//err1 = grpcServer.Serve(listen)
	if err1 != nil {
		fmt.Println("启动grpc服务失败：", err1)
		return
	}
}
