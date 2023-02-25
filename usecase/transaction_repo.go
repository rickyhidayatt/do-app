package usecase

import (
	"errors"

	"github.com/rickyhidayatt/do-app/model"
	"github.com/rickyhidayatt/do-app/repository"
)

type TransactionUseCase interface {
	AddWallet(AddBalance int, user *model.User) error
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
}

func (tx *transactionUseCase) AddWallet(addBalance int, user *model.User) error {
	if addBalance < 10000 {
		return errors.New("Transfer failed, minimum transfer amount is 10000")
	}
	err := tx.transactionRepo.AddBalance(addBalance, user)

	if err != nil {
		return err
	}

	return nil
}

func NewTransactionUseCase(txRepoArg repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: txRepoArg,
	}
}
