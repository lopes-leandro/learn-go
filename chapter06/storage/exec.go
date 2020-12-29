package storage

import (
	"context"
	"fmt"
	"os"
)

//Exec inicializa o armazenamento e, em seguida,
// executa as operações usando a interface de armazenamento
func Exec() error  {
	
	ctx := context.Background()
	
	m, err := NewMongoStorage(ctx, 
		fmt.Sprintf("mongodb://%s:%s@localhost:27017", 
		os.Getenv("MONGODBUSERNAME"), 
		os.Getenv("MONGODBPASSWORD")), 
		"gocoockbook", 
		"items")
	if err != nil {
		return err
	}

	if err := PerformOperations(m); err != nil {
		return err
	}

	if err := m.Client.Database(m.DB).Collection(m.Collection).Drop(ctx); err != nil {
		return err
	}

	return nil
}

//PerformOperations cria um
// item vela e o obtém
func PerformOperations(s Storage) error  {
	
	ctx := context.Background()

	i := Item{Name: "candles", Price: 100}
	if err := s.Put(ctx, &i); err != nil {
		return err
	}

	candles, err := s.GetByName(ctx, "candles")
	if err != nil {
		return err
	}

	fmt.Printf("Resultado: %#v\n", candles)

	return nil
}