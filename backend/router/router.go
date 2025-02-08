package router

import (
	"github.com/gorilla/mux"

	"backend/api"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/ping-results", api.CreatePingResult).Methods("POST")

	return r
}