package main

import (
	"log"
	"net/http"
)

func main() {

	//Basic Auth Accounts
	accounts := make([]string, 1)
	accounts = append(accounts, "hai:pwd")

	//Home
	http.HandleFunc("/", Log(Home))

	//Handle request on question endpoint
	http.HandleFunc("/v1/question", Group(Log(QuestionAPI), EnsureAuthorize(accounts)))

	log.Fatal(http.ListenAndServe(":1313", nil))
}
