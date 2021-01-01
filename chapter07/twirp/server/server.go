package main

import (
	"fmt"
	"net/http"

	"github.com/lopes-leandro/learn-go/chapter07/twirp/rpc/greeter"
)

func main() {

	server := &Greeter{}
	twirpHandler := greeter.NewGreeterServiceServer(server, nil)

	fmt.Println("Executando na porta :4444")
	http.ListenAndServe(":4444", twirpHandler)

}
