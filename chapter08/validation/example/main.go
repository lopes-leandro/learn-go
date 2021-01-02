package main

import (
	"fmt"
	"net/http"
	"github.com/lopes-leandro/learn-go/chapter08/validation"
)

func main()  {
	
	c := validation.New()
	http.HandleFunc("/", c.Process)
	fmt.Println("Executando na porta :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}