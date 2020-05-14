package main

import (
	"context"
	"fmt"
	"github.com/PhillipHuang2017/grpc_learn/hello"
	"log"
	"net"
	"google.golang.org/grpc"
)

var(
	ip = "127.0.0.1"
	port = "50051"
	addr = fmt.Sprintf("%s:%s", ip, port)
)

type Service struct {
	hello.UnimplementedHelloServer  // 实现了接口的结构体，但是里面啥也没有，要继承一下然后重载
}

// 函数声明可以在生成的go文件中找到，重写UnimplementedHelloServer的接口方法即可
func (s *Service) EchoHello(ctx context.Context, req *hello.HELLO_REQUEST) (*hello.HELLO_RESPONSE, error) {
	name := req.GetName()
	res := &hello.HELLO_RESPONSE{
		Message: fmt.Sprintf("Hello! %s!\n", name),
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	// 创建新的grpc服务
	s := grpc.NewServer()
	// 将服务注册到s上
	hello.RegisterHelloServer(s, &Service{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

