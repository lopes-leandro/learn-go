package handlers

import (
	"fmt"
	"net/http"
)

//HelloHandler recebe um parâmetro GET "name"
// e responde com Olá<name>! em texto simples
func HelloHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	name := req.URL.Query().Get("name")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Olá %s!\t\n", name)))
}
