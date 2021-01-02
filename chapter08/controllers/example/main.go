package main

import(
	"fmt"
	"net/http"
	"github.com/lopes-leandro/learn-go/chapter08/controllers"
)

func main()  {
	
	storage := controllers.MemStorage{}
	c := controllers.New(&storage)
	http.HandleFunc("/get", c.GetValue(false))
	http.HandleFunc("/get/default", c.GetValue(true))
	http.HandleFunc("/set", c.SetValue)

	fmt.Println("Executando na porta :3333")

	err := http.ListenAndServe(":3333", nil)
	panic(err)
}