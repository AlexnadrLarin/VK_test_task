package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"pinger/models"
)

func SendResultsToAPI(results models.PingResults) {
	if !isEmpty(results) {
		jsonData, err := json.Marshal(results)
		if err != nil {
			log.Printf("Ошибка сериализации JSON: %v", err)
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

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Ошибка отправки данных на API")
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			log.Println("Данные успешно отправлены на API")
		} else {
			log.Printf("Ошибка отправки данных, код ответа: %d\n", resp.StatusCode)
		}
	} else {
		log.Println("Нет успешно пропингованных контейнеров")
	}
}

func isEmpty(results models.PingResults) bool {
	return len(results.Results) == 0
}
