package usecase

import (
	"go-bankmate/model/entity"
	"go-bankmate/repository"
)

type PaymentUsecase interface {
	Create(id_customer int, token string, payment *entity.PaymentRequest) (*entity.Payment, error)
	FindOne(id_customer, id_payment int, token string) (*entity.Payment, error)
	FindAll(id_customer int, token string) ([]*entity.Payment, error)
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepo
}

func (u *paymentUsecase) Create(id_customer int, token string, payment *entity.PaymentRequest) (*entity.Payment, error) {
	return u.paymentRepo.CreatePayment(id_customer, token, payment)
}

func (u *paymentUsecase) FindOne(id_customer, id_payment int, token string) (*entity.Payment, error) {
	return u.paymentRepo.GetPayment(id_customer, id_payment, token)
}

func (u *paymentUsecase) FindAll(id_customer int, token string) ([]*entity.Payment, error) {
	return u.paymentRepo.GetAllPayment(id_customer, token)
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepo) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
	}
}
