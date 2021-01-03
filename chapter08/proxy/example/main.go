package main

import (
	"fmt"
	"net/http"
	"github.com/lopes-leandro/learn-go/chapter08/proxy"
)

func main()  {
	
	p := &proxy.Proxy{
		Client: http.DefaultClient,
		BaseURL: "http://www.golang.org",
	}

	http.Handle("/", p)
	fmt.Println("Executando na porta :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
	
}