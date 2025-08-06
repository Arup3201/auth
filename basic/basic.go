package basic

import (
	"fmt"
	"net/http"
)

func HandleBasicAuth(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%v\n", req.Method)
}
