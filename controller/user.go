package controller

import (
	"Tmage/controller/status"
	"Tmage/logic"
	"Tmage/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {
	//获取请求参数，校验数据
	var client *models.RegisterFrom
	if err := c.ShouldBindJSON(&client); err != nil {
		zap.L().Error("Register with invalid parameter", zap.Error(err))
		//判断 err 是否为 validator 类型
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, status.StatusInvalidParams) // 请求参数错误
			return
		}
		ResponseErrorWithMsg(c, status.StatusInvalidParams)
		// 翻译错误
		return
	}
	// 用户注册
	if err := logic.Register(client); err != nil {
		zap.L().Error("logic.Register failed", zap.Error(err))
		ResponseError(c, err)
		return
	}

	ResponseSuccess(c)
}

func LoginHandler(c *gin.Context) {
	//获取请求参数，校验参数
	var user *models.LoginForm
	if err := c.ShouldBindJSON(&user); err != nil {
		zap.L().Error("log in with invalid parameters", zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, err)
			return
		}
		ResponseErrorWithMsg(c, status.StatusInvalidParams)
		return
	}

	//业务处理：登录
	_, err := logic.Login(user)
	if err != nil {
		zap.L().Error("logic.login failed", zap.String("user", user.Email), zap.Error(err))
		ResponseError(c, err)
		return
	}

	//返回响应
	ResponseSuccess(c)
}

/*
func RefreshTokenHandler(c *gin.Context) {

}
func JWTAuthMiddleWare(c *gin.Context) {

}*/
