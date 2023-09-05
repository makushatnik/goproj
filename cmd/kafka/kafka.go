package main

import (
	"goproj/internal/config"
	"goproj/internal/kafka/consumer"
	"goproj/internal/kafka/producer"
	"goproj/utils"
	"time"
)

func main() {
	cfg, err := config.New()
	utils.LogError(err)

	go producer.Produce(cfg)
	go consumer.Consume(cfg)

	time.Sleep(60 * time.Second)
}
