package main

import(
	"github.com/lopes-leandro/learn-go/chapter06/pools"
)

func main()  {
	if err := pools.ExecWithTimeout(); err != nil {
		panic(err)
	}
}