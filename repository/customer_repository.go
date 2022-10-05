package repository

import (
	"fmt"

	"github.com/golang-intensive/model"
)

type CustomerRespository interface {
	GetAll() []model.Customer
	Create(customer model.Customer) model.Customer
	Delete(id string) string
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

func (c *customerRespository) Delete(id string) string {
	var customers []model.Customer
	for i := 0; i < len(c.db); i++ {
		if c.db[i].Id == id {
			customers = append(c.db[:i], c.db[i+1:]...)
		}
	}
	c.db = customers
	return fmt.Sprintf("Data customer with ID %s deleted", id)
}

func NewCustomerRepoistory() CustomerRespository {
	cstRepo := new(customerRespository)
	return cstRepo
}
