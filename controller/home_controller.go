package controller

import (
	"fmt"
	"net/http"
)

type HomeController struct {
}

func (controller *HomeController) Home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "text/html; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprintf(writer, "<h2 align=\"center\">Welcome to Golang Viet Nam Workshop 2 - FAQ app</h2>")

	return
}
