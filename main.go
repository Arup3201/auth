package main

import (
	"auths/methods/basic"
	"net/http"
)

func main() {
	http.HandleFunc("/api", basic.HandleBasicAuth)

	http.ListenAndServe(":8000", nil)
}
