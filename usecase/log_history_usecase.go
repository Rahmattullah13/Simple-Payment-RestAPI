package usecase

import (
	"simple-payment/model"
	"simple-payment/repository"
)

type LogHistoryUsaCase interface {
	LogHistory() (*[]model.LogHistory, error)
}

type logHistoryUseCase struct {
	repo repository.LogHistoryRepository
}

func (lhu *logHistoryUseCase) LogHistory() (*[]model.LogHistory, error) {
	return lhu.repo.LogHistory()
}

func NewLogHistoryUseCase(repo repository.LogHistoryRepository) LogHistoryUsaCase {
	return &logHistoryUseCase{repo: repo}
}
