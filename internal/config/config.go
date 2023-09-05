package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDSN          string
	Host           string
	HTTPPort       int
	GRPCPort       int
	KafkaBroker    string
	KafkaTopic     string
	KafkaPartition int
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatal("Http Port can't be turned into an integer")
		return nil, err
	}

	grpcPort, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatal("Grpc Port can't be turned into an integer")
		return nil, err
	}

	conf := &Config{
		DBDSN:          os.Getenv("DBDSN"),
		Host:           "localhost",
		HTTPPort:       httpPort,
		GRPCPort:       grpcPort,
		KafkaBroker:    "kafka",
		KafkaTopic:     "test",
		KafkaPartition: 1,
	}
	return conf, nil
}
