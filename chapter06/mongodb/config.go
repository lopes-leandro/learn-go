package mongodb

import(
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

//Setup inicializa um cliente mongo
func Setup(ctx context.Context, address string) (*mongo.Client, error)  {
	
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	//cancel será chamado quando sair da configuração
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	return client, nil
}