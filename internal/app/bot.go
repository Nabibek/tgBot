package app

import (
	"tgBot/internal/config"
	"tgBot/internal/handler"
	"tgBot/internal/repository"
	"tgBot/internal/service"
	"tgBot/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api       *tgbotapi.BotAPI
	config    *config.Config
	logger    *logger.Logger
	handler   *handler.CommandHandler
	scheduler *service.Scheduler
}

func NewBot(cfg *config.Config, l *logger.Logger) (*Bot, error) {
	// Инициализируем Telegram Bot API
	botAPI, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, err
	}

	botAPI.Debug = cfg.Debug

	// Инициализируем репозитории
	subscriberRepo := repository.NewInMemorySubscriberRepo()
	quotesRepo := repository.NewInMemoryQuotesRepo()

	// Инициализируем сервисы
	subscriberSvc := service.NewSubscriberService(subscriberRepo)
	quotesSvc := service.NewQuotesService(quotesRepo)

	// Инициализируем обработчики
	cmdHandler := handler.NewCommandHandler(botAPI, subscriberSvc, quotesSvc, l)

	// Инициализируем планировщик
	scheduler := service.NewScheduler(botAPI, subscriberSvc, quotesSvc, l)

	return &Bot{
		api:       botAPI,
		config:    cfg,
		logger:    l,
		handler:   cmdHandler,
		scheduler: scheduler,
	}, nil
}

func (b *Bot) Start() {
	b.logger.Info("Bot started: @%s", b.api.Self.UserName)

	// Запускаем планировщик
	b.scheduler.Start()
	defer b.scheduler.Stop()

	// Настраиваем получение обновлений
	u := tgbotapi.NewUpdate(b.config.UpdateOffset)
	u.Timeout = b.config.Timeout

	updates := b.api.GetUpdatesChan(u)

	// Обрабатываем обновления
	for update := range updates {
		b.handleUpdate(update)
	}
}

func (b *Bot) handleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	message := update.Message

	// Обрабатываем команды
	switch {
	case message.IsCommand():
		b.handleCommand(message)
	default:
		b.handleMessage(message)
	}
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		b.handler.HandleStart(message)
	case "subscribe":
		b.handler.HandleSubscribe(message)
	case "unsubscribe":
		b.handler.HandleUnsubscribe(message)
	case "quote":
		b.handler.HandleQuote(message)
	case "help":
		b.handler.HandleHelp(message)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID,
			"❌ Неизвестная команда. Введите /help для справки.")
		b.api.Send(msg)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	// Можно добавить обработку обычных сообщений
	// Например, ответ на приветствие
	text := message.Text
	if contains([]string{"привет", "здравствуй", "hello", "hi"}, text) {
		msg := tgbotapi.NewMessage(message.Chat.ID,
			"Привет! 👋 Введите /help для списка команд.")
		b.api.Send(msg)
	}
}

// Вспомогательная функция
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
