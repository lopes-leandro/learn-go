package controllers

import (
	"encoding/json"
	"net/http"
)

//GetValue é um closure que envolve um HandlerFunc,
// se UseDefault for true, o valor sempre será "default",
// caso contrário, será o que estiver armazenado no armazenamento.
func (c *Controller) GetValue(UseDefault bool) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		value := "default"
		if !UseDefault {
			value = c.storage.Get()
		}

		w.WriteHeader(http.StatusOK)
		p := Payload{Value: value}
		if payload, err := json.Marshal(p); err == nil {
			w.Write(payload)
		}
	}
}
