package storage

 import "context"

 //Item representa um item em
 // uma loja
 type Item struct {
	 Name string
	 Price int64
 }

 //Storage é a nossa interface de armazenamento
//Vamos implementá-lo com o
// armazenamento Mongo
type Storage interface {
	GetByName(context.Context, string) (*Item, error)
	Put(context.Context, *Item) error
}