package main

import (
	"time"
	"os"
	"strconv"
	"log"

	"pinger/internal/ping"
	"pinger/api"
)

func main() {
	intervalStr := os.Getenv("TIME_INTERVAL")
	interval, err := strconv.Atoi(intervalStr)
	if err != nil || interval <= 0 {
		interval = 10
		log.Printf("Некорректный интервал, используем значение по умолчанию: %d секунд", interval)
	}
	
	for {
		time.Sleep(time.Duration(interval) * time.Second)

		results := ping.GetPingResults()
		api.SendResultsToAPI(results)
	}
}
