package redis

import(
	"fmt"
	"time"
	rds "gopkg.in/redis.v5"
)

//Exec executa algumas operações de redis
func Exec() error  {
	conn, err := Setup()
	if err != nil {
		return err
	}

	c1 := "value"

	//value é uma interface, podemos armazenar qualquer
	// que seja o último argumento da expiração do redis
	conn.Set("key", c1, 5*time.Second)

	var result string
	if err := conn.Get("key").Scan(&result) ; err != nil {
		switch err {
			//isso significa que a
			// chave não foi encontrada
		case rds.Nil:
			return nil
		default:
			return err
		}
	}

	fmt.Println("resultado = ", result)

	return nil
}