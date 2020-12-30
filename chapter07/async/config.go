package async

import "net/http"

//Client armazena um cliente e tem dois
// canais para agregar respostas e erros
type Client struct{
	*http.Client
	Resp chan *http.Response
	Err chan error
}

//NewClient cria um novo cliente e
// define seus canais apropriados
func NewClient(client *http.Client, bufferSize int) *Client  {
	respch := make(chan *http.Response, bufferSize)
	errch := make(chan error, bufferSize)
	return &Client{
		Client: client,
		Resp: respch,
		Err: errch,
	}
}

//AsyncGet executa um Get e retorna o
// resp/error para o canal apropriado
func (c *Client) AsyncGet(url string)  {
	
	resp, err := c.Get(url)
	if err != nil {
		c.Err <- err
		return
	}
	c.Resp <- resp
}