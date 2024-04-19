# Macintosh Publisher
```
.
├── api
│   └── Api.go
├── data
│   ├── Kafka.go
│   └── Processing.go
├── go.mod
├── go.sum
├── main.go
├── README.md
└── test
    └── data.txt
```

This project is designed to automatically fetch and process weather data from all observatories in Korea using the Korea Meteorological Administration's REST API. <br>
It's built with Go and utilizes a Kafka Cluster to handle the streaming data efficiently.<br>

## Key Features

- **Data Collection**: Collects comprehensive weather data including temperature, humidity, wind direction, wind speed, and barometric pressure from all Korean observatories.
- **Data Parsing**: Parses the received data and converts it into JSON format suitable for transmission.
- **Data Streaming**: Streams the formatted data to a designated Kafka topic specific to each observatory ID (e.g., `weather.{observatoryID}`).
- **Automation**: Uses GitHub Actions to automatically update the weather data in Kafka every hour.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- You must have Go installed on your machine (visit [Go's official site](https://golang.org/dl/) for download & installation instructions).
- Access to a Kafka Cluster where you can create and manage topics.

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Monologue2/Macintosh_publisher.git
   cd Macintosh_publisher
   ```

2. **Set up the required secrets**
- You need to provide your Korea Meteorological Administration REST API key and your Kafka Cluster's listener address as environment variables


3. **Run the application**
   ```bash
   go mod tidy
   go run .
   ```

### Usage
The system is set to run automatically via GitHub Actions, which triggers every hour to fetch and update the weather data.
