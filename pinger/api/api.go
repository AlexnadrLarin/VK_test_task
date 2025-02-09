package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"pinger/models"
)

func SendResultsToAPI(results models.PingResults) {
	if isEmpty(&results) {
		log.Println("Нет успешно пропингованных контейнеров")
		return
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Printf("Ошибка сериализации JSON: %v", err)
		return
	}

	apiURL := os.Getenv("BACKEND_API_URL")
	if apiURL == "" {
		log.Println("BACKEND_API_URL не задан")
		return
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Ошибка создания запроса: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка отправки данных на API: %v", err)
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Ошибка закрытия тела ответа: %v", err)
		}
	}()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Println("Данные успешно отправлены на API")
	} else {
		log.Printf("Ошибка отправки данных, код ответа: %d\n", resp.StatusCode)
	}
}

func isEmpty(results *models.PingResults) bool {
	return results == nil || len(results.Results) == 0
}
