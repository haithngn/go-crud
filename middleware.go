package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Log(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Requested At: ", time.Now().String())
		fmt.Println("======== REQUEST LINE ========")
		fmt.Println("Method: ", request.Method)
		fmt.Println("Host: ", request.Host)
		fmt.Println("Path: ", request.URL.Path)
		fmt.Println("Query: ", request.URL.RawQuery)
		fmt.Println("Fragment: ", request.URL.RawFragment)
		fmt.Println("======== REQUEST HEADER ========")
		fmt.Println("User-Agent: ", request.Header.Get("User-Agent"))
		fmt.Println("Content-type: ", request.Header.Get("Content-type"))
		fmt.Println("Connection: ", request.Header.Get("Connection"))
		fmt.Println("Accept: ", request.Header.Get("Accept"))
		fmt.Println("Accept-Encoding: ", request.Header.Get("Accept-Encoding"))

		handlerFunc(writer, request)
	}
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Group(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

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
