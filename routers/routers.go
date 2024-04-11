package routers

import (
	"Tmage/controller"
	"Tmage/logger"
	"Tmage/middleWares"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/pprof"
	"time"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置发布模式
	}
	r := gin.New()

	// 设置中间件
	r.Use(logger.GinLogger(),
		logger.GinRecovery(true),
		middleWares.RateLimitMiddleWare(2*time.Second, 40),
	)

	r.LoadHTMLFiles() // 加载HTML
	r.Static()        // 加载静态文件

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置单张图片最大为32MB r.MaxMultipartMemory 默认为32MB
	v1 := r.Group("/api/v1")
	v1.POST("/register", controller.RegisterHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)
	v1.GET("/share/:token", controller.OpenShareHandler) // 打开分享图片,无需登录

	// 中间件
	v1.Use(middleWares.JWTAuthMiddleWare()) //JWT认证
	{
		v1.GET("/images", controller.HomeHandler) //首页

		//文件业务
		v1.POST("/upload", controller.UploadHandler)     // 图片上传
		v1.POST("/delete", controller.DeleteHandler)     // 删除图片
		v1.POST("/edit", controller.EditHandler)         // 编辑图片信息
		v1.POST("/share", controller.ShareHandler)       // 分享图片
		v1.POST("/search", controller.SearchHandler)     // 图片搜索
		v1.POST("/download", controller.DownloadHandler) // 图片下载

		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}

	pprof.Register(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
