package pools

import(
	"context"
	"time"
)

//ExecWithTimeout atingirá o tempo limite
// tentando obter a hora atual
func ExecWithTimeout() error  {
	db, err := Setup()
	if err != nil {
		return err
	}

	ctx := context.Background()

	//queremos interromper o tempo imediatamente
	ctx, cancel := context.WithDeadline(ctx, time.Now())

	//chame cancel depois de completarmos
	defer cancel()

	//nossa transação é sensível ao contexto
	_, err = db.BeginTx(ctx, nil)

	return err
}