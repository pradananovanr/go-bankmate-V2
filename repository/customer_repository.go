/*
 * Author : Pradana Novan Rianto (https://github.com/pradananovanr)
 * Created on : Thu Apr 13 2023
 * Copyright : Pradana Novan Rianto © 2023. All rights reserved
 */

package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-bankmate/model/entity"
	"go-bankmate/util"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type CustomerRepo interface {
	Create(newCustomer *entity.Customer) (entity.Customer, error)
	Delete(id_customer int) error
	Login(username, password string) (string, error)
	Logout(id_customer int) error
	InsertToken(id_customer int, token string) error
	UpdateToken(id_customer int) error
	ValidateToken(id_customer int, token string) error
}

type customerRepo struct {
	db *sql.DB
}

func (c *customerRepo) Create(newCustomer *entity.Customer) (entity.Customer, error) {
	query := "INSERT INTO m_customer (username, password, email, phone) VALUES ($1, $2, $3, $4) RETURNING id_customer"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newCustomer.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return entity.Customer{}, err
	}

	var customerID int
	err = c.db.QueryRow(query, newCustomer.Username, string(hashedPassword), newCustomer.Email, newCustomer.Phone).Scan(&customerID)
	if err != nil {
		log.Println(err)
		log.Println(err)
		return entity.Customer{}, err
	}

	activity := fmt.Sprintf("create new customer with username %s", newCustomer.Username)

	query = `INSERT INTO t_log (id_customer, activity) VALUES ($1, $2)`
	_, err = c.db.Exec(query, customerID, activity)
	if err != nil {
		log.Println(err)
		return entity.Customer{}, err
	}

	newCustomer.ID_Customer = customerID

	return *newCustomer, nil
}

func (c *customerRepo) Delete(id_customer int) error {
	query := "DELETE FROM m_customer WHERE id_customer = $1"
	result, err := c.db.Exec(query, id_customer)
	if err != nil {
		log.Println(err)
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("customer with id %d not found", id_customer)
	}

	activity := fmt.Sprintf("customer with id %d deleted", id_customer)

	query = `INSERT INTO t_log (id_customer, activity) VALUES ($1, $2)`
	_, err = c.db.Exec(query, id_customer, activity)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *customerRepo) Login(username, password string) (string, error) {

	var err error

	u := entity.CustomerLogin{}

	query := "SELECT id_customer, username, password FROM m_customer WHERE username = $1"
	row := c.db.QueryRow(query, username)
	err = row.Scan(&u.ID_Customer, &u.Username, &u.Password)

	if err != nil {
		log.Println(err)
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println(err)
		return "", err
	}

	token, err := util.GenerateToken(u.ID_Customer)

	if err != nil {
		log.Println(err)
		return "", err
	}

	err = c.InsertToken(u.ID_Customer, token)

	if err != nil {
		log.Println(err)
		return "", err
	}

	activity := fmt.Sprintf("customer with id %d logged in", &u.ID_Customer)

	query = `INSERT INTO t_log (id_customer, activity) VALUES ($1, $2)`
	_, err = c.db.Exec(query, &u.ID_Customer, activity)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}

func (c *customerRepo) Logout(id_customer int) error {

	err := c.UpdateToken(id_customer)

	if err != nil {
		log.Println(err)
		return err
	}

	activity := fmt.Sprintf("customer with id %d logged out", id_customer)

	query := `INSERT INTO t_log (id_customer, activity) VALUES ($1, $2)`
	_, err = c.db.Exec(query, id_customer, activity)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *customerRepo) UpdateToken(id_customer int) error {
	query := "UPDATE t_token SET revoked = $1 WHERE id_customer = $2"
	result, err := c.db.Exec(query, true, id_customer)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("token for customer with id %d not found", id_customer)
	}

	return nil
}

func (c *customerRepo) InsertToken(id_customer int, token string) error {
	query := "INSERT INTO t_token (id_customer, token, revoked) VALUES ($1, $2, $3)"
	result, err := c.db.Exec(query, id_customer, token, false)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (c *customerRepo) ValidateToken(id_customer int, token string) error {
	var tokenString string

	query := "SELECT token FROM t_token WHERE id_customer = $1 AND revoked = false LIMIT 1"
	row := c.db.QueryRow(query, id_customer)
	err := row.Scan(&tokenString)

	if err != nil {
		log.Println(err)
		return err
	}

	if tokenString != token {
		return errors.New("invalid token")
	}

	return nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepo {
	repo := new(customerRepo)
	repo.db = db
	return repo
}
