package repository

import (
	"simple-payment/model"
	"simple-payment/util"

	"github.com/jmoiron/sqlx"
)

type LogHistoryRepository interface {
	LogHistory() (*[]model.LogHistory, error)
}

type logHistoryRepository struct {
	db *sqlx.DB
}

func (lhr *logHistoryRepository) LogHistory() (*[]model.LogHistory, error) {
	logHistory := new([]model.LogHistory)

	if err := lhr.db.Select(logHistory, util.ALL_HISTORIES); err != nil {
		return nil, err
	}
	return logHistory, nil
}

func NewLogHistoryRepository(db *sqlx.DB) LogHistoryRepository {
	return &logHistoryRepository{db: db}
}
