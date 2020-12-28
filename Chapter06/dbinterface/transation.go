package dbinterface

import "database/sql"

//DB é uma interface que é satisfeita
// por um sql.DB ou um sql.Transaction
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

//Transaction pode fazer qualquer coisa que
// uma consulta pode fazer, mais Commit, Rollback ou Stmt
type Transaction interface{
	DB
	Commit() error
	Rolback() error
}