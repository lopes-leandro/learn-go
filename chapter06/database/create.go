package database

import (
	"database/sql"

	//importamos bibliotecas com suporte para database/sql
	_ "github.com/go-sql-driver/mysql"
)

//Create cria uma tabela chamada
// example e a preenche
func Create(db *sql.DB) error {
	//cria o banco de dados
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name , created) values ("Aaron", NOW())`); err != nil {
		return err
	}

	return nil
}
