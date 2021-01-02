package main

import (
	"fmt"
	"net/http"
	"github.com/lopes-leandro/learn-go/chapter08/handlers"
)

func main()  {
	
	http.HandleFunc("/name", handlers.HelloHandler)
	http.HandleFunc("/greeting", handlers.GreetingHandler)
	fmt.Println("Executando na porta :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}