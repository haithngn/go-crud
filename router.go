package main

import (
	"github.com/haithngn/go-crud/controller"
	"github.com/haithngn/go-crud/middleware"
	"net/http"
)

type Router struct {
	HomeCtrlr     controller.HomeController
	QuestionCtrlr controller.QuestionController
}

func (router *Router) Question(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		router.GetQuestion(writer, request)
		break
	case http.MethodPost:
		router.CreateQuestion(writer, request)
	default:
		router.abortRequest(writer)
		break
	}
}

func (router *Router) GetQuestion(writer http.ResponseWriter, request *http.Request) {
	//Re-direct the quest base on the http method
	router.QuestionCtrlr.GetQuestion(writer, request)
}

func (router *Router) CreateQuestion(writer http.ResponseWriter, request *http.Request) {
	router.QuestionCtrlr.CreateQuestion(writer, request)
}

//Handle Unexpected request
func (router *Router) abortRequest(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusBadRequest)
	return
}

func (router *Router) Setup() {

	accounts := make([]string, 1)
	accounts = append(accounts, "hai:pwd")

	//Home
	http.HandleFunc("/", middleware.Log(router.HomeCtrlr.Home))

	//Handle request on question endpoint
	http.HandleFunc("/v1/question", middleware.Group(middleware.Log(router.Question), middleware.EnsureAuthorize(accounts)))
}
