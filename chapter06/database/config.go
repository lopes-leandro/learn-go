package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	//importamos bibliotecas com suporte para database/sql
	_ "github.com/go-sql-driver/mysql"
)

//Example contém os resultados de nossas consultas
type Example struct {
	Name    string
	Created *time.Time
}

//Setup configura e retorna nossos
// pools de conexão de banco de dados
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(localhost:8083)/gocookbook?parseTime=true",
			os.Getenv("MYSQLUSERNAME"), os.Getenv("MYSQLPASSWORD")))
	if err != nil {
		return nil, err
	}
	return db, nil
}
