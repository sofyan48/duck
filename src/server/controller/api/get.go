package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/duck/src/libs"
	"github.com/sofyan48/duck/src/libs/scheme"
	helper "github.com/sofyan48/duck/src/server/helper"
)

// GetTaskController Object
type GetTaskController struct{}

// GetTask Function Controller
func (ctx GetTaskController) GetTask(c *gin.Context) {
	data := scheme.GetTask{}
	c.ShouldBindJSON(&data)
	result, err := libs.GetResult(data.Duck.Action.Worker, data.Duck.Action.Trigger, data.Duck.UUID)
	if err != nil {
		helper.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.ResponseData(c, http.StatusOK, result)
	return
}
