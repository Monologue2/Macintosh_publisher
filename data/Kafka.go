package data

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func GetConfigSet() *kafka.ConfigMap {
	bootstrap := os.Getenv("KAFKA_BOOTSTRAP_SERVER")

	config := &kafka.ConfigMap{
		"bootstrap.servers": bootstrap,
	}
	return config
}

func ProduceJsonToBroker(p *kafka.Producer, m []byte, station_id int) {
	topic := fmt.Sprintf(
		"weather.%d",
		station_id,
	)

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          m,
	}, nil)
}
