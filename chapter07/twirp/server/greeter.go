package main

import (
	"context"
	"fmt"

	"github.com/lopes-leandro/learn-go/chapter07/twirp/rpc/greeter"
)

//Greeter implementa a interface
// gerada pelo protoc
type Greeter struct {
	Exclaim bool
}

//Greet implementa twirp Greet
func (g *Greeter) Greet(ctx context.Context, r *greeter.GreetRequest) (*greeter.GreetResponse, error) {

	msg := fmt.Sprintf("%s %s", r.GetGreeting(), r.GetName())
	if g.Exclaim {
		msg += "!"
	} else {
		msg += "."
	}

	return &greeter.GreetResponse{Response: msg}, nil
}
