package repository

import (
	"database/sql"
	"go-bankmate/model/entity"
	"log"
)

type LogRepo interface {
	ShowAll() ([]*entity.Log, error)
}

type logRepo struct {
	db *sql.DB
}

func (l *logRepo) ShowAll() ([]*entity.Log, error) {
	var logs []*entity.Log

	query := "SELECT * FROM t_log"
	row, err := l.db.Query(query)

	if err != nil {
		log.Println(err)
		return []*entity.Log{}, err
	}

	defer row.Close()
	for row.Next() {
		var log entity.Log
		if err := row.Scan(&log.ID_Log, &log.ID_Customer, &log.Activity, &log.Date_Time); err != nil {
			return []*entity.Log{}, err
		}
		logs = append(logs, &log)
	}
	if err := row.Err(); err != nil {
		return []*entity.Log{}, err
	}

	return logs, nil
}

func NewLogRepository(db *sql.DB) LogRepo {
	repo := new(logRepo)
	repo.db = db
	return repo
}
