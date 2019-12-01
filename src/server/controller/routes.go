package controller

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sofyan48/duck/app/controller/api"

	"github.com/sofyan48/duck/app/middlewares"
)

// RoutesController | Create Route Controller Rest API
func RoutesController(r *gin.Engine) {
	// Get Controller
	login := new(controller.LoginController)
	ping := new(controller.PingController)

	// user := new(users.UsersController)

	// Create Routes No Auth Declare Here
	auth := r.Group("api")
	{
		auth.POST("/login", login.LoginUsers)
	}
	// Create Routes With Auth Declare Here
	r.Use(middlewares.AuthACL())
	//r.Use(middlewares.AuthToken())
	api := r.Group("api")
	{
		api.GET("/ping", ping.Ping)
	}
}
