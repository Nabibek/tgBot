package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken     string
	Debug        bool
	QuoteTimes   []string
	LogFile      string
	UpdateOffset int
	Timeout      int
}

func Load() (*Config, error) {
	// Загружаем .env файл
	_ = godotenv.Load()

	return &Config{
		BotToken:     getEnv("TELEGRAM_BOT_TOKEN", ""),
		Debug:        getEnvAsBool("BOT_DEBUG", false),
		QuoteTimes:   getEnvAsSlice("QUOTE_TIMES", []string{"09:00", "18:00"}),
		LogFile:      getEnv("LOG_FILE", "bot_logs.txt"),
		UpdateOffset: getEnvAsInt("UPDATE_OFFSET", 0),
		Timeout:      getEnvAsInt("TIMEOUT", 60),
	}, nil
}

// Вспомогательные функции для чтения переменных окружения
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return []string{value} // Простая реализация
	}
	return defaultValue
}
