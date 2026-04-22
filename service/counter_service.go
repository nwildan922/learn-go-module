package service

import (
	"context"
	"time"

	"github.com/nwildan922/learn-go-module/model"
	"github.com/nwildan922/learn-go-module/repository"
)

type CounterService struct {
	repo  *repository.CounterRepository
	appId string
}

func NewCounterService(repo *repository.CounterRepository, appId string) *CounterService {
	return &CounterService{repo: repo, appId: appId}
}

// Business logic
func (s *CounterService) SaveCounter(
	ctx context.Context,
	counter int32,
	timestamp time.Time,
) error {

	data := &model.Counter{
		Counter:   counter,
		Timestamp: timestamp,
		AppId:     s.appId,
	}

	return s.repo.Create(ctx, data)
}
