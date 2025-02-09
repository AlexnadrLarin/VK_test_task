package api

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/internal/database"
	"backend/models"
)

func CreatePingResult(repo *database.Repository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer r.Body.Close()

		var pingResults models.PingResults

		if err := json.NewDecoder(r.Body).Decode(&pingResults); err != nil {
			log.Printf("Ошибка при декодировании данных: %v", err)
			respondWithJSON(w, http.StatusBadRequest, ErrorMessage("Неверный формат данных"))
			return
		}

		err := repo.UpsertPingResults(pingResults)
        if err != nil {
            log.Printf("Ошибка вставки данных в базу данных: %v", err)
            respondWithJSON(w, http.StatusInternalServerError, ErrorMessage("Не удалось обработать запрос"))
            return
        }

		respondWithJSON(w, http.StatusCreated, SuccessMessage("Результат добавлен в базу данных"))
	}
}

func GetAllPingResults(repo *database.Repository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        pingResults, err := repo.GetPingResults()
        if err != nil {
            log.Printf("Ошибка получения данных из базы: %v", err)
            respondWithJSON(w, http.StatusInternalServerError, ErrorMessage("Не удалось обработать запрос"))
            return
        }

        respondWithJSON(w, http.StatusOK, pingResults)
    }
}

func ErrorMessage(message string) map[string]string {
    return map[string]string{"error": message}
}

func SuccessMessage(message string) map[string]string {
    return map[string]string{"message": message}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}
