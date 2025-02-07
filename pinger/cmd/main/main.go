package main

import (
	"time"

	"pinger/internal/ping"
	"pinger/api"
)

func main() {
	for {
		time.Sleep(5 * time.Second)

		results := ping.GetPingResults()
		api.SendResultsToAPI(results)
	}
}
