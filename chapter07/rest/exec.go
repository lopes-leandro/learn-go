package rest

import "fmt"

//Exec cria um cliente API e usa seu método
// GetGoogle, a seguir imprime o resultado
func Exec() error {

	c := NewAPIClient("username", "password")

	StatusCode, err := c.GetGoogle()
	if err != nil {
		return err
	}

	fmt.Println("Resultado do GetGoogle: ", StatusCode)

	return nil
}
