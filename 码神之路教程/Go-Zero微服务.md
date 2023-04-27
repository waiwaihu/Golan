# Go-Zero微服务

> 容器是微服务架构的绝佳示例，现代云原生应用使用容器来构建微服务

架构图来源：https://github.com/dotnet-architecture/eShopOnContainers

![img](https://github.com/dotnet-architecture/eShopOnContainers/raw/dev/img/eShopOnContainers-architecture.png)

# 官网

https://go-zero.dev/cn/

# 1、环境准备

在正式进入实际开发之前，我们需要做一些准备工作，比如：Go环境的安装，grpc代码生成使用的工具安装， 必备工具Goctl的安装，Golang环境配置等，本节将包含以下小节：

## golang安装

```bash
# 1、linux 安装Go
下载Go for Linux,解压压缩包至/usr/local
$ tar -C /usr/local -xzf go1.15.8.linux-amd64.tar.gz

添加/usr/local/go/bin到环境变量
$ $HOME/.profile
$ export PATH=$PATH:/usr/local/go/bin
$ source $HOME/.profile

验证安装结果
$ go version
go version go1.15.1 linux/amd64


# 2、Windows安装Go
下载并安装Go for Windows 验证安装结果
$ go version
go version go1.15.1 windows/amd64
```

## go module配置

```bash
$ go env -w GO111MODULE="on"
$ go env -w GOPROXY=https://goproxy.cn

如果目录不为空或者/dev/null，请跳过。
$ go env -w GOMODCACHE=$GOPATH/pkg/mod
```

## goctl安装

goctl（官方建议读 go control）是go-zero微服务框架下的代码生成工具，使用goctl可显著提升开发效率，让开发人员将时间重点放在业务快发上，其功能有：

- api服务生成
- rpc服务生成
- model代码生成
- 模板管理

```bash
# Go 1.15 及之前版本
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/zeromicro/go-zero/tools/goctl@latest

# Go 1.16 及以后版本
GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest

#设置完go代理
go install github.com/zeromicro/go-zero/tools/goctl@v1.4.4
通过此命令可以将goctl工具安装到 $GOPATH/bin 目录下

$ goctl -v
goctl version 1.1.4 darwin/amd64
```

## protoc & protoc-gen-go安装

```bash
#protoc是一款用C++编写的工具，其可以将proto文件翻译为指定语言的代码。在go-zero的微服务中，我们采用grpc进行服务间的通信，而grpc的编写就需要用到protoc和翻译成go语言rpc stub代码的插件protoc-gen-go。
```

### 方式一：goctl一键安装

```bash
#快捷命令
$ goctl env check -i -f --verbose  
```

### 方式二： Homebrew（macOS）

```bash
$ brew install protobuf protoc-gen-go protoc-gen-go-grpc
$ protoc --version
libprotoc x.x.x
```

## 其他环境

- [etcd](https://etcd.io/docs/v3.5/)
- [redis](https://redis.io/)
- [mysql](https://www.mysql.com/)

[其他](https://go-zero.dev/cn/docs/prepare/prepare-other)

# 单体服务

## 1、HelloWorld

```bash
E:\GoWork\src\GoProject> mkdir helloworld
E:\GoWork\src\GoProject> cd helloworld
E:\GoWork\src\GoProject> goctl api new hello
E:\GoWork\src\GoProject> cd hello
E:\GoWork\src\GoProject> go mod tidy
```

