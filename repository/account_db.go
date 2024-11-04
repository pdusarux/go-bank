package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) Create(account Account) (*Account, error) {
	query := "insert into accounts (customer_id, opening_date, account_type, amount, status) values ($1, $2, $3, $4, $5) RETURNING account_id"
	err := r.db.QueryRow(
		query,
		account.CustomerID,
		account.OpeningDate,
		account.AccountType,
		account.Amount,
		account.Status,
	).Scan(&account.AccountID)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r accountRepositoryDB) GetAll(customerID int) ([]Account, error) {
	accounts := []Account{}
	query := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where customer_id = $1"

	err := r.db.Select(&accounts, query, customerID)
	if err != nil {
		log.Printf("Error fetching accounts for customerID %d: %v", customerID, err)
		return nil, err
	}

	return accounts, nil
}
