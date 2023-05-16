package repositories

import (
	"database/sql"
	"log"

	"github.com/Uchel/chels-laundry/models"
)

type CustomerRepo interface {
	GetPasscode(customer *models.Customer) string
}

type customerRepo struct {
	db *sql.DB
}

func NewCustomerRepo(db *sql.DB) CustomerRepo {
	repo := new(customerRepo)
	repo.db = db
	return repo

}

func (r customerRepo) GetPasscode(customer *models.Customer) string {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	// fill query
	name := r.InsertCustomer(customer, tx)
	passcode := r.GetPassCode(name, tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err.Error()
	} else {
		return passcode
	}
}

func (r customerRepo) InsertCustomer(customer *models.Customer, tx *sql.Tx) string {
	query := "insert into customer (name,phone) values($1,$2)"

	_, err := tx.Exec(query, &customer.Name, &customer.Phone)
	Validate(err, "Insert", tx)
	return customer.Name
}

func (r customerRepo) GetPassCode(name string, tx *sql.Tx) string {
	query := "select pass_code from customer where name = $1"
	passcode := ""
	err := tx.QueryRow(query, name).Scan(&passcode)
	Validate(err, "Select", tx)
	return passcode
}
