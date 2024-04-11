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
		ResponseError(c, status.StatusBusy, nil)
		return
	}
	ResponseSuccess(c, info)
}

func DeleteHandler(c *gin.Context) {
	//接收文件id
	var imageIds models.FormImageIds
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
	uploadImageIds := imageIds.ImageIds
	code := logic.Delete(uploadImageIds, userID)
	if code != status.StatusSuccess {
		zap.L().Error("logic.delete failed", zap.Error(errors.New(code.Msg())))
		ResponseError(c, status.StatusBusy, nil)
		return
	}

	ResponseSuccess(c, "删除成功")
}

func EditHandler(c *gin.Context) {
	//接收要修改的图片和tags
	var image models.ModifyImage
	if err := c.BindJSON(&image); err != nil {
		zap.L().Error("edit with invalid parameters", zap.Error(err))
		ResponseError(c, status.StatusInvalidParams, nil)
		return
	}

	//获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("log in and edit", zap.Error(err))
		ResponseError(c, status.StatusNotLogin, nil)
		return
	}

	//业务处理，编辑图片标签
	code := logic.Edit(image, userID)
	if code != status.StatusSuccess {
		zap.L().Error("logic.edit failed", zap.Error(errors.New(code.Msg())))
		ResponseError(c, status.StatusBusy, nil)
		return
	}

	ResponseSuccess(c, "修改成功")
}

func DownloadHandler(c *gin.Context) {
	// 接收文件id
	var imageIds models.FormImageIds
	if err := c.BindJSON(&imageIds); err != nil {
		zap.L().Error("download with invalid parameters", zap.Error(err))
		ResponseError(c, status.StatusInvalidParams, nil)
		return
	}

	// 获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("log in and download", zap.Error(err))
		ResponseError(c, status.StatusNotLogin, nil)
		return
	}

	// 获取图片文件
	uploadImageIds := imageIds.ImageIds
	images, code := logic.Download(uploadImageIds, userID)
	if code != status.StatusSuccess {
		zap.L().Error("logic.download failed", zap.Error(errors.New(code.Msg())))
		ResponseError(c, status.StatusBusy, nil)
		return
	}

	// 业务处理，下载图片
	for _, image := range images {
		// 设置响应头
		c.Header("Content-Type", image.Type)
		c.Header("Content-Disposition", "attachment; imagename="+image.ImageName)
		c.Writer.Write(image.Content)
	}

	ResponseSuccess(c, nil)
}

func ShareHandler(c *gin.Context) {
	// 接收文件id
	var imageIds models.FormImageIds
	if err := c.BindJSON(&imageIds); err != nil {
		zap.L().Error("share with invalid parameters", zap.Error(err))
		ResponseError(c, status.StatusInvalidParams, nil)
		return
	}

	// 获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("log in and share", zap.Error(err))
		ResponseError(c, status.StatusNotLogin, nil)
		return
	}

	// 业务处理，分享图片
	uploadImageIds := imageIds.ImageIds
	shareURL, code := logic.Share(uploadImageIds, userID)
	if code != status.StatusSuccess {
		zap.L().Error("logic.share failed", zap.Error(errors.New(code.Msg())))
		ResponseError(c, status.StatusBusy, nil)
		return
	}

	ResponseSuccess(c, shareURL)
}

func SearchHandler(c *gin.Context) {
	// 接收tags
	var tags models.FormTags
	if err := c.BindJSON(&tags); err != nil {
		zap.L().Error("search with invalid parameters", zap.Error(err))
		ResponseError(c, status.StatusInvalidParams, nil)
		return
	}

	// 获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("log in and search", zap.Error(err))
		ResponseError(c, status.StatusNotLogin, nil)
		return
	}

	// 业务处理，搜索图片
	uploadTags := tags.Tags
	images, code := logic.Search(uploadTags, userID)
	if code != status.StatusSuccess {
		zap.L().Error("logic.search failed", zap.Error(errors.New(code.Msg())))
		ResponseError(c, status.StatusBusy, nil)
		return
	}

	ResponseSuccess(c, images)
}

func OpenShareHandler(c *gin.Context) {
	// 接收token
	token := c.Param("token")

	// 业务处理，返回被分享的图片
	sharedImages, code := logic.OpenShare(token)
	if code != status.StatusSuccess {
		zap.L().Error("logic.openShare failed", zap.Error(errors.New(code.Msg())))
		if code == status.StatusShareInvalid {
			ResponseError(c, status.StatusShareInvalid, nil)
		}
		ResponseError(c, status.StatusBusy, nil)
		return
	}

	ResponseSuccess(c, sharedImages)
}
