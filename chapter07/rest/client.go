package rest

import "net/http"

//APIClient é nosso cliente customizado
type APIClient struct {
	*http.Client
}

//NewAPIClient o construtor que inicializa o cliente
// com nosso transporte personalizado
func NewAPIClient(username, password string) *APIClient {

	t := http.Transport{}
	return &APIClient{
		Client: &http.Client{
			Transport: &APITransport{
				Transport: &t,
				username:  username,
				password:  password,
			},
		},
	}
}

//GetGoogle é uma chamada de API - nós
// abstraímos os aspectos REST
func (c *APIClient) GetGoogle() (int, error) {

	resp, err := c.Get("http://www.google.com")
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}
