package ping

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

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
	var results []models.PingResult

	var resultsLock sync.Mutex 

	for _, container := range containers {
		wg.Add(1)
		go func(containerID string) {
			defer wg.Done()
			containerResults := inspectAndPingContainer(cli, containerID)
			resultsLock.Lock()
			results = append(results, containerResults...)
			resultsLock.Unlock()
		}(container.ID)
	}

	wg.Wait()

	return models.PingResults{Results: results}
}

func inspectAndPingContainer(cli *client.Client, containerID string) []models.PingResult {
	containerInfo, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		log.Printf("Ошибка инспекции контейнера %s: %v", containerID, err)
		return nil
	}

	if containerInfo.NetworkSettings == nil || len(containerInfo.NetworkSettings.Networks) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	var results []models.PingResult
	var resultsLock sync.Mutex

	for _, network := range containerInfo.NetworkSettings.Networks {
		ip := network.IPAddress
		if ip != "" {
			for port := range containerInfo.NetworkSettings.Ports {
				wg.Add(1)
				go func(ip, port string) {
					defer wg.Done()
					pingTime, date, success := pingContainer(ip, port)
					if success {
						resultsLock.Lock()
						results = append(results, models.PingResult{
							IP:       ip,
							PingTime: pingTime,
							Date:     date,
						})
						resultsLock.Unlock()
					}
				}(ip, port.Port())
			}
		}
	}

	wg.Wait()

	return results
}

func pingContainer(ip string, port string) (string, string, bool) {
	pingTimeout := 5 * time.Second

	timezone := os.Getenv("TIME_ZONE")
	if timezone == "" {
		timezone = "UTC"
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Printf("Ошибка загрузки часового пояса: %v", err)
		location = time.UTC
	}

	startTime := time.Now().In(location)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", ip, port), pingTimeout)
	if err != nil {
		log.Printf("Ошибка пинга %s:%s %v", ip, port, err)
		return "", "", false
	}
	defer conn.Close()

	pingTime := time.Since(startTime).String()

	return pingTime, startTime.Format("02:01:2006:15:04:05.000"), true
}
