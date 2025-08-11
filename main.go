package main

import (
	"auths/methods/basic"
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func main() {
	http.HandleFunc("/api", basic.BasicAuthMiddleware(index))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
