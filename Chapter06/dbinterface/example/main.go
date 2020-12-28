package main

import (
	"github.com/lopes-leandro/learn-go/chapter06/database"
	"github.com/lopes-leandro/learn-go/chapter06/dbinterface"
	//importamos bibliotecas com suporte para database/sql
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	defer tx.Rollback()

	if err := dbinterface.Exec(tx); err != nil {
		panic(err)
	}

	if err := tx.;err != nil {
		return nil, err
	}
	
}