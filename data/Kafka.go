package data

import (
	"bufio"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func GetConfigSet() *kafka.ConfigMap {
	var bootstrap string
	path := os.Getenv("KAFKA_BOOTSTRAP_SERVER")
	if path == "" {
		fmt.Println("The KAFKA_BOOTSTRAP_SERVER variable is not set.")
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening The KAFKA_BOOTSTRAP_SERVER secret file:", err)
		return nil
	}

	Reader := bufio.NewReader(file)
	fmt.Fscan(Reader, &bootstrap)
	defer file.Close()

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
