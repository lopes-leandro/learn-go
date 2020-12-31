package oauthcli

import (
	"context"
	"fmt"
	"os"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

//Setup retorna um oauth2Config configurado
// para falar com o github, você precisa de variáveis
// de ambiente definidas para seu id e segredo
func Setup() *oauth2.Config  {
	return &oauth2.Config{
		ClientID: os.Getenv("GITHUB_CLIENT"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		Scopes: []string{"repo", "user"},
		Endpoint: github.Endpoint,
	}
}

//GetToken recupera um token github oauth2
func GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error)  {
	
	url := conf.AuthCodeURL("state")
	
	fmt.Printf("Digite a seguinte url em seu navegador e siga as instruções na tela: %v\n", url)
	fmt.Println("Cole o código retornado na URL de redirecionamento e pressione Enter:")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}

	return conf.Exchange(ctx, code)
}