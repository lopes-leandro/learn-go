package dbinterface

import (
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	//importamos bibliotecas com suporte para database/sql
	_ "github.com/go-sql-driver/mysql"
)

func TestExec(t *testing.T) {

	tests := []struct {
		name      string
		createErr bool
		queryErr  bool
		wantErr   bool
	}{
		{"base-case", false, false, false},
		{"create error", true, false, true},
		{"query error", false, true, true},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("um erro '%s' não era esperado ao abrir uma conexão de banco de dados")
			}

			defer db.Close()

			mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(0, 0))
			c := mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(0, 0))

			if tt.createErr {
				c.WillReturnResult(errors.New("falhou"))
			}

			if !tt.createErr {
				m := mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]stirng{"name", "created"}))
				if tt.queryErr {
					m.WillReturnError(errors.New("falhou"))
				}
			}

			if err := Exec(db); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("havia expectativas não atendidas: %s", err)
			}

		})
	}
}
