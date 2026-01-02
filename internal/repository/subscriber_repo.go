package repository

import (
	"sync"
	"tgBot/internal/domain"
)

type SubscriberRepository interface {
	Add(subscriber *domain.Subscriber) error
	Remove(chatID int64) error
	Get(chatID int64) (*domain.Subscriber, error)
	GetAll() ([]*domain.Subscriber, error)
	Exists(chatID int64) bool
	Count() int
}

type InMemorySubscriberRepo struct {
	subscribers map[int64]*domain.Subscriber
	mu          sync.RWMutex
}

func NewInMemorySubscriberRepo() *InMemorySubscriberRepo {
	return &InMemorySubscriberRepo{
		subscribers: make(map[int64]*domain.Subscriber),
	}
}

func (r *InMemorySubscriberRepo) Add(subscriber *domain.Subscriber) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.subscribers[subscriber.ChatID] = subscriber
	return nil
}

func (r *InMemorySubscriberRepo) Remove(chatID int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.subscribers, chatID)
	return nil
}

func (r *InMemorySubscriberRepo) Get(chatID int64) (*domain.Subscriber, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	subscriber, exists := r.subscribers[chatID]
	if !exists {
		return nil, nil
	}
	return subscriber, nil
}

func (r *InMemorySubscriberRepo) GetAll() ([]*domain.Subscriber, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	subscribers := make([]*domain.Subscriber, 0, len(r.subscribers))
	for _, sub := range r.subscribers {
		subscribers = append(subscribers, sub)
	}

	return subscribers, nil
}

func (r *InMemorySubscriberRepo) Exists(chatID int64) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.subscribers[chatID]
	return exists
}

func (r *InMemorySubscriberRepo) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.subscribers)
}
