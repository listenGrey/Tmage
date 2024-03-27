package util

import (
	"Tmage/controller/status"
	"Tmage/models"
	"fmt"
	"github.com/segmentio/kafka-go"
	"strconv"
	"time"

	"context"
)

func Upload(uploadImages []*models.UploadImage) status.Code {
	ctx := context.Background()
	topic := strconv.FormatInt(uploadImages[0].UserID, 10)
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true, //实际情况是false
	}

	defer writer.Close()

	// 发送消息
	for i := 0; i < 3; i++ {
		err := writer.WriteMessages(
			ctx,
			kafka.Message{
				Key:   []byte("1"),
				Value: []byte("l"),
			},
			kafka.Message{
				Key:   []byte("2"),
				Value: []byte("i"),
			},
			kafka.Message{
				Key:   []byte("3"),
				Value: []byte("s"),
			},
		) // 原子操作，全部成功或全部不成功

		if err != nil {
			if err == kafka.LeaderNotAvailable { //第一次写，topic不存在
				time.Sleep(500 * time.Microsecond)
				continue
			} else {
				fmt.Printf("批量写入失败：%v\n", err)
			}
		} else {
			break
		}
	}

}
