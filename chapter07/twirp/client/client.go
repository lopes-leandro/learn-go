package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lopes-leandro/learn-go/chapter07/twirp/rpc/greeter"
)

func main() {

	//Você pode colocar um cliente customizado para controles
	// mais rígidos em tempos limite e assim por diante.
	client := greeter.NewGreeterServiceProtobufClient("http://localhost:4444", &http.Client{})

	ctx := context.Background()
	req := greeter.GreetRequest{Greeting: "Hello", Name: "Reader"}
	resp, err := client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}

	fmt.Print(resp)

	req.Greeting = "Goodbye"
	res, err := client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
