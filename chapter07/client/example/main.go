package main

import "github.com/lopes-leandro/learn-go/chapter07/client"

func main() {

	//seguro e opcional!
	cli := client.Setup(true, false)
	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}

	if err := client.DoOps(cli); err != nil {
		panic(err)
	}

	c := client.Controller{Client: cli}
	if err := c.DoOps(); err != nil {
		panic(err)
	}

	//seguro e não opcional
	client.Setup(true, true)
	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}

}
