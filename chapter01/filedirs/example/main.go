package main

import (
	filedirs "github.com/lopes-leandro/learn-go/chapter01/fieldirs"
)

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.CapitalizerExample(); err != nil {
		panic(err)
	}
}
