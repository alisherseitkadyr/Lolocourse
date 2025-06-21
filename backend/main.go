package main

import (
	"landing/backend/storage"
	"landing/backend/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port="8080"
	}
	storage.InitDB()

	http.HandleFunc("/userdata", handlers.HandleLead)

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	log.Println("Server is run on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
