package util

import (
	"Tmage/models"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"

	"context"
)

func OperationLog(info models.OperationInfo) {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "operationLog",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true, //实际情况是false
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", info.UserID)) // key=id+time
	key = append(key, '+')
	key = append(key, []byte(info.Time)...)
	value, err := json.Marshal(info) // value
	if err != nil {
		return
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
		return
	}
}
