package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-bankmate/model/entity"
	"log"
	"time"
)

type DepositRepo interface {
	ValidateToken(id_customer int, token string) error
	CreateDeposit(id_customer int, token string, deposit *entity.DepositRequest) (*entity.Deposit, error)
	GetDeposit(id_customer, id_deposit int, token string) (*entity.Deposit, error)
	GetAllDeposit(id_customer int, token string) ([]*entity.Deposit, error)
}

type depositRepo struct {
	db *sql.DB
}

func (d *depositRepo) ValidateToken(id int, token string) error {
	var tokenString string

	query := "SELECT token FROM t_token WHERE id_customer = $1 AND revoked = false LIMIT 1"
	row := d.db.QueryRow(query, id)
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

func (d *depositRepo) CreateDeposit(id_customer int, token string, deposit *entity.DepositRequest) (*entity.Deposit, error) {
	err := d.ValidateToken(id_customer, token)

	if err != nil {
		return &entity.Deposit{}, err
	}

	tx, err := d.db.Begin()
	if err != nil {
		return &entity.Deposit{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		} else {
			err = tx.Commit()
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}()

	var id_deposit int
	var date_time time.Time

	query := `INSERT INTO t_deposit (id_customer, deposit_amount, deposit_description) VALUES ($1, $2, $3) RETURNING id_deposit, date_time`
	err = tx.QueryRow(query, id_customer, deposit.Deposit_Amount, deposit.Deposit_Description).Scan(&id_deposit, &date_time)
	if err != nil {
		return &entity.Deposit{}, err
	}

	query = "SELECT COUNT(*) FROM t_wallet WHERE id_customer = $1"
	var count int
	err = tx.QueryRow(query, id_customer).Scan(&count)
	if err != nil {
		return &entity.Deposit{}, err
	}

	if count > 0 {
		query = `UPDATE t_wallet SET wallet_amount = wallet_amount + $1 WHERE id_customer = $2`
		_, err = tx.Exec(query, deposit.Deposit_Amount, id_customer)
		if err != nil {
			return &entity.Deposit{}, err
		}
	} else {
		query = `INSERT INTO t_wallet (id_customer, wallet_amount) VALUES ($1, $2)`
		_, err = tx.Exec(query, id_customer, deposit.Deposit_Amount)
		if err != nil {
			return &entity.Deposit{}, err
		}
	}

	activity := fmt.Sprintf("customer with id %d created new deposit", id_customer)

	query = `INSERT INTO t_log (id_customer, activity) VALUES ($1, $2)`
	_, err = tx.Exec(query, id_customer, activity)
	if err != nil {
		log.Println(err)
		return &entity.Deposit{}, err
	}

	response := &entity.Deposit{
		ID_Deposit:          id_deposit,
		Deposit_Amount:      deposit.Deposit_Amount,
		Deposit_Description: deposit.Deposit_Description,
		Date_Time:           date_time,
	}

	return response, nil
}

func (d *depositRepo) GetDeposit(id_customer, id_deposit int, token string) (*entity.Deposit, error) {
	err := d.ValidateToken(id_customer, token)

	if err != nil {
		return &entity.Deposit{}, err
	}

	var deposit entity.Deposit

	query := "SELECT id_deposit, id_customer, deposit_amount, deposit_description, date_time FROM t_deposit WHERE id_deposit = $1"
	row := d.db.QueryRow(query, id_deposit)
	err = row.Scan(&deposit.ID_Deposit, &deposit.ID_Customer, &deposit.Deposit_Amount, &deposit.Deposit_Description, &deposit.Date_Time)

	if err != nil {
		log.Println(err)
		return &entity.Deposit{}, err
	}

	activity := fmt.Sprintf("customer with id %d get deposit history by id", id_customer)

	query = `INSERT INTO t_log (id_customer, activity) VALUES ($1, $2)`
	_, err = d.db.Exec(query, id_customer, activity)
	if err != nil {
		log.Println(err)
		return &entity.Deposit{}, err
	}

	return &deposit, nil
}

func (d *depositRepo) GetAllDeposit(id_customer int, token string) ([]*entity.Deposit, error) {
	err := d.ValidateToken(id_customer, token)

	if err != nil {
		return []*entity.Deposit{}, err
	}

	var deposits []*entity.Deposit

	query := "SELECT * FROM t_deposit WHERE id_customer = $1"
	row, err := d.db.Query(query, id_customer)

	if err != nil {
		log.Println(err)
		return []*entity.Deposit{}, err
	}

	defer row.Close()
	for row.Next() {
		var deposit entity.Deposit
		if err := row.Scan(&deposit.ID_Deposit, &deposit.ID_Customer, &deposit.Deposit_Amount, &deposit.Deposit_Description, &deposit.Date_Time); err != nil {
			return []*entity.Deposit{}, err
		}
		deposits = append(deposits, &deposit)
	}
	if err := row.Err(); err != nil {
		return []*entity.Deposit{}, err
	}

	activity := fmt.Sprintf("customer with id %d get all payment history", id_customer)

	query = `INSERT INTO t_log (id_customer, activity) VALUES ($1, $2)`
	_, err = d.db.Exec(query, id_customer, activity)
	if err != nil {
		log.Println(err)
		return []*entity.Deposit{}, err
	}

	return deposits, nil
}

func NewDepositRepository(db *sql.DB) DepositRepo {
	repo := new(depositRepo)
	repo.db = db
	return repo
}
