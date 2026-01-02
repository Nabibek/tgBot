package app

import (
	"tgBot/internal/config"
	"tgBot/pkg/logger"
)

type App struct {
	config *config.Config
	logger *logger.Logger
	bot    *Bot
}

func New() (*App, error) {
	// Загружаем конфигурацию
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Инициализируем логгер
	l, err := logger.New(cfg.LogFile, true)
	if err != nil {
		return nil, err
	}

	// Инициализируем бота
	bot, err := NewBot(cfg, l)
	if err != nil {
		return nil, err
	}

	return &App{
		config: cfg,
		logger: l,
		bot:    bot,
	}, nil
}

func (a *App) Run() error {
	defer a.logger.Close()

	a.logger.Info("Application started")
	a.bot.Start()

	return nil
}
