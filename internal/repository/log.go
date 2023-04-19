package repository

import (
	model "GO-Payment/internal/model/entity"
	"database/sql"
	"fmt"
)

type LogRepository interface {
	Save(log *model.Log) error
	GetAllLogs() ([]*model.Log, error)
}

type logRepository struct {
	db *sql.DB
}

func NewLogRepository(db *sql.DB) LogRepository {
	return &logRepository{
		db: db,
	}
}

func (r *logRepository) Save(log *model.Log) error {
	if r.db == nil {
		return fmt.Errorf("database connection is nil")
	}
	query := `INSERT INTO logs (user_id, event, created_at) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, log.UserID, log.Event, log.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *logRepository) GetAllLogs() ([]*model.Log, error) {
	if r.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	query := `SELECT id, user_id, event, created_at FROM logs`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logs := []*model.Log{}
	for rows.Next() {
		log := &model.Log{}
		err := rows.Scan(&log.ID, &log.UserID, &log.Event, &log.CreatedAt)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}
