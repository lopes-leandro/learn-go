package database

import (
	"database/sql"

	//importamos bibliotecas com suporte para database/sql
	_ "github.com/go-sql-driver/mysql"
)

//Exec substitui o Exec da receita anterior
func Exec(db *sql.DB) error {
	//erro não detectado na limpeza,
	// mas sempre queremos limpar
	defer db.Exec("DROP TABLE example")

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db, "Aaron"); err != nil {
		return err
	}

	return nil
}
