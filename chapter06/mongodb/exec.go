package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

//State é nosso modelo de dados
type State struct {
	Name       string `bson:"name"`
	Population int    `bson:"pop"`
}

//Exec cria e consulta um exemplo
func Exec(address string) error {
	ctx := context.Background()
	db, err := Setup(ctx, address)
	if err != nil {
		return err
	}

	coll := db.Database("gocookbook").Collection("example")

	vals := []interface{}{&State{"Washington", 7062000}, &State{"Oregon", 3970000}}

	//podemos inserir várias linhas de uma vez
	if _, err := coll.InsertMany(ctx, vals); err != nil {
		return err
	}

	var s State
	if err := coll.FindOne(ctx, bson.M{"name": "Washington"}).Decode(&s); err != nil {
		return err
	}

	var ss State
	if err := coll.FindOne(ctx, bson.M{"name": "Oregon"}).Decode(&ss); err != nil {
		return err
	}

	if err := coll.Drop(ctx); err != nil {
		return err
	}

	fmt.Printf("State: %#v\n", s)

	fmt.Printf("State: %#v\n", ss)

	return nil
}
