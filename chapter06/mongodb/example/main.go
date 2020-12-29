package main

import (
	"fmt"
	"os"

	"github.com/lopes-leandro/learn-go/chapter06/mongodb"
)

func main() {
	if err := mongodb.Exec(fmt.Sprintf("mongodb://%s:%s@localhost:27017",
		os.Getenv("MONGODBUSERNAME"),
		os.Getenv("MONGODBPASSWORD"))); err != nil {
		panic(err)
	}
}
