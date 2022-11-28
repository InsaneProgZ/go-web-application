package controller

import "net/http"

func Router() {

	http.HandleFunc("/", Checkout)

	http.ListenAndServe(":80", nil)
}
