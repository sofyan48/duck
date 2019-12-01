package controller

import (
	"net/http"
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
	if err != nil {
		helper.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	result, err := libs.SendTask(srv, data)
	if err != nil {
		helper.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.ResponseData(c, http.StatusOK, result)
	return
}
