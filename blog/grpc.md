## Grpc
RPC（Remote Procedure Call），即远程过程调用，主要是为了解决在分布式系统中，服务之间的调用问题。

首先在分布式系统中，部署在不同机器上的服务需要协同工作就必须要相互通讯，它们之间的通讯并不能像是本地函数调用一样轻松，本地函数的调用因为是在同一地址空间，所以通过方法栈和参数栈就可以实现相互之间通讯，但部署在不同机器上的服务，并没有共用的地址空间，所以只能通过网络通讯。

HTTP通常是网络传输的首选，但是HTTP比较重，为了优化这个重复编写 httpClient 的操作，可以用代理模式实现调用，在这个代理内部实现 httpClient 请求等一系列繁杂的操作。目前其实有很多 RPC 框架是采用这种设计理念，如 Motan，dubbo 等。
由此可见 RPC 不仅仅是为了解决服务之间的调用，更是为了让服务之间的调用像本地函数调用一样，方便快捷。

### RPC三个部分
网络传输，函数名映射，序列化和反序列化
## Demo
使用 gRPC 主要分为三步：

1. 编写 .proto pb 文件，制定通讯协议。
2. 利用对应插件将 .proto pb文件编译成对应语言的代码。
3. 根据生成的代码编写业务代码。
### Requirements
protobuf
```
go get -u github.com/golang/protobuf/protoc-gen-go
```
grpc
```
go get -u google.golang.org/grpc
```

### 目录结构
```
|—— hello/
    |—— client/
        |—— main.go   // 客户端
    |—— server/
        |—— main.go   // 服务端
|—— proto/
    |—— hello/
        |—— hello.proto   // proto描述文件
        |—— hello.pb.go   // proto编译后文件
```

### Proto
```proto
syntax = "proto3"; // 指定proto版本
package hello;     // 指定默认包名

// 指定golang包名
option go_package = "proto/hello";

// 定义Hello服务
service Hello {
  // 定义SayHello方法
  rpc SayHello(HelloRequest) returns (HelloResponse) {}
}

// HelloRequest 请求结构
message HelloRequest {
  string name = 1;
}

// HelloResponse 响应结构
message HelloResponse {
  string message = 1;
}
```
### 编译proto
```bash
protoc -I proto/ proto/hello/hello.proto --go_out=plugins=grpc:.

```
### 实现服务端
```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc/proto/hello" // 确保proto包路径正确
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:8081"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	log.Println("Listening on " + Address)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

```

### 实现客户端
```go
package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/proto/hello" // 确保proto包路径正确
	"log"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:8081"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	log.Println("Response from server:", res.Message)
}

```


### 测试grpc
```
go run main.go
Listen on 127.0.0.1:8081  //服务端已开启并监听50052端口
```
```
$ go run main.go
Hello gRPC.    // 接收到服务端响应
```

