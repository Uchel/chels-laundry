package usecases

import (
	"fmt"

	"github.com/Uchel/chels-laundry/models"
	"github.com/Uchel/chels-laundry/repositories"
)

type CustomerUsecases interface {
	GetPasscode(customer *models.Customer) string
}

type customerUsecases struct {
	customerRepo repositories.CustomerRepo
}

// =======================================================================================================
func (u customerUsecases) GetPasscode(customer *models.Customer) string {
	fmt.Println("wooouy")
	return u.customerRepo.GetPasscode(customer)
}

// =======================================================================================================
func NewCustomerUC(customerRepo repositories.CustomerRepo) CustomerUsecases {
	return &customerUsecases{
		customerRepo: customerRepo,
	}
}
