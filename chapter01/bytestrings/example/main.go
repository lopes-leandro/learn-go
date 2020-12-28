package main

import (
	"github.com/lopes-leandro/learn-go/chapter01/bytestrings"	
)

func main()  {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}	

	//cada um desses imprime em stdout
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}