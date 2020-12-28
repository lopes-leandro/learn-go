package main

import (
	"bytes"
	"fmt"
	"github.com/lopes-leandro/learn-go"
)

func main()  {
	in := bytes.NewReader([]byte("example\t"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer = ", out.String())

	fmt.Print("stdout on PipeExample = ")

	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}

}