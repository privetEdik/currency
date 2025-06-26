package service

import (
	"currency/internal/model"
	"errors"
	"sync"
)

type Service struct {
	mu         sync.Mutex
	currencies []model.Currency
	nextID     int
}

func NewService() *Service {
	return &Service{
		currencies: []model.Currency{},
		nextID:     1,
	}
}

func (s *Service) GetAll() []model.Currency {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.currencies
}

func (s *Service) GetByCode(code string) (*model.Currency, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, c := range s.currencies {
		if c.Code == code {
			return &c, nil
		}
	}
	return nil, errors.New("currency not found")
}

func (s *Service) Add(c model.Currency) model.Currency {
	s.mu.Lock()
	defer s.mu.Unlock()

	c.ID = s.nextID
	s.nextID++
	s.currencies = append(s.currencies, c)
	return c
}
