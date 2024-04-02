package logic

import (
	"Tmage/controller/status"
	"Tmage/models"
	"Tmage/util"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

const fileSize = 50 * 1024 * 1024 //单张图片最大50MB

func Upload(files []*multipart.FileHeader, tags []string, userID int64) (err error) {
	var uploadImages []models.UploadImage
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
			Tags:       tags,
			UploadTime: now.Format("2006-01-02 15:04:05"),
			ImageName:  file.Filename,
			Size:       file.Size,
			Content:    fileBytes,
		}
		uploadImages = append(uploadImages, *uploadImage)
	}
	res := util.Upload(uploadImages)
	if res != status.StatusSuccess {
		return errors.New(res.Msg())
	}
	return nil
}

func Delete(imageIds []string, userID int64) (err error) {
	res := util.Delete(imageIds, userID)
	if res != status.StatusSuccess {
		return errors.New(res.Msg())
	}
	return nil
}
