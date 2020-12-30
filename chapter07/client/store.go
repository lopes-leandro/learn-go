package client

import (
	"fmt"
	"net/http"
)

//Controller incorpora um http.Client
// e o usa internamente
type Controller struct{
	*http.Client
}

//DoOps com um objeto controlador
func (c *Controller) DoOps() error  {
	
	res, err := c.Client.Get("http://www.google.com")
	if err != nil {
		return err
	}

	fmt.Println("resultados do client.DoOps: ", res.StatusCode)

	return nil
}