package service

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"net"
)

type Servicew struct {
	UnimplementedSayHelloServer
}

func (s *Servicew) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	//获取元数据信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未传输Token")
	}
	var appId string
	var appKey string
	if v, ok := md["appId"]; ok {
		appId = v[0]
	}
	if v, ok := md["appKey"]; ok {
		appKey = v[0]
	}
	//用户appId
	if appId != "kuangshen" || appKey != "123456789" {
		return nil, errors.New("Token 不正确!!!")
	}

	fmt.Printf("hello" + req.RequestName)
	return &HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func Gservicew() {
	//开启端口
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("开启端口监听失败:", err)
		return
	}

	//创建grpc服务
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials())) //不加密

	//在grpc服务端中去注册我们自己编写的服务
	RegisterSayHelloServer(grpcServer, &Servicew{})

	//启动服务
	err1 := grpcServer.Serve(listen)
	//err1 = grpcServer.Serve(listen)
	if err1 != nil {
		fmt.Println("启动grpc服务失败：", err1)
		return
	}
}

func Gserviceq() {
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
	RegisterSayHelloServer(grpcServer, &Servicew{})

	//启动服务
	err1 := grpcServer.Serve(listen)
	//err1 = grpcServer.Serve(listen)
	if err1 != nil {
		fmt.Println("启动grpc服务失败：", err1)
		return
	}
}