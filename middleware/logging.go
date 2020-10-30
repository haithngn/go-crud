package middleware

import (
	"fmt"
	"net/http"
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
