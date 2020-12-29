package main

import "github.com/lopes-leandro/learn-go/chapter06/storage"

func main()  {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}