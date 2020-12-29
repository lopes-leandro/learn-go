package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

//GetByName consulta mongodb para um item
// com o nome correto
func (m *MongoStorage) GetByName(ctx context.Context, name string) (*Item, error) {
	
	c := m.Client.Database(m.DB).Collection(m.Collection)
	
	var i Item
	if err := c.FindOne(ctx, bson.M{"name": name}).Decode(&i); err != nil {
		return nil, err
	}

	return &i, nil
}

//Put adiciona um item à nossa instância mongo
func (m *MongoStorage) Put(ctx context.Context, i *Item) error  {
	
	c := m.Client.Database(m.DB).Collection(m.Collection)

	_, err := c.InsertOne(ctx, i)

	return err
}