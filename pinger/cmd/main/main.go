package main

import (
	"fmt"
	"time"

	"pinger/internal/ping"
)

func main() {
	for {
		time.Sleep(5 * time.Second)

		results := ping.GetPingResults()
		fmt.Println(results)
	}
}
