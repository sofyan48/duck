package controller

import (
	"net/http"

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
	result, err := libs.Send(data.Duck.Action.Worker, data.Duck.Action.Trigger, data)
	if err != nil {
		helper.ResponseMsg(c, http.StatusInternalServerError, err)
	}
	helper.ResponseData(c, http.StatusOK, result)
	return
}
