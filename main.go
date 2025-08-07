package main

import (
	"auths/methods/basic"
	"auths/methods/logging"
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/api", logging.LoggingMiddleware(basic.BasicAuthMiddleware(index)))

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Server failed to start")
	}
}
