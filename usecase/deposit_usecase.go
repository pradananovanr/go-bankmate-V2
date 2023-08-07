package usecase

import (
	"go-bankmate/model/entity"
	"go-bankmate/repository"
)

type DepositUsecase interface {
	Add(id_customer int, token string, deposit *entity.DepositRequest) (*entity.Deposit, error)
	FindOne(id_customer, id_deposit int, token string) (*entity.Deposit, error)
	FindAll(id_customer int, token string) ([]*entity.Deposit, error)
}

type depositUsecase struct {
	depositRepo repository.DepositRepo
}

func (u *depositUsecase) Add(id_customer int, token string, deposit *entity.DepositRequest) (*entity.Deposit, error) {
	return u.depositRepo.CreateDeposit(id_customer, token, deposit)
}

func (u *depositUsecase) FindOne(id_customer, id_deposit int, token string) (*entity.Deposit, error) {
	return u.depositRepo.GetDeposit(id_customer, id_deposit, token)
}

func (u *depositUsecase) FindAll(id_customer int, token string) ([]*entity.Deposit, error) {
	return u.depositRepo.GetAllDeposit(id_customer, token)
}

func NewDepositUsecase(depositRepo repository.DepositRepo) DepositUsecase {
	return &depositUsecase{
		depositRepo: depositRepo,
	}
}
