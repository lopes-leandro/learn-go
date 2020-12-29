package pools

import (
	"database/sql"
	"fmt"
	"os"
		//importamos bibliotecas com suporte para database/sql
		_ "github.com/go-sql-driver/mysql"
)

//Setup configura o banco de dados junto com o
// número de conjuntos de conexões e muito mais
func Setup() (*sql.DB, error)  {
	db, err := sql.Open("mysql", 
	fmt.Sprintf("%s:%s@tcp(localhost:8083)/gocoockbook?parseTime=true", 
	os.Getenv("$MYSQLUSERNAME"), 
	os.Getenv("$MYSQLPASSWORD")))
	if err != nil {
		return nil, err
	}

	//sempre haverá apenas 24 conexões abertas
	db.SetMaxOpenConns(24)

	//MaxIdleConns nunca pode ser menor que abertura máxima
	//SetMaxOpenConns, caso contrário, o padrão será esse valor
	db.SetMaxIdleConns(24)

	return db, nil
}