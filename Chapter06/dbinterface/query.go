package dbinterface

import (
	"fmt"

	"github.com/lopes-leandro/learn-go/chapter06/database"
)

//Query pega uma nova conexão,
// cria tabelas e depois as descarta
// e emite algumas consultas
func Query(db DB) error {

	name := "Aaron"
	rows, err := db.Query("SELECT name, created FROM example WHERE name=?", name)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var e database.Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Resultados:\n\tNome: %s\n\tCreated: %v\n", e.Name, e.Created)
	}

	return rows.Err()
}
