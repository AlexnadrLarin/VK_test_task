package router

import (
	"github.com/gorilla/mux"

	"backend/api"
	"backend/internal/database"
)

func SetupRouter(repo *database.Repository) *mux.Router {
	r := mux.NewRouter()

	// r.HandleFunc("/ping-results", getPingResults).Methods("GET")
	r.HandleFunc("/api/v1/ping-results", api.CreatePingResult(repo)).Methods("POST")

	return r
}