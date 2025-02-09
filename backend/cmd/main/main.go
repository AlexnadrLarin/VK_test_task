package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/cors"

	"backend/router"
	"backend/internal/database"
)

func main() {
	repo, err := database.NewRepository(10, 2*time.Second)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer repo.Close()

	r := router.SetupRouter(repo)

	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		log.Fatal("CORS_ALLOWED_ORIGINS не установлена")
	}

	originsList := strings.Split(allowedOrigins, ",")
	for i, origin := range originsList {
		originsList[i] = strings.TrimSpace(origin)
	}

	corsOptions := cors.Options{
		AllowedOrigins:     originsList, 
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials:   true,
	}

	corsHandler := cors.New(corsOptions).Handler(r)

	backendPort := os.Getenv("BACKEND_PORT")
	if backendPort == "" {
		log.Fatal("BACKEND_PORT не установлена")
	}

	log.Println("Сервер запущен")
	log.Fatal(http.ListenAndServe(":"+backendPort, corsHandler))
}
