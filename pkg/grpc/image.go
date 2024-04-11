package grpc

import (
	"Tmage/controller/status"
	"Tmage/models"
	"github.com/listenGrey/TmagegRpcPKG/images"
	"google.golang.org/grpc"
)

// 定义gRpc客户端服务器的类型码

type ImageClient int64

const (
	Download  ImageClient = 2003
	Search    ImageClient = 2004
	OpenShare ImageClient = 2005
)

func ImageClientServer(funcCode ImageClient) (client interface{}) {
	conn, err := grpc.Dial("localhost:8964", grpc.WithInsecure()) //server IP
	if err != nil {
		return status.StatusConnGrpcServerERR
	}
	switch funcCode {
	case Download:
		client = images.NewDownloadServiceClient(conn)
	case Search:
		client = images.NewSearchServiceClient(conn)
	case OpenShare:
		client = images.NewShareServiceClient(conn)
	default:
		client = nil
	}
	return client
}

func ParseImages(receive *images.Images) (images []models.UploadImage, info status.Code) {
	code := receive.GetInfo()
	if code == status.StatusSuccess.Code() {
		info = status.StatusSuccess
	} else if code == status.StatusBusy.Code() {
		info = status.StatusBusy
	} else if code == status.StatusConnDBERR.Code() {
		info = status.StatusConnDBERR
	}

	data := receive.GetImages()
	for _, v := range data {
		var image models.UploadImage
		image.UserID = v.GetUserID()
		image.Tags = v.GetTags()
		image.UploadTime = v.GetUploadTime()
		image.ImageName = v.GetImageName()
		image.Size = v.GetSize()
		image.Type = v.GetType()
		image.Content = v.GetContent()

		images = append(images, image)
	}

	return
}
