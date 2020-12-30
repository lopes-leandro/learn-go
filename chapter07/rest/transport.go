package rest

import "net/http"

//APITransport faz um SetBasicAuth
// para cada requisição
type APITransport struct {
	*http.Transport
	username, password string
}

//RoundTrip faz a autenticação básica
// antes de adiar para o transporte padrão
func (t *APITransport) RoundTrip(req *http.Request) (*http.Response, error) {

	req.SetBasicAuth(t.username, t.password)

	return t.Transport.RoundTrip(req)
}
