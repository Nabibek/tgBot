package domain

// BotHandler интерфейс для обработчиков
type BotHandler interface {
	HandleStart(message Message) error
	HandleSubscribe(message Message) error
	HandleUnsubscribe(message Message) error
	HandleQuote(message Message) error
	HandleHelp(message Message) error
}

// Message представляет сообщение от пользователя
type Message struct {
	ChatID   int64
	UserID   int64
	Text     string
	Username string
	Command  string
}
