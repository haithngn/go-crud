package router

import (
	"github.com/gin-gonic/gin"
	"github.com/haithngn/go-crud/controller"
	"github.com/haithngn/go-crud/middleware"
)

type QuestionRouter struct {
	Controller       *controller.QuestionContoller
	SecureMiddleware *middleware.SecureMiddleware
}

func (router *QuestionRouter) Start(server *gin.Engine) {
	version1 := server.Group("v1", gin.BasicAuth(gin.Accounts{
		"hai":   "pwd",
		"admin": "pwd",
	}))
	{
		version1.GET("/question/:id", router.Controller.GetQuestion)
		version1.POST("/question", router.SecureMiddleware.EnsureOTP(), router.Controller.CreatingNewQuestion)
		version1.PUT("/question/:id", router.SecureMiddleware.EnsureOTP(), router.Controller.UpdateQuestion)
	}
	version2 := server.Group("v2", gin.BasicAuth(gin.Accounts{
		"hai":   "pwd",
		"admin": "pwd",
	}))
	{
		version2.DELETE("/question/:id", router.SecureMiddleware.EnsureOTP(), router.Controller.RemoveQuestion)
	}
}
