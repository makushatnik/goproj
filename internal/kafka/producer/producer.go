package producer

import (
	"context"
	"goproj/internal/config"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const messagesCount = 20

func Produce(cfg *config.Config) {
	conn, err := kafka.DialLeader(context.Background(), "tcp",
		cfg.KafkaBroker, cfg.KafkaTopic, cfg.KafkaPartition)
	if err != nil {
		log.Fatalln("failed to dial leader:", err)
	}

	for i := 0; i < messagesCount; i++ {
		_, err := conn.WriteMessages(
			kafka.Message{
				Key:   []byte("push"),
				Value: []byte("It works!"),
			},
		)
		if err != nil {
			log.Fatalln("failed to write a message:", err)
		}

		time.Sleep(1 * time.Second)
	}

	if err := conn.Close(); err != nil {
		log.Fatalln("failed to close writer:", err)
	}
}
