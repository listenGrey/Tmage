package logic

import (
	"Tmage/controller/status"
	"Tmage/models"
	"Tmage/util"
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
		return "标签为空", status.StatusNoTag
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
			Content:    fileBytes,
		}
		uploadImages = append(uploadImages, *uploadImage)
	}
	res := util.Upload(uploadImages)
	if res != status.StatusSuccess {
		return "", res
	}
	success = len(uploadImages)
	failed = total - success
	return fmt.Sprintf("总共上传%d张图片，%d张成功，%d张失败", total, success, failed), status.StatusSuccess
}

func Delete(imageIds []string, userID int64) (code status.Code) {
	code = util.Delete(imageIds, userID)
	if code != status.StatusSuccess {
		return code
	}
	return
}

func Edit() (info string, code status.Code) {
	return
}

func Download() (info string, code status.Code) {
	return
}

func Share() (info string, code status.Code) {
	return
}

func Search(tags []string, userID int64) (images []models.UploadImage, code status.Code) {
	images, code = util.Search(tags, userID)
	if code != status.StatusSuccess {
		return nil, code
	}
	return
}
