package main

import (
	"log"
	"tgBot/internal/app"
)

func main() {
	// Создаем и запускаем приложение
	application, err := app.New()
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}

	// Запускаем бота
	if err := application.Run(); err != nil {
		log.Fatalf("Bot stopped with error: %v", err)
	}
}
