package service

import (
	"fmt"
	"tgBot/internal/domain"
	"tgBot/internal/repository"
)

type QuotesService struct {
	repo repository.QuotesRepository
}

func NewQuotesService(repo repository.QuotesRepository) *QuotesService {
	return &QuotesService{repo: repo}
}

func (s *QuotesService) GetRandomQuote() (string, error) {
	quote, err := s.repo.GetRandom()
	if err != nil {
		return "", err
	}

	if quote == nil {
		return "Нет доступных цитат", nil
	}

	return fmt.Sprintf("✨ *%s*\n\n_— %s_", quote.Text, quote.Author), nil
}

func (s *QuotesService) GetAllQuotes() ([]*domain.Quote, error) {
	return s.repo.GetAll()
}

func (s *QuotesService) AddQuote(text, author string) error {
	// Генерируем ID (в реальном приложении используйте базу данных)
	quotes, _ := s.repo.GetAll()
	id := len(quotes) + 1

	quote := &domain.Quote{
		ID:     id,
		Text:   text,
		Author: author,
	}

	return s.repo.Add(quote)
}
