package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"


	"backend/api"
	"backend/internal/database"
)

func SetupRouter(repo *database.Repository) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/ping-results", api.GetAllPingResults(repo)).Methods("GET")
	r.HandleFunc("/api/v1/ping-results", api.CreatePingResult(repo)).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Маршрут не найден: %s %s", r.Method, r.URL)
		http.Error(w, "Ресурс не найден", http.StatusNotFound)
	})

	return r
}
