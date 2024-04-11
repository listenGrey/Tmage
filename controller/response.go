package controller

import (
	"Tmage/controller/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseContent struct {
	Code    status.Code `json:"code"`
	Err     string      `json:"error"`
	Content interface{} `json:"content,omitempty"` // content为空时不显示
}

func ResponseError(c *gin.Context, code status.Code, content interface{}) {
	re := &ResponseContent{
		Code:    code,
		Err:     code.Msg(),
		Content: content,
	}
	c.JSON(http.StatusOK, re)
}

func ResponseSuccess(c *gin.Context, content interface{}) {
	re := &ResponseContent{
		Code:    status.StatusSuccess,
		Err:     status.StatusSuccess.Msg(),
		Content: content,
	}
	c.JSON(http.StatusOK, re)
}
