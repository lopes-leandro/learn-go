package client

import (
	"crypto/tls"
	"net/http"
)

//Setup configura nosso cliente e redefine
// o DefaultClient global
func Setup(isSecure, nop bool) *http.Client {

	c := http.DefaultClient

	//Às vezes, para teste, queremos
	// desligar a verificação SSL
	if !isSecure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	}

	if nop {
		c.Transport = &NopTransport{}
	}

	http.DefaultClient = c
	return c
}

//NopTransport é um transporte sem operação
type NopTransport struct {
}

//RoundTrip implementa interface RoundTripper
func (n *NopTransport) RoundTrip(*http.Request) (*http.Response, error) {
	//observe que esta é uma resposta não inicializada
	// se você estiver vendo cabeçalhos e assim por diante.
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}
