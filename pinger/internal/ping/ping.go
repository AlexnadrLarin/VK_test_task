package ping

import (
	"context"
	"log"
	"net"
	"time"
	"fmt"
	"sync"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"pinger/models"
)

func GetPingResults() models.PingResults {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Ошибка подключения к Docker: %v", err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), containertypes.ListOptions{All: true})
	if err != nil {
		log.Fatalf("Ошибка получения списка контейнеров: %v", err)
	}

	var wg sync.WaitGroup
	resultsChannel := make(chan []models.PingResult)

	for _, container := range containers {
		wg.Add(1)
		go func(containerID string) {
			defer wg.Done()
			containerResults := inspectAndPingContainer(cli, containerID)
			resultsChannel <- containerResults
		}(container.ID)
	}

	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	var results models.PingResults
	for containerResults := range resultsChannel {
		results.Results = append(results.Results, containerResults...)
	}

	return results
}

func inspectAndPingContainer(cli *client.Client, containerID string) []models.PingResult {
	containerInfo, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		log.Printf("Ошибка инспекции контейнера %s: %v", containerID, err)
		return nil
	}

	var wg sync.WaitGroup
	resultsChannel := make(chan models.PingResult)

	for _, network := range containerInfo.NetworkSettings.Networks {
		ip := network.IPAddress
		if ip != "" {
			for port := range containerInfo.NetworkSettings.Ports {
				wg.Add(1)
				go func(ip, port string) {
					defer wg.Done()
					pingTime, date, success := pingContainer(ip, port)
					if success {
						resultsChannel <- models.PingResult{
							IP:       ip,
							PingTime: pingTime,
							Date:     date,
						}
					}
				}(ip, port.Port())
			}
		}
	}

	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	var results []models.PingResult
	for result := range resultsChannel {
		results = append(results, result)
	}

	return results
}

func pingContainer(ip string, port string) (string, string, bool) {
	pingTimeout := 5 * time.Second

	startTime := time.Now()
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", ip, port), pingTimeout)
	if err != nil {
		log.Printf("Ошибка пинга %s:%s %v", ip, port, err)
		return "", "", false
	}
	defer conn.Close()

	pingTime := time.Since(startTime).String()
	date := startTime.Format(time.RFC3339)

	return pingTime, date, true
}
