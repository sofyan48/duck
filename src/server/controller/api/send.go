package controller

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/duck/src/libs"
	"github.com/sofyan48/duck/src/libs/scheme"
	helper "github.com/sofyan48/duck/src/server/helper"
)

// SendController Object
type SendController struct{}

// SendTask Function Controller
func (send SendController) SendTask(c *gin.Context) {
	data := scheme.SendTask{}
	c.ShouldBindJSON(&data)

	srv, err := libs.InitServer(os.Getenv("REST_ENV_PATH"))
	libs.Check(err)
	// send task
	result, err := libs.SendTask(srv, data)
	libs.Check(err)
	helper.ResponseData(c, 200, result)
	return
}
