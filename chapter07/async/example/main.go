package main

import (
	"fmt"
	"net/http"
	"github.com/lopes-leandro/learn-go/chapter07/async"
)

func main()  {
	urls := []string{
		"https://www.google.com",
		"https://golang.org",
		"https://www.github.com",
		"http://www.tuiaseguros.com.br",
		"http://www.amazon.com",
		"http://www.abobrinha.com",
	}
	c := async.NewClient(http.DefaultClient, len(urls))
	async.FetchAll(urls, c)

	for i := 0; i < len(urls); i++ {
		select{
		case resp := <- c.Resp:
			fmt.Printf("Status recebido para %s: %d\n", resp.Request.URL, resp.StatusCode)
		case err := <- c.Err:
			fmt.Printf("Erro recebido: %s\n", err)
		}
	}
}