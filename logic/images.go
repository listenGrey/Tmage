package logic

import (
	"Tmage/controller/status"
	"Tmage/models"
	"Tmage/util"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

const fileSize = 50 * 1024 * 1024 //单张图片最大50MB

func Upload(files []*multipart.FileHeader, tags []string, userID int64) (info string, code status.Code) {
	total := len(files)
	success, failed := 0, 0
	var uploadImages []models.UploadImage

	// 转换tag类型
	var formTags models.FormTags
	err := json.Unmarshal([]byte(tags[0]), &formTags)
	if err != nil {
		return "", status.StatusBusy
	}
	uploadTags := formTags.Tags

	for _, file := range files {
		// 校验文件类型是否为图片
		ext := strings.ToLower(path.Ext(file.Filename))

		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			continue
		}

		// 校验文件大小是否超过限制
		if file.Size > fileSize {
			continue
		}

		// 获取当前时间
		now := time.Now()

		// 获取图片内容
		fileContent, err := file.Open()
		if err != nil {
			continue
		}
		defer fileContent.Close()

		fileBytes, err := ioutil.ReadAll(fileContent)
		if err != nil {
			continue
		}

		// 将符合条件的图片上传
		uploadImage := &models.UploadImage{
			UserID:     userID,
			Tags:       uploadTags,
			UploadTime: now.Format("2006-01-02 15:04:05"),
			ImageName:  file.Filename,
			Size:       file.Size,
			Type:       ext[1:],
			Content:    fileBytes,
		}
		uploadImages = append(uploadImages, *uploadImage)
	}
	res := util.Upload(uploadImages)
	if res != status.StatusSuccess {
		return "", res
	}

	// 发送日志信息
	now := time.Now()
	operationInfo := &models.OperationInfo{
		UserID:    userID,
		Operation: "upload",
		Time:      now.Format("2006-01-02 15:04:05"),
		Status:    res.Msg(),
		Data:      uploadImages,
	}
	util.OperationLog(*operationInfo)

	success = len(uploadImages)
	failed = total - success
	return fmt.Sprintf("总共上传%d张图片，%d张成功，%d张失败", total, success, failed), status.StatusSuccess
}

func Delete(imageIds []string, userID int64) (code status.Code) {
	now := time.Now()
	code = util.Delete(imageIds, userID, now.Format("2006-01-02 15:04:05"))

	// 发送日志信息
	operationInfo := &models.OperationInfo{
		UserID:    userID,
		Operation: "delete",
		Time:      now.Format("2006-01-02 15:04:05"),
		Status:    code.Msg(),
		Data:      imageIds,
	}
	util.OperationLog(*operationInfo)

	return
}

func Edit(image models.ModifyImage, userID int64) (code status.Code) {
	now := time.Now()
	code = util.Edit(image, userID, now.Format("2006-01-02 15:04:05"))

	// 发送日志信息
	operationInfo := &models.OperationInfo{
		UserID:    userID,
		Operation: "edit",
		Time:      now.Format("2006-01-02 15:04:05"),
		Status:    code.Msg(),
		Data:      image,
	}
	util.OperationLog(*operationInfo)

	return
}

func Download(imageIds []string, userID int64) (images []models.UploadImage, code status.Code) {
	images, code = util.Download(imageIds, userID)

	// 发送日志信息
	now := time.Now()
	operationInfo := &models.OperationInfo{
		UserID:    userID,
		Operation: "download",
		Time:      now.Format("2006-01-02 15:04:05"),
		Status:    code.Msg(),
		Data:      images,
	}
	util.OperationLog(*operationInfo)

	return
}

func Share(imageIds []string, userID int64) (url string, code status.Code) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", status.StatusBusy
	}

	encodedToken := base64.URLEncoding.EncodeToString(tokenBytes)

	// 分享链接到期时间：14天
	now := time.Now()
	expirationTime := now.Add(14 * 24 * time.Hour)

	shareImage := &models.ShareImage{
		UserID:         userID,
		ImagesIds:      imageIds,
		Token:          encodedToken,
		ShareTime:      now.Format("2006-01-02 15:04:05"),
		ExpirationTime: expirationTime.Format("2006-01-02 15:04:05"),
	}

	code = util.Share(*shareImage)
	if code != status.StatusSuccess {
		return "", code
	}

	// 构建 URL
	url = fmt.Sprintf("https://tmage.org/share/%s", encodedToken)

	// 发送日志信息
	operationInfo := &models.OperationInfo{
		UserID:    userID,
		Operation: "share",
		Time:      now.Format("2006-01-02 15:04:05"),
		Status:    code.Msg(),
		Data:      imageIds,
	}
	util.OperationLog(*operationInfo)

	return url, status.StatusSuccess
}

func Search(tags []string, userID int64) (images []models.UploadImage, code status.Code) {
	images, code = util.Search(tags, userID)
	if code != status.StatusSuccess {
		return nil, code
	}

	// 发送日志信息
	now := time.Now()
	operationInfo := &models.OperationInfo{
		UserID:    userID,
		Operation: "search",
		Time:      now.Format("2006-01-02 15:04:05"),
		Status:    code.Msg(),
		Data:      images,
	}
	util.OperationLog(*operationInfo)

	return
}

func OpenShare(token string) (images []models.UploadImage, code status.Code) {
	images, code = util.OpenShare(token)
	if code != status.StatusSuccess {
		return nil, code
	}

	// 发送日志信息
	now := time.Now()
	operationInfo := &models.OperationInfo{
		UserID:    0,
		Operation: "openShare",
		Time:      now.Format("2006-01-02 15:04:05"),
		Status:    code.Msg(),
		Data:      token,
	}
	util.OperationLog(*operationInfo)

	return
}
