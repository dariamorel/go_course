package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"project/proto"
	"time"
)

func main() {
	name := flag.String("name", "world", "Name to greet")

	flag.Parse()

	conn, err := grpc.NewClient("0.0.0.0:4567", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.SayHello(ctx, &proto.HelloRequest{Name: *name})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Greeting: %s", res.GetMessage())
}
