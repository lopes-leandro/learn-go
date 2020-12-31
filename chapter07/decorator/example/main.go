package main

import (
	"github.com/lopes-leandro/learn-go/chapter07/decorator"
)

func main()  {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}