package dbinterface

import(
	"fmt"

	"github.com/lopes-leandro/learn-go/chapter06/database"
)

//Query pega uma nova conexão,
// cria tabelas e depois as descarta
// e emite algumas consultas
func Query(db DB)  error {
	 name := "Aeron"
	 rows, err := Query("SELECT name, created FROM example WHERE name=?", name)
	 if err != nil {
		 return err
	 }

	 defer rows.
}