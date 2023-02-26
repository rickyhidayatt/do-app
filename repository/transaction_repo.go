package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rickyhidayatt/do-app/model"
)

type TransactionRepository interface {
	AddBalance(userId string, amount int) error
	GetBalance(userId string) ([]int, error)
	SendBalance(userId string, amount int) error
}

type transactionRepository struct {
	db *sqlx.DB
}

func (r *transactionRepository) AddBalance(userId string, amount int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// check if user exists
	var user []model.User
	err = tx.Select(&user, "SELECT * FROM users WHERE id=$1", userId)

	if err != nil {
		return fmt.Errorf("failed to add balance: %v", err)
	}

	// add balance to user's account
	balance := model.Balances{
		UserId:  userId,
		Balance: amount,
	}

	_, err = tx.NamedExec("UPDATE balances SET balance = balance + :balance WHERE user_id = :user_id", &balance)

	fmt.Println(err)

	if err != nil {
		return fmt.Errorf("failed to add balance: %v", err)
	}

	return nil
}

func (r *transactionRepository) SendBalance(userId string, amount int) error {

	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// check if user exists
	var user []model.User
	err = tx.Select(&user, "SELECT * FROM users WHERE id=$1", userId)

	if err != nil {
		return fmt.Errorf("failed to add balance: %v", err)
	}

	// add balance to user's account
	balance := model.Balances{
		UserId:  userId,
		Balance: amount,
	}

	_, err = tx.NamedExec("UPDATE balances SET balance = balance - :balance WHERE user_id = :user_id", &balance)

	fmt.Println(err)

	if err != nil {
		return fmt.Errorf("failed to add balance: %v", err)
	}

	return nil
}

func (r *transactionRepository) GetBalance(userId string) ([]int, error) {
	var balances []model.Balances
	var balanceInt []int

	err := r.db.Select(&balances, "SELECT * FROM balances WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	for _, v := range balances {
		balanceInt = append(balanceInt, v.Balance)
	}

	return balanceInt, nil
}

func NewTransactionRepository(dbArg *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: dbArg,
	}
}
