package dbinterface

import (
	"errors"
	"testing"
	"github.com/golang/mock/gomock"
	//importamos bibliotecas com suporte para database/sql
	_ "github.com/go-sql-driver/mysql"

)

func TestCreate(t *testing.T)  {
	
	tests := []struct {
		name string
		ExecErr1 bool
		ExecErr2 bool
		wantErr bool
	}{
		{"base-case", false, false, false},
		{"1st exec error", true, false, false},
		{"2st exec error", false, true, true},
	}

	for _, tt := range tests {
		
		t.Run(tt.name, func (t *testing.T)  {
			
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			mockDB := NewMockDB(ctrl)
			m := mockDB
		})
	}
}