package client

import (
	"fmt"
	"net/http"
)

//DoOps pega um cliente e, em seguida,
// busca google.com
func DoOps(c *http.Client) error  {
	
	res, err := c.Get("http://www.tuiaseguros.com.br")
	if err != nil {
		return err
	}

	fmt.Println("resultados do DoOps: ", res.StatusCode)

	return nil
}

//DefaultGetGolang usa o cliente padrão para
// obter golang.org
func DefaultGetGolang() error  {
	res, err := http.Get("https://www.golang.org")
	if err != nil {
		return err
	}

	fmt.Println("resultados do DefaultGetGolang: ", res.StatusCode)

	return nil
}