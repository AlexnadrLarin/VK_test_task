package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"backend/router"
	"backend/internal/database"
)

func main() {
	repo, err := database.NewRepository()
    if err != nil {
        log.Fatalf("Ошибка подключения к Postgres: %v", err)
    }
    defer repo.Close()

	r := router.SetupRouter(repo)

	fmt.Println("Сервер запущен")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("BACKEND_PORT"), r))
}
