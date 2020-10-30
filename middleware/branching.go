package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func Method(method string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			//Detect rq method and redirect the request
			if r.Method != method {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			//execute next handling
			f(w, r)
		}
	}
}

func EnsureAuthorize(accounts []string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

			if len(auth) != 2 || auth[0] != "Basic" {
				http.Error(w, "authorization failed", http.StatusUnauthorized)
				return
			}

			payload, _ := base64.StdEncoding.DecodeString(auth[1])
			pair := strings.SplitN(string(payload), ":", 2)

			if len(pair) != 2 || !validate(pair[0], pair[1], accounts) {
				http.Error(w, "authorization failed", http.StatusUnauthorized)
				return
			}

			//execute next handling
			f(w, r)
		}
	}
}

func validate(username, password string, accounts []string) bool {
	for _, account := range accounts {
		if account == (username + ":" + password) {
			return true
		}
	}

	return false
}
