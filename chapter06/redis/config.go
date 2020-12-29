package redis

import (
	"os"
	rds "gopkg.in/redis.v5"
)

//Setup inicializa um cliente redis
func Setup() (*rds.Client , error)  {
	client := rds.NewClient(&rds.Options{
		Addr: "localhost:6379",
		Password: os.Getenv("REDISPASSWORD"),
		DB: 0,
	})

	_, err := client.Ping().Result()

	return client, err
}