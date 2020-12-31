package oauthstore

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

//Storage é nossa interface de armazenamento genérica
type Storage interface {
	GetToken() (*oauth2.Token, error)
	SetToken(*oauth2.Token) error
}

//GetToken recupera um token github oauth2
func GetToken(ctx context.Context, conf Config) (*oauth2.Token, error) {

	token, err := conf.Storage.GetToken()
	if err == nil && token.Valid() {
		return token, err
	}

	url := conf.AuthCodeURL("state")
	fmt.Printf("Digite o seguinte url em seu navegador e siga as instruções na tela: %v\n", url)
	fmt.Println("Cole o código retornado no URL de redirecionamento e pressione Enter: ")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}

	return conf.Exchange(ctx, code)
}
