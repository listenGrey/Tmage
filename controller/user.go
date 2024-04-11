package controller

import (
	"Tmage/controller/status"
	"Tmage/logic"
	"Tmage/models"
	"Tmage/pkg/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func RegisterHandler(c *gin.Context) {
	//获取请求参数，校验数据
	var client *models.RegisterFrom
	if err := c.ShouldBindJSON(&client); err != nil {
		zap.L().Error("Register with invalid parameter", zap.Error(err))
		//判断 err 是否为 validator 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, status.StatusInvalidParams, nil) // 请求参数错误
			return
		}
		ResponseError(c, status.StatusInvalidParams, errs.Translate(trans))
		// 翻译错误
		return
	}
	// 用户注册
	if code := logic.Register(client); code != status.StatusSuccess {
		zap.L().Error("logic.Register failed", zap.Error(errors.New(code.Msg())))
		if code == status.StatusUserExist {
			ResponseError(c, status.StatusUserExist, nil)
		}
		ResponseError(c, status.StatusBusy, nil)
		return
	}

	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	//获取请求参数，校验参数
	var user *models.LoginForm
	if err := c.ShouldBindJSON(&user); err != nil {
		zap.L().Error("log in with invalid parameters", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, status.StatusInvalidParams, nil)
			return
		}
		ResponseError(c, status.StatusInvalidParams, errs.Translate(trans))
		return
	}

	//业务处理：登录
	curUser, code := logic.Login(user)
	if code != status.StatusSuccess {
		zap.L().Error("logic.login failed", zap.String("user", user.Email), zap.Error(errors.New(code.Msg())))
		if code == status.StatusUserNotExist || code == status.StatusInvalidPwd {
			ResponseError(c, code, nil)
		}
		ResponseError(c, status.StatusBusy, nil)
		return
	}

	//返回响应
	ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", curUser.UserID),
		"user_name":     curUser.UserName,
		"access_token":  curUser.AccessToken,
		"refresh_token": curUser.RefreshToken,
	})
}

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseError(c, status.StatusInvalidToken, "请求头缺少Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseError(c, status.StatusInvalidToken, "Token格式错误")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	zap.L().Error("jwt.RefreshToken failed", zap.Error(err))
	if err != nil {
		ResponseError(c, status.StatusInvalidToken, "刷新Token错误")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
