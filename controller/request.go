package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "userID"
)

var (
	ErrorUserNotLogin = errors.New("当前用户未登录")
)

// getCurrentUserID 获取当前登录用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	_userID, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = _userID.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
