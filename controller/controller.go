package controller

import (
	"Tmage/logic"
	"Tmage/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	//获取请求参数，校验数据
	var client *models.RegisterFrom
	if err := c.ShouldBindJSON(&client); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		/*
				// translate errors
				c.JSON(http.StatusBadRequest, gin.H{
					"msg: ": removeTopStruct(errs.Translate(trans)),
				})
				return
			}
		*/
	}
	// client register
	if err := logic.Register(client); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

/*
func LoginHandler(c *gin.Context) {
	//获取请求参数，校验参数
	var u models.LoginForm
	if err := c.ShouldBindJSON(&u); err != nil {

	}
	//业务处理：登录
	//返回响应
}

func RefreshTokenHandler(c *gin.Context) {

}
func JWTAuthMiddleWare(c *gin.Context) {

}*/
