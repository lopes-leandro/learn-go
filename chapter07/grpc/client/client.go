package main

import(
	"context"
	"fmt"
	"github.com/lopes-leandro/learn-go/chapter07/grpc/greeter"
	"google.golang.org/grpc"
)

func main()  {
	conn, err := grpc.Dial(":4444", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := greeter.NewGreeterServiceClient(conn)

	ctx := context.Background()
	req := greeter.GreetRequest{Greeting: "Hello", Name: "Reader"}
	res, err := client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	req.Greeting = "Goodbye"
	res, err = client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}