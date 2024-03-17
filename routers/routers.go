package routers

import (
	"Tmage/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.POST("/register", controller.RegisterHandler)
	v1.POST("/login", controller.LoginHandler)
	//v1.GET("/refresh_token", controller.RefreshTokenHandler)
	/*
		v1.Use(controller.JWTAuthMiddleWare())
		{
			v1.GET()
			v1.GET()
			v1.POST()
		}

	*/
	/*
		v1.Use(middleWares.JWTAuthMiddleWare())
		{

		}
	*/
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
