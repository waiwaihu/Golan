package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

type Servicel struct {
	UnimplementedSayHelloServer
}

func (s *Servicel) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	fmt.Printf("hello" + req.RequestName)
	return &HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func Gservicel() {
	//开启加密认证,不用要相对路径，我这里是测试一下相对路径
	//自签名文件和私钥文件
	creds, _ := credentials.NewServerTLSFromFile("service/key/test.pem", "service/key/test.key")

	//开启端口
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("开启端口监听失败:", err)
		return
	}

	//创建grpc服务
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	//在grpc服务端中去注册我们自己编写的服务
	RegisterSayHelloServer(grpcServer, &Servicel{})

	//启动服务
	err1 := grpcServer.Serve(listen)
	//err1 = grpcServer.Serve(listen)
	if err1 != nil {
		fmt.Println("启动grpc服务失败：", err1)
		return
	}
}
