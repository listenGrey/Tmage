package util

import (
	"Tmage/controller/status"
	"Tmage/models"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"strconv"
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
		AllowAutoTopicCreation: true, //实际情况是false
	}

	defer writer.Close()

	// 发送消息
	for _, uploadImage := range uploadImages {
		sendK := []byte(strconv.Itoa(int(uploadImage.UserID))) // key=id+image name
		sendK = append(sendK, ' ')
		sendK = append(sendK, []byte(uploadImage.ImageName)...)
		sendV, _ := json.Marshal(uploadImage) // value
		err := writer.WriteMessages(
			ctx,
			kafka.Message{
				Key:   sendK,
				Value: sendV,
			},
		)

		if err != nil {
			return status.StatusKafkaSendERR
		}
	}
	return status.StatusSuccess
}

func Delete(imageIds []string, userID int64) status.Code {
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

	// 发送消息
	sendK := []byte(strconv.Itoa(int(userID))) // key=id
	var sendV []byte                           // value
	for k, imageId := range imageIds {
		image := []byte(imageId)
		if k != len(imageIds)-1 {
			image = append(image, ' ')
		}
		sendV = append(sendV, image...)
	}
	err := writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   sendK,
			Value: sendV,
		},
	)
	if err != nil {
		return status.StatusKafkaSendERR
	}

	// 接收回复的消息
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "delete",
		GroupID:     "producer",
		StartOffset: kafka.FirstOffset,
		MaxWait:     10 * time.Millisecond,
	})

	var info []byte
	for {
		msg, err := reader.ReadMessage(ctx)
		info = msg.Value
		if err != nil {
			return status.StatusKafkaReceiveERR
		}
	}
	res, err := strconv.Atoi(string(info))
	if err != nil {
		return status.StatusKafkaReceiveERR
	}
	writer.Close()
	reader.Close()

	if res != int(status.StatusSuccess) {
		return status.StatusKafkaReceiveERR
	}

	return status.StatusSuccess
}

func Edit(image models.ModifyImage, userID int64) (code status.Code) {
	return status.StatusSuccess

}

func Download(imageIds []string, userID int64) (images []models.UploadImage, code status.Code) {
	return nil, status.StatusSuccess

}

func Share(userID int64, imagesIds []string, encodedToken string, expirationTime time.Time) status.Code {
	return status.StatusSuccess

}

func Search(tags []string, userID int64) (images []models.UploadImage, code status.Code) {
	return
}

func OpenShare(token string) (images []models.UploadImage, code status.Code) {
	return
}
