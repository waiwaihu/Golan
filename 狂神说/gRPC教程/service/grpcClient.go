package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func Gclient() {
	//连接到service端，此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()

	//建立连接
	client := NewSayHelloClient(conn)

	//执行rpc调用(这个方法在服务端来实现并返回结果)
	resp, err := client.SayHello(context.Background(), &HelloRequest{RequestName: "kuangshen"})
	if err != nil {
		log.Fatalf("客户端rpc调用失败：%v", err)
	}
	fmt.Println(resp.GetResponseMsg())
}
