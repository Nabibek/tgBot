package service

import (
	"tgBot/internal/domain"
	"tgBot/internal/repository"
)

type SubscriberService struct {
	repo repository.SubscriberRepository
}

func NewSubscriberService(repo repository.SubscriberRepository) *SubscriberService {
	return &SubscriberService{repo: repo}
}

func (s *SubscriberService) Subscribe(chatID int64, firstName, username string) error {
	subscriber := &domain.Subscriber{
		ChatID:    chatID,
		FirstName: firstName,
		Username:  username,
		Active:    true,
	}

	return s.repo.Add(subscriber)
}

func (s *SubscriberService) Unsubscribe(chatID int64) error {
	return s.repo.Remove(chatID)
}

func (s *SubscriberService) IsSubscribed(chatID int64) bool {
	return s.repo.Exists(chatID)
}

func (s *SubscriberService) GetSubscribers() ([]*domain.Subscriber, error) {
	return s.repo.GetAll()
}

func (s *SubscriberService) GetSubscriberCount() int {
	return s.repo.Count()
}
