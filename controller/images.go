package controller

import (
	"Tmage/controller/status"
	"Tmage/logic"
	"Tmage/models"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HomeHandler(c *gin.Context) {
	ResponseSuccess(c, "index")
}

func UploadHandler(c *gin.Context) {
	// Multipart form
	//接收文件和标签
	form, err := c.MultipartForm()
	if err != nil {
		zap.L().Error("upload files with invalid parameters", zap.Error(err))
		ResponseError(c, status.StatusInvalidParams, nil)
		return
	}

	//获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("log in and upload", zap.Error(err))
		ResponseError(c, status.StatusNotLogin, nil)
		return
	}

	//业务处理，上传图片
	files := form.File["uploads"]
	tags := form.Value["tags"]
	info, code := logic.Upload(files, tags, userID)
	if code != status.StatusSuccess {
		zap.L().Error("logic.upload failed", zap.Error(errors.New(code.Msg())))
		ResponseError(c, code, nil)
		return
	}
	ResponseSuccess(c, info)
}

func DeleteHandler(c *gin.Context) {
	//接收文件id
	var imageIds []string
	if err := c.BindJSON(&imageIds); err != nil {
		zap.L().Error("delete with invalid parameters", zap.Error(err))
		ResponseError(c, status.StatusInvalidParams, nil)
		return
	}

	//获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("log in and delete", zap.Error(err))
		ResponseError(c, status.StatusNotLogin, nil)
		return
	}

	//业务处理，删除图片
	code := logic.Delete(imageIds, userID)
	if code != status.StatusSuccess {
		zap.L().Error("logic.delete failed", zap.Error(errors.New(code.Msg())))
		ResponseError(c, code, nil)
		return
	}

	ResponseSuccess(c, "删除成功")
}

func EditHandler(c *gin.Context) {

}

func DownloadHandler(c *gin.Context) {

}

func ShareHandler(c *gin.Context) {

}

func SearchHandler(c *gin.Context) {
	//接收tags
	var tags models.FormTags
	if err := c.BindJSON(&tags); err != nil {
		zap.L().Error("delete with invalid parameters", zap.Error(err))
		ResponseError(c, status.StatusInvalidParams, nil)
		return
	}

	//获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("log in and delete", zap.Error(err))
		ResponseError(c, status.StatusNotLogin, nil)
		return
	}

	//业务处理，搜索图片
	uploadTags := tags.Tags
	images, code := logic.Search(uploadTags, userID)
	if code != status.StatusSuccess {
		zap.L().Error("logic.delete failed", zap.Error(errors.New(code.Msg())))
		ResponseError(c, code, nil)
		return
	}

	ResponseSuccess(c, images)
}
