package main

import "github.com/lopes-leandro/learn-go/chapter07/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
