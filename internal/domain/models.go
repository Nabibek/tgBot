package domain

// Subscriber представляет подписчика
type Subscriber struct {
	ChatID    int64
	FirstName string
	Username  string
	Active    bool
}

// Quote представляет цитату
type Quote struct {
	ID     int
	Text   string
	Author string
	Tags   []string
}
