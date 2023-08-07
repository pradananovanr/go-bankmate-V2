package usecase

import (
	"go-bankmate/model/entity"
	"go-bankmate/repository"
)

type CustomerUsecase interface {
	Add(newCustomer *entity.Customer) (entity.Customer, error)
	Remove(id int) error
	Login(username, password string) (string, error)
	Logout(id int) error
}

type customerUsecase struct {
	customerRepo repository.CustomerRepo
}

func (u *customerUsecase) Add(newCustomer *entity.Customer) (entity.Customer, error) {
	return u.customerRepo.Create(newCustomer)
}

func (u *customerUsecase) Remove(id int) error {
	return u.customerRepo.Delete(id)
}

func (u *customerUsecase) Login(username, password string) (string, error) {
	return u.customerRepo.Login(username, password)
}

func (u *customerUsecase) Logout(id int) error {
	return u.customerRepo.Logout(id)
}

func NewCustomerUsecase(customerRepo repository.CustomerRepo) CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}
