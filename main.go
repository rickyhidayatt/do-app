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
	txRepo := repository.NewTransactionRepository(db)
	// txUsecase := usecase.NewTransactionUseCase(txRepo, userRepo)

	userId := "53-331-6070"
	addBalance := 30000

	// balanceAdded, err := txUsecase.AddWallet(userId, addBalance)
	// if err != nil {
	// 	// handle the error
	// 	fmt.Println("gagal")
	// }

	// fmt.Println(txUsecase)
	// fmt.Println(txRepo.GetBalance(userId))

	err := txRepo.SendBalance(userId, addBalance)
	if err != nil {
		fmt.Println("gagal")
	}
	fmt.Println("berhasi;")

}

// cek := repository.NewUserRepository(db)
// lihat, _ := cek.GetUserById("53-331-6070")

// fmt.Println("cek isi transactionUsecase", transactionUsecase)
// fmt.Println(lihat)

// transactionUsecase.AddWallet("53-331-6070", 200000)
// router := gin.Default()

// // // menambahkan route untuk mengambil saldo
// router.POST("/saldo/:id", func(c *gin.Context) {
// 	var ballance *model.Balance

// 	saldo, err := transactionUsecase.AddWallet(ballance.UserId, ballance.Balance)
// 	// saldo, err := transactionUsecase.AddWallet("STRING", 200000)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// mengirimkan saldo sebagai response
// 	c.JSON(200, gin.H{"saldo": saldo})
// })

// // menjalankan server
// router.Run()
