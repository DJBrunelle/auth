package service

import (
	"errors"
	"time"
)

type Account struct {
	ID             int       `json:"id" db:"id"`
	FirstName      string    `json:"firstName" db:"first_name"`
	LastName       string    `json:"lastName" db:"last_name"`
	Email          string    `json:"email" db:"email"`
	HashedPassword string    `db:"hashed_password"`
	Admin          bool      `json:"admin" db:"admin"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
}

func (db *DB) CreateAccount(acc Account) error {

	query := `SELECT id
                FROM account
                WHERE email=$1`

	var id int

	if err := db.Get(&id, query, acc.Email); err == nil {
		err = errors.New("Error creating new account: email already exists")
		return err
	}

	query = `INSERT INTO account (first_name,last_name,email,hashed_password,admin,created_at,updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7)`

	currTime := time.Now()

	if _, err := db.Exec(query, acc.FirstName, acc.LastName, acc.Email, acc.HashedPassword, acc.Admin, currTime, currTime); err != nil {
		return err
	}

	return nil

}

func (db *DB) UpdateAccount(acc Account) error {
	return nil

}

func (db *DB) DeleteAccount(acc Account) error {
	return nil

}

func (db *DB) GetAccount(email string) (Account, error) {

	query := `SELECT *
                FROM account
                WHERE email = $1`

	account := Account{}

	if err := db.Get(&account, query, email); err != nil {
		return account, err
	}

	return account, nil
}
