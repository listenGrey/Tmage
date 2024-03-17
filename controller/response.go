package controller

import (
	"Tmage/controller/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"msg": err.Error(),
	})
}

func ResponseErrorWithMsg(c *gin.Context, code status.Code) {
	c.JSON(http.StatusOK, gin.H{
		"msg": code.Msg(),
	})
}

func ResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": status.StatusSuccess,
	})
}
