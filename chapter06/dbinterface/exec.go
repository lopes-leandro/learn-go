package dbinterface

//Exec substitui o Exec da receita anterior
func Exec(db DB) error  {

    //erro não detectado na limpeza,
		// mas sempre queremos limpar
		defer db.Exec("DROP TABLE example")

		if err := Create(db); err != nil {
			return err
		}

		if err := Query(db); err != nil {
			return err
		}

		return nil
}