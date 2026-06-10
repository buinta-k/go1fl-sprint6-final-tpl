package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	logger.Println("Инициализация настроек сервера...")

	srv := server.New(logger)

	err := srv.Start()

	if err != nil {
		logger.Fatalf("Ошибка %v", err)
	}
	
}
