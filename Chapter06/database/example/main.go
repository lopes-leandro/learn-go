package main

import (
	"github.com/lopes-leandro/learn-go/chapter06/database"
	//importamos bibliotecas com suporte para database/sql
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	if err := database.Exec(db); err != nil {
		panic(err)
	}
}