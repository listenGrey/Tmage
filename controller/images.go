package controller

import (
	"Tmage/controller/status"
	"Tmage/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

func PreviewHandler(c *gin.Context) {
	tag := c.Param("tag")

	//使用tag获取图片文件名列表 filenames

	filenames, err := logic.GetFilesByTags(tag)
	if err != nil {
		zap.L().Error("logic.GetFilesByTags failed", zap.Error(err))
		ResponseError(c, err)
		return
	}
	if len(filenames) == 0 {
		ResponseErrorWithMsg(c, status.StatusImagesNotFound)
		return
	}

	// 构造预览图片的HTML代码
	var previewURls []string
	for _, filename := range filenames {
		previewURls = append(previewURls, "/preview/"+filename)
	}

	c.JSON(http.StatusOK, gin.H{
		"tag":    tag,
		"images": previewURls,
	})
}

func GetImagesHandler(c *gin.Context) {
	filename := c.Param("filename")
	filepath := "http://server2/uploads/" + filename // 假设服务器2上的存储系统的地址为"http://server2/uploads/"

	// 获取图片文件内容
	resp, err := http.Get(filepath)
	//if err != nil {
	//c.JSON(http.StatusOK)
	//return
	//}
	//defer resp.Body.Close()

	// 读取图片文件内容
	data, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//c.JSON()
	//return
	//}

	// 将图片文件内容返回给客户端
	c.Data(http.StatusOK, "image/jpg", data)
}
