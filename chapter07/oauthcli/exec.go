package oauthcli

import (
	"os"
	"io"
	"fmt"
	"net/http"
)

//GetUsers usa um cliente oauth2 inicializado
// para obter informações sobre um usuário
func GetUsers(client *http.Client) error  {
	
	url := fmt.Sprintf("http://api.github.com/user")

	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	fmt.Println("Código de estatus da ", url, ":", resp.StatusCode)
	io.Copy(os.Stdout, resp.Body)
	return nil
}