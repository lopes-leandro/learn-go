package oauthstore

import (
	"context"
	"net/http"
	"golang.org/x/oauth2"
)

//Config envolve o oauth2.Config padrão e
// adiciona nosso armazenamento
type Config struct{
	*oauth2.Config
	Storage
}

//Exchange armazena um token após a recuperação
func (c *Config) Exchange(ctx context.Context, code string) (*oauth2.Token, error){
	
	token, err := c.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	if err := c.Storage.SetToken(token); err != nil {
		return nil, err
	}
	return token, nil
}

//TokenSource pode receber um token que
// é armazenado ou, quando um novo é
// recuperado, é armazenado
func (c *Config) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource  {
	return StorageTokenSource(ctx, c, t)
}

//Client está anexado ao nosso TokenSource
func (c *Config) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return oauth2.NewClient(ctx, c.TokenSource(ctx, t))
}