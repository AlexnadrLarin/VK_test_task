package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"backend/router"
)

func main() {
	r := router.SetupRouter()

	fmt.Println("Сервер запущен")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("BACKEND_PORT"), r))
}
