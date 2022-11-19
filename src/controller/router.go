package controller

import "net/http"

func Router() {

	http.HandleFunc("/", GetCheckout)

	http.ListenAndServe(":80", nil)
}
