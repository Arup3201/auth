package basic

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getCredentialsFromAuthHeader(authHeader string) (string, string, error) {
	credentials := strings.Replace(authHeader, "Basic", "", 1)
	fmt.Println("Credentials: ", credentials)

	decoded, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		log.Fatal("Decode error: ", err)
		return "", "", errors.New("Decoding failed")
	}

	result := strings.Split(string(decoded), ":")

	return result[0], result[1], nil
}

func BasicAuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		username, password, err := getCredentialsFromAuthHeader(authorizationHeader)
		if err != nil {
			w.Header().Add("WWW-Authenticate", "Basic realm='BasicAuth'")
			http.Error(w, http.StatusText(http.StatusNetworkAuthenticationRequired), http.StatusNetworkAuthenticationRequired)
			return
		}

		if (username != "arup") || (password != "1234") {
			log.Fatal("Incorrect credentials")
			w.Header().Add("WWW-Authenticate", "Basic realm='BasicAuth'")
			http.Error(w, http.StatusText(http.StatusNetworkAuthenticationRequired), http.StatusNetworkAuthenticationRequired)
			return
		}

		f(w, r)
	}
}
