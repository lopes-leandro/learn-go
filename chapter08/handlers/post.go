package handlers

import (
	"encoding/json"
	"net/http"
)

//GreetingResponse é a resposta JSON
// que GreetingHandler retorna
type GreetingResponse struct {
	Payload struct {
		Greeting string `json:"greeting"`
		Name string `json:"name"`
		Error string `json:"error"`
	} `json:"payload"`
	Successful bool `json:"successful"`
}

//GreetingHandler retorna um GreetingResponse
// que contém erros ou uma carga útil útil
func GreetingHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var gr GreetingResponse
	if err := req.ParseForm(); err != nil {
		gr.Payload.Error = "bad request"
		if payload, err := json.Marshal(gr); err == nil {
			w.Write(payload)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	name := req.FormValue("name")
	greeting := req.FormValue("greeting")

	w.WriteHeader(http.StatusOK)

	gr.Successful = true
	gr.Payload.Name = name
	gr.Payload.Greeting = greeting
	if payload, err := json.Marshal(gr); err == nil {
		w.Write(payload)
	}

}
