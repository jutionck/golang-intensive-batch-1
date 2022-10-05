package main

import (
	"fmt"

	"github.com/golang-intensive/model"
	"github.com/golang-intensive/repository"
)

func main() {

	cstRepo := repository.NewCustomerRepoistory()

	// Create
	newCustomer := cstRepo.Create(model.Customer{Id: "C001", Name: "Jack", Address: "USA", Job: "Actrist"})
	fmt.Println(newCustomer)

	// GetAll
	customers := cstRepo.GetAll()
	fmt.Println(customers)
}
