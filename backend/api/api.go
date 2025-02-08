package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend/models"
)

func CreatePingResult(w http.ResponseWriter, r *http.Request) {
	var pingResults models.PingResults

	if err := json.NewDecoder(r.Body).Decode(&pingResults); err != nil {
		respondWithJSON(w, http.StatusBadRequest, ErrorMessage("Неверный формат данных"))
		return
	}

	fmt.Println(pingResults)

	respondWithJSON(w, http.StatusOK, SuccessMessage("Результат добавлен в базу данных"))
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
