package util

import (
	"Tmage/controller/status"
	"Tmage/models"
	"Tmage/pkg/grpc"
	"encoding/json"
	"fmt"
	images2 "github.com/listenGrey/TmagegRpcPKG/images"
	"github.com/segmentio/kafka-go"
	"time"

	"context"
)

func Upload(uploadImages []models.UploadImage) status.Code {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "upload",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", uploadImages[0].UserID)) // key=id+time
	key = append(key, '+')
	key = append(key, []byte(uploadImages[0].UploadTime)...)
	value, err := json.Marshal(uploadImages) // value
	if err != nil {
		return status.StatusBusy
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)

	if err != nil {
		return status.StatusKafkaSendERR
	}

	return status.StatusSuccess
}

func Delete(imageIds []string, userID int64, deleteTime string) status.Code {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "delete",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true, //实际情况是false
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", userID)) // key=id+time
	key = append(key, '+')
	key = append(key, []byte(deleteTime)...)
	value, err := json.Marshal(imageIds) // value
	if err != nil {
		return status.StatusBusy
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return status.StatusKafkaSendERR
	}

	return status.StatusSuccess
}

func Edit(image models.ModifyImage, userID int64, editTime string) (code status.Code) {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "edit",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true, //实际情况是false
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", userID)) // key=id+time
	key = append(key, '+')
	key = append(key, []byte(editTime)...)
	value, err := json.Marshal(image) // value
	if err != nil {
		return status.StatusBusy
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return status.StatusKafkaSendERR
	}

	return status.StatusSuccess
}

func Share(shareImage models.ShareImage) status.Code {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "share",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true, //实际情况是false
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", shareImage.UserID)) // key=id+time
	key = append(key, '+')
	key = append(key, []byte(shareImage.ShareTime)...)
	value, err := json.Marshal(shareImage) // value
	if err != nil {
		return status.StatusBusy
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return status.StatusKafkaSendERR
	}

	return status.StatusSuccess
}

func Download(imageIds []string, userID int64) (images []models.UploadImage, code status.Code) {
	client := grpc.ImageClientServer(grpc.Download)
	if client == status.StatusConnGrpcServerERR {
		return nil, status.StatusConnGrpcServerERR
	}
	send := &images2.DownloadInfo{ImageIds: imageIds, UserID: userID}
	res, err := client.(images2.DownloadServiceClient).Download(context.Background(), send)
	if err != nil {
		return nil, status.StatusRecvGrpcSerInfoERR
	}
	images, code = grpc.ParseImages(res)

	return
}

func Search(tags []string, userID int64) (images []models.UploadImage, code status.Code) {
	client := grpc.ImageClientServer(grpc.Search)
	if client == status.StatusConnGrpcServerERR {
		return nil, status.StatusConnGrpcServerERR
	}
	send := &images2.SearchInfo{UserID: userID, Tags: tags}
	res, err := client.(images2.SearchServiceClient).Search(context.Background(), send)
	if err != nil {
		return nil, status.StatusRecvGrpcSerInfoERR
	}
	images, code = grpc.ParseImages(res)

	return
}

func OpenShare(token string) (images []models.UploadImage, code status.Code) {
	client := grpc.ImageClientServer(grpc.OpenShare)
	if client == status.StatusConnGrpcServerERR {
		return nil, status.StatusConnGrpcServerERR
	}
	send := &images2.ShareInfo{Token: token}
	res, err := client.(images2.ShareServiceClient).Share(context.Background(), send)
	if err != nil {
		return nil, status.StatusRecvGrpcSerInfoERR
	}
	images, code = grpc.ParseImages(res)

	return
}
