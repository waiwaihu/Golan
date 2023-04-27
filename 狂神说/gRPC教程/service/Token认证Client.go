package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

/*
grpc提供我们的一个接口，有两个方法，接口位于credentials包下，这个接口需要客户端来实现

	type PerRPCCredentials interface {
		GetRequestMetadada(ctx context.Context,uri ...string) (map[string]string,error)
		RequireTransportSecurity() bool
	}

第一个方法作用是获取元数据信息，也就是客户端提供的key，value。context 用于控制超时和取消，uri是请求入口的uri
第二个方法的作用是否需要基于TLS认证进行安全传输，如果返回值是true，则必须加上TLS认证，返回值是false则不用
*/

type ClientTokenAuth struct{}

func (c ClientTokenAuth) GetRequestMetadada(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "kuangshen",
		"appKey": "123456789",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	//开启安全验证，直接返回true
	return false
}

// 不需要安全连接函数
func Tclient() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	//这儿new报错，暂时隐藏，后续找到解决方法在打开

	//连接到service端，此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:8082", opts...)
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

// 加密
func Ctoken() {
	//客户端加密验证
	creds, _ := credentials.NewClientTLSFromFile("service/key/test.pem", "*.kuangstudy.com")
	//真实地址是浏览器去获取地址

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	//opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	//连接到service端，此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:8082", opts...)
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
