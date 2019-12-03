package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/duck/src/libs/scheme"
	helper "github.com/sofyan48/duck/src/server/helper"
)

// SendController Object
type SendController struct{}

// SendTask Function Controller
func (send SendController) SendTask(c *gin.Context) {
	data := scheme.SendTask{}
	c.ShouldBindJSON(&data)

	helper.ResponseData(c, http.StatusOK, data)
	return
}
