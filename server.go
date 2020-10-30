package main

import "net/http"

func main() {
	storage := Storage{[]Question{}}
	handler := RequestHandler{storage}

	//Home
	http.HandleFunc("/", handler.Home)

	//Handle request on question endpoint
	http.HandleFunc("/question", handler.Question)

	http.ListenAndServe(":1313", nil)
}
