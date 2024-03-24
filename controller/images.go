package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func UploadHandler(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["uploads"]
	for _, file := range files {
		// 上传文件至指定目录
		d, x := c.Get(ContextUserIDKey)
		c.SaveUploadedFile(file, "./"+file.Filename)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
