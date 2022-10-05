package repository

import (
	"github.com/golang-intensive/model"
)

type CustomerRespository interface {
	GetAll() []model.Customer
	Create(customer model.Customer) model.Customer
}

type customerRespository struct {
	db []model.Customer
}

func (c *customerRespository) GetAll() []model.Customer {
	var customer []model.Customer
	for _, cst := range c.db {
		customer = append(customer, cst)
	}
	return customer
}

func (c *customerRespository) Create(customer model.Customer) model.Customer {
	c.db = append(c.db, customer)
	return customer
}

func NewCustomerRepoistory() CustomerRespository {
	cstRepo := new(customerRespository)
	return cstRepo
}
