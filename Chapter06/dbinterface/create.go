package dbinterface

import (
		//importamos bibliotecas com suporte para database/sql
		_ "github.com/go-sql-driver/mysql"
)

//Create cria uma tabela chamada
// example e a preenche
func Create(db DB) error  {
	
	//cria a tabela
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME"); err != nil {
		return err
	}

	//insere um registro
	if _, err := db.Exec(`INSERT INTO (name, created) VALUES("Aaron", Now())`); err != nil {
		return err
	}

	return nil
}