package main

import (
	"context"
	"fmt"
	"github.com/PhillipHuang2017/grpc_learn/hello"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

var(
	ip = "127.0.0.1"
	port = "50051"
	addr = fmt.Sprintf("%s:%s", ip, port)
	defaultName = "Phillip"
)

func main() {
	// 创建连接，加上Insecure和阻塞的选项，默认是Secure的，但是服务端要开Secture选项才能用Secture去连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial: %v\n", err)
	}
	defer conn.Close()
	client := hello.NewHelloClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.EchoHello(ctx, &hello.HELLO_REQUEST{
		Name: name,
	})
	if err != nil {
		log.Fatalf("could not Hello: %v\n", err)
	}
	log.Printf("Message from server: %v\n", res.GetMessage())
}