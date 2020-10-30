package main

import (
	"github.com/haithngn/go-crud/controller"
	"github.com/haithngn/go-crud/db"
	"log"
	"net/http"
)

func main() {
	database, err := db.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	storage := db.Storage{Database: database}

	handler := Router{HomeCtrlr: controller.HomeController{}, QuestionCtrlr: controller.QuestionController{storage}}
	handler.Setup()

	http.ListenAndServe(":1313", nil)
}
