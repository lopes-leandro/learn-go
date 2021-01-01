package main

import(
	"fmt"
	"net"
	"github.com/lopes-leandro/learn-go/chapter07/grpc/greeter"
	"google.golang.org/grpc"
)

func main()  {
	
	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServiceServer(grpcServer, &Greeter{Exclaim: true})
	list, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}
	fmt.Println("Executando na porta :4444")
	grpcServer.Serve(list)
}