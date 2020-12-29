package redis

import(
	"fmt"
rds "gopkg.in/redis.v5"
)

//Sort executa operações de classificação do redis
func Sort() error  {
	
	conn, err := Setup()
	if err != nil {
		return err
	}

	listkey := "list"

	if err := conn.LPush(listkey, 1).Err(); err != nil {
		return err
	}

	//isto irá limpar a chave da lista se
	// qualquer um dos comandos subsequentes falhar
	defer conn.Del(listkey)

	if err := conn.LPush(listkey, 3).Err(); err != nil {
		return err
	}

	if err := conn.LPush(listkey, 2).Err(); err != nil {
		return err
	}

	res, err := conn.Sort(listkey, rds.Sort{Order: "ASC"}).Result()
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}