package middleWares

import (
	"Tmage/controller"
	"Tmage/controller/status"
	"Tmage/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// 基于 JWT 的认证中间件
func JWTAuthMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Token的方式 1.放在请求头 2.放在请求体 3.放在URI
		// Token放在Header的Authorization中，使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, status.StatusInvalidToken, nil)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, status.StatusInvalidToken, nil)
			c.Abort()
			return
		}
		// 解析Token
		user, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, status.StatusInvalidToken, nil)
			c.Abort()
			return
		}
		// 将user信息保存到context中
		c.Set(controller.ContextUserIDKey, user.UserID)
		c.Next()
		// 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}
