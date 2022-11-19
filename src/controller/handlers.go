package controller

import (
	"fmt"
	"net/http"
)

func GetCheckout(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "404 Only Post is available", http.StatusNotFound)
		return
	}

	error := r.ParseForm()

	if error != nil {
		http.Error(w, "parse error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Welcome to my website!")
}
