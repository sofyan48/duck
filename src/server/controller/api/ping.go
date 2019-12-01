package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/duck/src/server/helper"
)

// PingController Object
type PingController struct{}

// Ping Function Controller
func (p PingController) Ping(c *gin.Context) {
	helper.ResponseMsg(c, 200, "Pong Pong!")
	return
}
