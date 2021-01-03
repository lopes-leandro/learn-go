package proxy

import (
	"log"
	"net/http"
)

//Proxy mantém nosso cliente configurado
// e BaseURL para proxy para
type Proxy struct{
	Client *http.Client
	BaseURL string
}

//ServeHTTP significa que o proxy implementa
// a interface Handler. Ele manipula a requisição,
// encaminha-a para BaseURL e retorna a resposta
func (p *Proxy) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	
	if err := p.ProcessRequest(req); err != nil {
		log.Printf("um erro ocorreu durante o processo de requisição: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := p.Client.Do(req)
	if err != nil {
		log.Printf("um erro ocorreu durante a operação do cliente: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()
	CopyRespose(w, res)
}