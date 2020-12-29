package main

import(

	"github.com/lopes-leandro/learn-go/chapter06/dbinterface"

	"github.com/lopes-leandro/learn-go/chapter06/database"

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

	//isso não fará nada se o commit for bem sucedido
	defer tx.Rollback()

	if err := dbinterface.Exec(tx); err != nil {
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}
}