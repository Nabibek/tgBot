package service

import (
	"fmt"
	"time"

	"tgBot/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Scheduler struct {
	bot           *tgbotapi.BotAPI
	subscriberSvc *SubscriberService
	quotesSvc     *QuotesService
	logger        *logger.Logger
	stopChan      chan struct{}
}

func NewScheduler(bot *tgbotapi.BotAPI, subscriberSvc *SubscriberService, quotesSvc *QuotesService, logger *logger.Logger) *Scheduler {
	return &Scheduler{
		bot:           bot,
		subscriberSvc: subscriberSvc,
		quotesSvc:     quotesSvc,
		logger:        logger,
		stopChan:      make(chan struct{}),
	}
}

func (s *Scheduler) Start() {
	go s.run()
}

func (s *Scheduler) Stop() {
	close(s.stopChan)
}

func (s *Scheduler) run() {
	ticker := time.NewTicker(1 * time.Minute) // Проверяем каждую минуту
	defer ticker.Stop()

	for {
		select {
		case <-s.stopChan:
			s.logger.Info("Scheduler stopped")
			return
		case <-ticker.C:
			s.sendScheduledQuotes()
		}
	}
}

func (s *Scheduler) sendScheduledQuotes() {
	now := time.Now()
	hour := now.Hour()
	minute := now.Minute()

	// Отправляем в 9:00 и 18:00
	if (hour >= 9 && hour < 18) && minute == 0 {
		s.sendQuotesToAllSubscribers()
	}
}

func (s *Scheduler) sendQuotesToAllSubscribers() {
	subscribers, err := s.subscriberSvc.GetSubscribers()
	if err != nil {
		s.logger.Error("Failed to get subscribers: %v", err)
		return
	}

	s.logger.Info("Sending quotes to %d subscribers", len(subscribers))

	for _, subscriber := range subscribers {
		quote, err := s.quotesSvc.GetRandomQuote()
		if err != nil {
			s.logger.Error("Failed to get quote for subscriber %d: %v", subscriber.ChatID, err)
			continue
		}

		text := fmt.Sprintf("✨ *Ваша ежедневная мотивация:*\n\n%s\n\n_Имейте веру в себя! 💪_", quote)

		msg := tgbotapi.NewMessage(subscriber.ChatID, text)
		msg.ParseMode = "Markdown"

		if _, err := s.bot.Send(msg); err != nil {
			s.logger.Error("Failed to send quote to %s (%d): %v",
				subscriber.FirstName, subscriber.ChatID, err)
		}
	}

	s.logger.Info("Finished sending quotes to %d subscribers", len(subscribers))
}
