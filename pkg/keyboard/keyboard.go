package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type KeyboardType string

const (
	Inline KeyboardType = "inline"
	Reply  KeyboardType = "reply"
	Remove KeyboardType = "remove"
)

// CreateMainMenu создает основное меню
func CreateMainMenu(chatID int64, keyboardType KeyboardType) tgbotapi.Chattable {
	text := "Выберите действие:"

	switch keyboardType {
	case Inline:
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("📝 Получить цитату", "get_quote"),
				tgbotapi.NewInlineKeyboardButtonData("⭐ Подписаться", "subscribe"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("📊 Статистика", "stats"),
				tgbotapi.NewInlineKeyboardButtonData("❓ Помощь", "help"),
			),
		)
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = keyboard
		return msg

	case Reply:
		keyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("📝 Цитата"),
				tgbotapi.NewKeyboardButton("⭐ Подписаться"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("📊 Статистика"),
				tgbotapi.NewKeyboardButton("❓ Помощь"),
			),
		)
		keyboard.OneTimeKeyboard = true
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = keyboard
		return msg

	case Remove:
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		return msg

	default:
		return tgbotapi.NewMessage(chatID, text)
	}
}

// CreateSubscriptionKeyboard создает клавиатуру для подписки
func CreateSubscriptionKeyboard(chatID int64) tgbotapi.Chattable {
	text := "Вы хотите получать мотивационные цитаты ежедневно?"

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅ Да, подписаться", "subscribe_yes"),
			tgbotapi.NewInlineKeyboardButtonData("❌ Нет, спасибо", "subscribe_no"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	return msg
}
