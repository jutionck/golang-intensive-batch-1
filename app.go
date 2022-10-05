package main

import (
	"fmt"

	"github.com/golang-intensive/model"
	"github.com/golang-intensive/repository"
)

func main() {

	cstRepo := repository.NewCustomerRepoistory()

	// Create
	cstRepo.Create(model.Customer{Id: "C001", Name: "Jack", Address: "USA", Job: "Actrist"})
	cstRepo.Create(model.Customer{Id: "C002", Name: "Bulan", Address: "UK", Job: "Actrist"})

	// GetAll
	fmt.Println("customers 01", cstRepo.GetAll())

	// Delete
	cstRepo.Delete("C001")

	// Get All Again
	fmt.Println("customers 02", cstRepo.GetAll())

}
