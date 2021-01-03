package main

import (
	"fmt"
	"net/http"

	"github.com/lopes-leandro/learn-go/chapter08/negociate"
)

func main() {

	http.HandleFunc("/", negociate.Handler)
	fmt.Println("Executando na porta :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)

}
