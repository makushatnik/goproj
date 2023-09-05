package consumer

import (
	"context"
	"fmt"
	"goproj/internal/config"
	"log"

	"github.com/segmentio/kafka-go"
)

func Consume(cfg *config.Config) {
	conn, err := kafka.DialLeader(context.Background(), "tcp",
		cfg.KafkaBroker, cfg.KafkaTopic, cfg.KafkaPartition)
	if err != nil {
		log.Fatalln("failed to dial leader:", err)
	}

	for {
		m, err := conn.ReadMessage(1e6)
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := conn.Close(); err != nil {
		log.Fatalln("failed to close writer:", err)
	}
}
