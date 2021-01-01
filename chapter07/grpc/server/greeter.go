package main

import (
	"fmt"
	"github.com/lopes-leandro/learn-go/chapter07/grpc/greeter"
	"golang.org/x/net/context"
)

//Greeter implementa a interface
// gerada pelo protoc
type Greeter struct {
	Exclaim bool
}

//Greet implementa grpc Greet
func (g *Greeter) Greet(ctx context.Context, r *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	
	msg := fmt.Sprintf("%s %s", r.GetGreeting(), r.GetName())
	if g.Exclaim {
		msg += "!"
	} else {
		msg += "."
	}

	return &greeter.GreetResponse{Response: msg}, nil
}