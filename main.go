package main

import (
	"fmt"

	"github.com/rickyhidayatt/do-app/config"
	"github.com/rickyhidayatt/do-app/repository"
)

func main() {
	config := config.NewConfig()
	db := config.DbConnect()

	// userRepo := repository.NewUserRepository(db)
	// txUsecase := usecase.NewTransactionUseCase(txRepo, userRepo)

	txRepo := repository.NewTransactionRepository(db)
	userId := "53-331-6070"
	addBalance := 30000

	err := txRepo.SendBalance(userId, addBalance)
	if err != nil {
		fmt.Println("gagal")
	}
	fmt.Println("berhasi;")

}
