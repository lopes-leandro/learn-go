package decorator

import "fmt"

//Exec cria um cliente, chama google.com
// e imprime a resposta
func Exec() error  {
	
	c := Setup()

	resp, err := c.Get("https://www.google.com")
	if err != nil {
		return err
	}

	fmt.Println("Código da resposta: ", resp.StatusCode)

	return nil
}