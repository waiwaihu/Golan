package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func Gtclient() {
	//客户端加密验证
	creds, _ := credentials.NewClientTLSFromFile("service/key/test.pem", "*.kuangstudy.com") //真实是浏览器去获取地址

	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithTransportCredentials(creds))
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
