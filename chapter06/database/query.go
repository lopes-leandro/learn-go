package database

import (
	"database/sql"
	"fmt"

	//importamos bibliotecas com suporte para database/sql
	_ "github.com/go-sql-driver/mysql"
)

//Query pega uma nova conexão,
// cria tabelas e depois as descarta
// e emite algumas consultas
func Query(db *sql.DB, name string) error {

	rows, err := db.Query("SELECT name, created FROM example WHERE name=?", name)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var e Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}

	return rows.Err()

}
