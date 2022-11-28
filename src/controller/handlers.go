package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"web-app/src/model"
	"web-app/src/service"
)

func Checkout(responseWriter http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(responseWriter, "I couldan't even simulate this error, congrats!! msg me please", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPost:
		postCheckout(responseWriter, body)

	default:
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func postCheckout(responseWriter http.ResponseWriter, body []byte) {

	var checkout model.Checkout

	err := json.Unmarshal(body, &checkout)

	if err != nil {
		http.Error(responseWriter, "Ajusta esse payload ai meu compatriota, é um map de string to any, não é possível que vc ta errando,"+
			"aceita literalmente qualquer coisa como valor.", http.StatusBadRequest)
		return
	}

	service.SaveCheckout(checkout)

}
