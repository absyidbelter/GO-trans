package usecase

import (
	"time"

	model "GO-Payment/internal/model/entity"
	"GO-Payment/internal/repository"
)

type LogUsecase interface {
	SaveLog(userID uint, event string) error
	GetAllLogs() ([]*model.Log, error)
}

type logUsecase struct {
	logRepository repository.LogRepository
}

func NewLogUsecase(logRepo repository.LogRepository) LogUsecase {
	return &logUsecase{
		logRepository: logRepo,
	}
}

func (lu *logUsecase) SaveLog(userID uint, event string) error {
	log := &model.Log{
		UserID:    userID,
		Event:     event,
		CreatedAt: time.Now(),
	}
	return lu.logRepository.Save(log)
}

func (lu *logUsecase) GetAllLogs() ([]*model.Log, error) {
	logs, err := lu.logRepository.GetAllLogs()
	if err != nil {
		return nil, err
	}
	return logs, nil
}
