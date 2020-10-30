package router

import (
	"github.com/gin-gonic/gin"
	"github.com/haithngn/go-crud/controller"
	"github.com/haithngn/go-crud/middleware"
	"gorm.io/gorm"
)

type Router interface {
	Start(server *gin.Engine)
}

func GetRoutes(database *gorm.DB) []Router {
	return []Router{&QuestionRouter{Controller: &controller.QuestionContoller{DB: database}, SecureMiddleware: &middleware.SecureMiddleware{}}}
}
