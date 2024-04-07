package controller

import (
	"Tmage/controller/status"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseContent struct {
	Code    status.Code `json:"code"`
	Err     error       `json:"error"`
	Content interface{} `json:"content,omitempty"` // content为空时不显示
}

func ResponseError(c *gin.Context, code status.Code, content interface{}) {
	re := &ResponseContent{
		Code:    code,
		Err:     errors.New(code.Msg()),
		Content: content,
	}
	c.JSON(http.StatusOK, re)
}

func ResponseSuccess(c *gin.Context, content interface{}) {
	re := &ResponseContent{
		Code:    status.StatusSuccess,
		Err:     nil,
		Content: content,
	}
	c.JSON(http.StatusOK, re)
}
