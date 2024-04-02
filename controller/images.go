package controller

import (
	"Tmage/controller/status"
	"Tmage/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HomeHandler(c *gin.Context) {
	ResponseSuccess(c)
}

func UploadHandler(c *gin.Context) {
	// Multipart form
	//接收文件
	form, err := c.MultipartForm()
	if err != nil {
		zap.L().Error("upload files with invalid parameters", zap.Error(err))
		ResponseErrorWithMsg(c, status.StatusInvalidParams)
		return
	}

	//接收标签
	var tags []string
	if err = c.BindJSON(&tags); err != nil {
		zap.L().Error("upload tags with invalid parameters", zap.Error(err))
		ResponseErrorWithMsg(c, status.StatusInvalidParams)
		return
	}

	//获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("log in and upload", zap.Error(err))
		ResponseError(c, err)
		return
	}

	//业务处理，上传图片
	files := form.File["uploads"]
	err = logic.Upload(files, tags, userID)
	if err != nil {
		zap.L().Error("logic.upload failed", zap.Error(err))
		ResponseError(c, err)
		return
	}

	ResponseSuccess(c)
}

func DeleteHandler(c *gin.Context) {
	//接收文件id
	var imageIds []string
	if err := c.BindJSON(&imageIds); err != nil {
		zap.L().Error("delete with invalid parameters", zap.Error(err))
		ResponseErrorWithMsg(c, status.StatusInvalidParams)
		return
	}

	//获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("log in and delete", zap.Error(err))
		ResponseError(c, err)
		return
	}

	//业务处理，删除图片
	err = logic.Delete(imageIds, userID)
	if err != nil {
		zap.L().Error("logic.delete failed", zap.Error(err))
		ResponseError(c, err)
		return
	}

	ResponseSuccess(c)
}
