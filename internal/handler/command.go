package handler

import (
	"fmt"
	"tgBot/internal/service"
	"tgBot/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandHandler struct {
	bot           *tgbotapi.BotAPI
	subscriberSvc *service.SubscriberService
	quotesSvc     *service.QuotesService
	logger        *logger.Logger
}

func NewCommandHandler(bot *tgbotapi.BotAPI, subscriberSvc *service.SubscriberService, quotesSvc *service.QuotesService, logger *logger.Logger) *CommandHandler {
	return &CommandHandler{
		bot:           bot,
		subscriberSvc: subscriberSvc,
		quotesSvc:     quotesSvc,
		logger:        logger,
	}
}

func (h *CommandHandler) HandleStart(message *tgbotapi.Message) {
	text := "👋 Добро пожаловать! Я мотивационный бот.\n\n" +
		"Я могу отправлять вам вдохновляющие цитаты несколько раз в день.\n\n" +
		"*Команды:*\n" +
		"/subscribe - подписаться на ежедневные цитаты\n" +
		"/unsubscribe - отписаться\n" +
		"/quote - получить случайную цитату\n" +
		"/help - справка"

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ParseMode = "Markdown"
	h.bot.Send(msg)

	h.logger.Info("New user: %s (%s) - ID: %d",
		message.From.FirstName,
		message.From.UserName,
		message.From.ID)
}

func (h *CommandHandler) HandleSubscribe(message *tgbotapi.Message) {
	if h.subscriberSvc.IsSubscribed(message.Chat.ID) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "✅ Вы уже подписаны на цитаты!")
		h.bot.Send(msg)
		return
	}

	err := h.subscriberSvc.Subscribe(
		message.Chat.ID,
		message.From.FirstName,
		message.From.UserName,
	)

	if err != nil {
		h.logger.Error("Failed to subscribe user %d: %v", message.Chat.ID, err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Произошла ошибка при подписке.")
		h.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID,
		"✅ Вы успешно подписались! Теперь вы будете получать мотивационные цитаты в 9:00 и 18:00.")
	h.bot.Send(msg)

	h.logger.Info("Subscribed: %s (%s) - ID: %d | Total: %d",
		message.From.FirstName,
		message.From.UserName,
		message.From.ID,
		h.subscriberSvc.GetSubscriberCount())
}

func (h *CommandHandler) HandleUnsubscribe(message *tgbotapi.Message) {
	if !h.subscriberSvc.IsSubscribed(message.Chat.ID) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Вы не подписаны на цитаты.")
		h.bot.Send(msg)
		return
	}

	err := h.subscriberSvc.Unsubscribe(message.Chat.ID)
	if err != nil {
		h.logger.Error("Failed to unsubscribe user %d: %v", message.Chat.ID, err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Произошла ошибка при отписке.")
		h.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Вы отписались от цитат.")
	h.bot.Send(msg)

	h.logger.Info("Unsubscribed: %s (%s) - ID: %d | Total: %d",
		message.From.FirstName,
		message.From.UserName,
		message.From.ID,
		h.subscriberSvc.GetSubscriberCount())
}

func (h *CommandHandler) HandleQuote(message *tgbotapi.Message) {
	quote, err := h.quotesSvc.GetRandomQuote()
	if err != nil {
		h.logger.Error("Failed to get random quote: %v", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "❌ Произошла ошибка при получении цитаты.")
		h.bot.Send(msg)
		return
	}

	text := "✨ *Мотивационная цитата:*\n\n" +
		quote + "\n\n" +
		"_Веди себя достойно! 💪_"

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ParseMode = "Markdown"
	h.bot.Send(msg)

	h.logger.Info("Quote sent to: %s (%d)", message.From.FirstName, message.Chat.ID)
}

func (h *CommandHandler) HandleHelp(message *tgbotapi.Message) {
	text := "📖 *Справка по командам:*\n\n" +
		"/start - главное меню\n" +
		"/subscribe - подписаться на цитаты\n" +
		"/unsubscribe - отписаться от цитат\n" +
		"/quote - получить случайную цитату\n" +
		"/help - показать эту справку"

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ParseMode = "Markdown"
	h.bot.Send(msg)
}
func (h *CommandHandler) HandleStats(message *tgbotapi.Message) {
	count := h.subscriberSvc.GetSubscriberCount()

	text := fmt.Sprintf("📊 *Статистика бота:*\n\n"+
		"• Подписчиков: %d\n"+
		"• Цитат в базе: %d\n"+
		"• Бот запущен", count)

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ParseMode = "Markdown"
	h.bot.Send(msg)
}
