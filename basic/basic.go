package basic

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"strings"
)

func getCredentialsFromAuthHeader(authHeader string) (string, string, error) {
	credentials := strings.Replace(authHeader, "Basic", "", 1)
	credentials = strings.Trim(credentials, " ")

	decoded, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		return "", "", errors.New("error while decoding")
	}

	result := strings.Split(string(decoded), ":")

	return result[0], result[1], nil
}

func httpAuthError(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", "Basic realm='BasicAuth'")
	w.WriteHeader(http.StatusUnauthorized)
}

func BasicAuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			httpAuthError(w)
			return
		}

		username, password, err := getCredentialsFromAuthHeader(authorizationHeader)
		if err != nil {
			log.Fatal(err)
			httpAuthError(w)
			return
		}

		if (username != "arup") || (password != "1234") {
			log.Fatal("incorrect credentials")
			httpAuthError(w)
			return
		}

		f(w, r)
	}
}
