package controller

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sofyan48/duck/src/server/controller/api"

	"github.com/sofyan48/duck/src/server/middlewares"
)

// RoutesController | Create Route Controller Rest API
func RoutesController(r *gin.Engine) {
	// Get Controller
	ping := new(controller.PingController)
	send := new(controller.SendController)
	get := new(controller.GetTaskController)
	// Create Routes With Auth Declare Here
	//r.Use(middlewares.AuthToken())
	api := r.Group("api")
	{
		api.GET("/ping", middlewares.AuthACL(), ping.Ping)
		api.POST("/send/task", middlewares.AuthACL(), send.SendTask)
		api.POST("/get/task", middlewares.AuthACL(), get.GetTask)
	}
}
