package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc"
	pb "my-simple-test/helloworld/pb"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	// 调用服务端的SayHello
	hello := &pb.HelloRequest{Name: "huang"}
	b, _ :=proto.Marshal(hello)
	fmt.Printf("%b\n", b)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "q1mi"})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	fmt.Printf("Greeting: %s !\n", r.Message)
}
