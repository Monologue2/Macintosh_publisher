package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/Monologue2/Early-Birds.git/api"
	"github.com/Monologue2/Early-Birds.git/data"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	producer *kafka.Producer
	err      error
)

// Please set up SECRET_APIKEY_PATH env variable before running
func main() {
	producer, err = kafka.NewProducer(data.GetConfigSet())
	if err != nil {
		// panic(err)
		fmt.Printf("Failed to create producer : %s", err)
		return
	}

	asos := api.New(
		api.WithTm(data.GetCurrentTime()),
		api.WithStn(0),
		api.WithHelp(0),
		api.WithAuthKey(),
	)

	// []byte stream
	responseBody, _ := api.AsosGetRequest(asos)
	// fmt.Printf("%s", responseBody)

	// PreProcessing Response
	reader := bytes.NewReader(responseBody)
	scanner := bufio.NewScanner(reader)
	startReading := false

	for scanner.Scan() {
		line := scanner.Text()

		// Check for start and end of the relevant data
		if strings.Contains(line, "#START7777") {
			startReading = true
		}

		if strings.Contains(line, "#7777END") {
			break
		}

		if strings.Contains(line, "{") {
			fmt.Printf("failed to authenticate api key : %v\n", string(responseBody))
			break
		}

		if startReading {
			fields := strings.Fields(line)
			if len(fields) < 16 {
				continue // skip lines that do not have enough data
			}

			if strings.Contains(line, "#") {
				continue
			}

			jsonData, station_id := data.ByteToJsonform(fields)
			data.ProduceJsonToBroker(producer, jsonData, station_id)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	}

	// Wait for message deliveries before shutting down
	producer.Flush(15 * 1000)
}
