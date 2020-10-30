package server

import (
	"github.com/gin-gonic/gin"
	"github.com/haithngn/go-crud/router"
	"log"
)

type Server struct {
	Router *gin.Engine
}

func (server *Server) Serve(routes []router.Router) {
	for _, route := range routes {
		route.Start(server.Router)
	}
	server.start()
}

func (server *Server) start() {
	server.Router.Use(gin.Logger())
	err := server.Router.Run(":1313")
	if err != nil {
		log.Fatal(err.Error())
	}
}
