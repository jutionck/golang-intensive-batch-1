package main

import (
	"fmt"

	"github.com/golang-intensive/model"
	"github.com/golang-intensive/repository"
	"github.com/golang-intensive/usecase"
)

func main() {

	cstRepo := repository.NewCustomerRepoistory()
	cstUseCase := usecase.NewCustomerUseCase(cstRepo)

	// Create
	newCustomer01, err := cstUseCase.RegisterNewCustomer(model.Customer{Id: "C001", Name: "Jution", Address: "USA", Job: "Actrist"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newCustomer01)

	// GetAll
	fmt.Println(cstUseCase.FindAllCustomer())

	// Delete
	// cstUseCase.DeleteCustomer("C001")
	// fmt.Println(cstUseCase.FindAllCustomer())

	// Update
	cstUseCase.UpdateCustomer(model.Customer{Id: "C001", Name: "Bintang", Address: "UK", Job: "Actrist"})
	fmt.Println(cstUseCase.FindAllCustomer())

}
