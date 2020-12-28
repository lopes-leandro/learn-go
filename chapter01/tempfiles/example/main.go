package main

import "github.com/lopes-leandro/learn-go/chapter01/tempfiles"

func main()  {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}