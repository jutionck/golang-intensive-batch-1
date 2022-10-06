package usecase

import (
	"errors"

	"github.com/golang-intensive/model"
	"github.com/golang-intensive/repository"
)

type CustomerUseCase interface {
	RegisterNewCustomer(customer model.Customer) (model.Customer, error)
	FindAllCustomer() []model.Customer
	UpdateCustomer(customer model.Customer) model.Customer
	DeleteCustomer(id string) string
}

type customerUseCase struct {
	customerRespository repository.CustomerRespository
}

func (c *customerUseCase) RegisterNewCustomer(customer model.Customer) (model.Customer, error) {
	if customer.Name == "" {
		return model.Customer{}, errors.New("Customer name can't be empty")
	}
	return c.customerRespository.Create(customer), nil
}
func (c *customerUseCase) FindAllCustomer() []model.Customer {
	return c.customerRespository.GetAll()
}
func (c *customerUseCase) UpdateCustomer(customer model.Customer) model.Customer {
	return c.customerRespository.Update(customer)
}
func (c *customerUseCase) DeleteCustomer(id string) string {
	return c.customerRespository.Delete(id)
}

func NewCustomerUseCase(customerRepository repository.CustomerRespository) CustomerUseCase {
	uc := new(customerUseCase)
	uc.customerRespository = customerRepository
	return uc
}
