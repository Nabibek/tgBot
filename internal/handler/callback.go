package handler

import (
	"tgBot/pkg/keyboard"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackHandler struct {
	bot            *tgbotapi.BotAPI
	commandHandler *CommandHandler
}

func NewCallbackHandler(bot *tgbotapi.BotAPI, cmdHandler *CommandHandler) *CallbackHandler {
	return &CallbackHandler{
		bot:            bot,
		commandHandler: cmdHandler,
	}
}

func (h *CallbackHandler) Handle(callbackQuery *tgbotapi.CallbackQuery) {
	// Отвечаем на callback
	callback := tgbotapi.NewCallback(callbackQuery.ID, "")
	h.bot.Request(callback)

	switch callbackQuery.Data {
	case "get_quote":
		// Отправляем цитату
		h.commandHandler.HandleQuote(callbackQuery.Message)

	case "subscribe":
		// Показываем клавиатуру подписки
		msg := keyboard.CreateSubscriptionKeyboard(callbackQuery.Message.Chat.ID)
		h.bot.Send(msg)

	case "subscribe_yes":
		// Подписываем пользователя
		h.commandHandler.HandleSubscribe(callbackQuery.Message)

	case "subscribe_no":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"Хорошо, если передумаете — используйте команду /subscribe")
		h.bot.Send(msg)

	case "stats":
		// Показываем статистику
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"Статистика скоро будет доступна!")
		h.bot.Send(msg)

	case "help":
		h.commandHandler.HandleHelp(callbackQuery.Message)
	}
}
