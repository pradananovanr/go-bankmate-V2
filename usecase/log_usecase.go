package usecase

import (
	"go-bankmate/model/entity"
	"go-bankmate/repository"
)

type LogUsecase interface {
	FindAll() ([]*entity.Log, error)
}

type logUsecase struct {
	logRepo repository.LogRepo
}

func (u *logUsecase) FindAll() ([]*entity.Log, error) {
	return u.logRepo.ShowAll()
}

func NewLogUsecase(logRepo repository.LogRepo) LogUsecase {
	return &logUsecase{
		logRepo: logRepo,
	}
}
