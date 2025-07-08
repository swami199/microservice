package service

import "github.com/swami199/microservice/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerSerivice struct {
	repo domain.CustomerRepositoryStub
}

func (s DefaultCustomerSerivice) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()w

}
