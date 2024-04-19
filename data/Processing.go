package data

import (
	"encoding/json"
	"strconv"
	"time"
)

type WeatherData struct {
	DateTime    string  `json:"datetime"`
	StationID   int     `json:"station_id"`
	WindDir     int     `json:"wind_dir"`
	WindSpeed   float64 `json:"wind_speed"`
	Pressure    float64 `json:"pressure"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func GetCurrentTime() string {
	// Get the current time
	now := time.Now()

	// Truncate to the nearest hour to zero out the minutes and seconds
	truncated := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())

	// Format the time to YYYYMMDDHHMM
	// Note that MM (for month) and mm (for minutes) use the same MM in this layout because Go uses the same representation for month and minute based on the position in the format string
	timeString := truncated.Format("200601021504")

	return timeString
}

func ByteToJsonform(fields []string) ([]byte, int) {
	datetime := fields[0]
	stationID, _ := strconv.Atoi(fields[1])
	windDir, _ := strconv.Atoi(fields[2])
	windSpeed, _ := strconv.ParseFloat(fields[3], 64)
	pressure, _ := strconv.ParseFloat(fields[7], 64)
	temperature, _ := strconv.ParseFloat(fields[11], 64)
	humidity, _ := strconv.ParseFloat(fields[13], 64)

	data := WeatherData{
		DateTime:    datetime,
		StationID:   stationID,
		WindDir:     windDir,
		WindSpeed:   windSpeed,
		Pressure:    pressure,
		Temperature: temperature,
		Humidity:    humidity,
	}

	jsonData, _ := json.Marshal(data)
	return jsonData, stationID
}
